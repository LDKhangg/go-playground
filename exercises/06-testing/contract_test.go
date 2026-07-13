//go:build exercise

package testingpractice

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
	"testing"
)

func TestExerciseContract(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "exercise_test.go", nil, 0)
	if err != nil {
		t.Fatalf("parse exercise_test.go: %v", err)
	}

	classifyTest := findFunction(file, "TestClassify")
	if classifyTest == nil {
		t.Error("add TestClassify as a table-driven test with negative, zero, and positive named cases")
	} else {
		checkClassifyTest(t, classifyTest)
	}

	benchmark := findFunction(file, "BenchmarkClassify")
	if benchmark == nil {
		t.Error("add BenchmarkClassify with a loop controlled by b.N")
	} else {
		checkBenchmark(t, benchmark)
	}
}

func findFunction(file *ast.File, name string) *ast.FuncDecl {
	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if ok && function.Name.Name == name {
			return function
		}
	}
	return nil
}

func checkClassifyTest(t *testing.T, function *ast.FuncDecl) {
	t.Helper()

	testParameter := parameterName(function)
	rangeStatement, runCall := findTableSubtest(function.Body, testParameter)
	if rangeStatement == nil || runCall == nil {
		t.Error("TestClassify must call t.Run from inside the loop over its test table")
		return
	}

	table := findTable(function.Body, rangeStatement.X)
	if !hasRequiredCases(table) {
		t.Error("TestClassify table must have exactly three cases: negative input -> negative, zero input -> zero, and positive input -> positive")
	}

	caseVariable, ok := rangeStatement.Value.(*ast.Ident)
	if !ok || caseVariable.Name == "_" {
		t.Error("TestClassify must keep the current table case in the range loop")
		return
	}
	if len(runCall.Args) < 2 || !isCaseSelector(runCall.Args[0], caseVariable.Name) {
		t.Error("t.Run must use the current table case's name")
		return
	}

	callback, ok := runCall.Args[1].(*ast.FuncLit)
	if !ok {
		t.Error("t.Run must execute a subtest function for each table case")
		return
	}

	classifyCalls := findCalls(callback.Body, "Classify")
	if len(classifyCalls) == 0 || len(classifyCalls[0].Args) != 1 || !isCaseSelector(classifyCalls[0].Args[0], caseVariable.Name) {
		t.Error("each subtest must call Classify with the current case's input")
		return
	}

	actualNames := assignedCallResults(callback.Body, "Classify")
	comparison := findResultComparison(callback.Body, actualNames, caseVariable.Name)
	if comparison == nil {
		t.Error("each subtest must compare the Classify result with the current case's expected value")
		return
	}
	if !failureReportsValues(comparison.Body, actualNames, caseVariable.Name, testParameter) {
		t.Error("the failure message must report the input, actual value, and expected value")
	}
}

func findTableSubtest(body *ast.BlockStmt, testParameter string) (*ast.RangeStmt, *ast.CallExpr) {
	var foundRange *ast.RangeStmt
	var foundRun *ast.CallExpr
	ast.Inspect(body, func(node ast.Node) bool {
		if foundRun != nil {
			return false
		}
		rangeStatement, ok := node.(*ast.RangeStmt)
		if !ok {
			return true
		}
		ast.Inspect(rangeStatement.Body, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			selector, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			receiver, receiverOK := selector.X.(*ast.Ident)
			if receiverOK && receiver.Name == testParameter && selector.Sel.Name == "Run" {
				foundRange = rangeStatement
				foundRun = call
				return false
			}
			return true
		})
		return foundRun == nil
	})
	return foundRange, foundRun
}

func findTable(body *ast.BlockStmt, expression ast.Expr) *ast.CompositeLit {
	if literal, ok := expression.(*ast.CompositeLit); ok {
		return literal
	}
	identifier, ok := expression.(*ast.Ident)
	if !ok {
		return nil
	}

	var table *ast.CompositeLit
	ast.Inspect(body, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.AssignStmt:
			for index, left := range node.Lhs {
				name, ok := left.(*ast.Ident)
				if ok && name.Name == identifier.Name && index < len(node.Rhs) {
					table, _ = node.Rhs[index].(*ast.CompositeLit)
				}
			}
		case *ast.ValueSpec:
			for index, name := range node.Names {
				if name.Name == identifier.Name && index < len(node.Values) {
					table, _ = node.Values[index].(*ast.CompositeLit)
				}
			}
		}
		return table == nil
	})
	return table
}

func hasRequiredCases(table *ast.CompositeLit) bool {
	if table == nil || len(table.Elts) != 3 {
		return false
	}

	found := map[string]bool{}
	for _, element := range table.Elts {
		stringsInCase := map[string]int{}
		var numbers []int64
		ast.Inspect(element, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.BasicLit:
				if node.Kind == token.STRING {
					value, err := strconv.Unquote(node.Value)
					if err == nil {
						stringsInCase[value]++
					}
				} else if node.Kind == token.INT {
					value, err := strconv.ParseInt(node.Value, 0, 64)
					if err == nil {
						numbers = append(numbers, value)
					}
				}
			case *ast.UnaryExpr:
				literal, ok := node.X.(*ast.BasicLit)
				if ok && node.Op == token.SUB && literal.Kind == token.INT {
					value, err := strconv.ParseInt(literal.Value, 0, 64)
					if err == nil {
						numbers = append(numbers, -value)
					}
				}
			}
			return true
		})

		for _, category := range []string{"negative", "zero", "positive"} {
			if stringsInCase[category] < 2 {
				continue
			}
			for _, number := range numbers {
				if category == "negative" && number < 0 || category == "zero" && number == 0 || category == "positive" && number > 0 {
					found[category] = true
				}
			}
		}
	}
	return found["negative"] && found["zero"] && found["positive"]
}

func findCalls(node ast.Node, name string) []*ast.CallExpr {
	var calls []*ast.CallExpr
	ast.Inspect(node, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		identifier, ok := call.Fun.(*ast.Ident)
		if ok && identifier.Name == name {
			calls = append(calls, call)
		}
		return true
	})
	return calls
}

func assignedCallResults(node ast.Node, functionName string) map[string]bool {
	results := map[string]bool{}
	ast.Inspect(node, func(node ast.Node) bool {
		assignment, ok := node.(*ast.AssignStmt)
		if !ok {
			return true
		}
		for index, right := range assignment.Rhs {
			calls := findCalls(right, functionName)
			if len(calls) > 0 && index < len(assignment.Lhs) {
				if name, ok := assignment.Lhs[index].(*ast.Ident); ok {
					results[name.Name] = true
				}
			}
		}
		return true
	})
	return results
}

func findResultComparison(node ast.Node, actualNames map[string]bool, caseVariable string) *ast.IfStmt {
	var found *ast.IfStmt
	ast.Inspect(node, func(node ast.Node) bool {
		statement, ok := node.(*ast.IfStmt)
		if !ok || found != nil {
			return found == nil
		}
		comparison, ok := statement.Cond.(*ast.BinaryExpr)
		if !ok || comparison.Op != token.NEQ && comparison.Op != token.EQL {
			return true
		}
		leftActual := containsAnyIdent(comparison.X, actualNames) || len(findCalls(comparison.X, "Classify")) > 0
		rightActual := containsAnyIdent(comparison.Y, actualNames) || len(findCalls(comparison.Y, "Classify")) > 0
		if leftActual && containsCaseSelector(comparison.Y, caseVariable) || rightActual && containsCaseSelector(comparison.X, caseVariable) {
			found = statement
			return false
		}
		return true
	})
	return found
}

func failureReportsValues(body *ast.BlockStmt, actualNames map[string]bool, caseVariable, testParameter string) bool {
	found := false
	ast.Inspect(body, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok || len(call.Args) < 4 {
			return true
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		receiver, receiverOK := selector.X.(*ast.Ident)
		if !receiverOK || receiver.Name != testParameter || selector.Sel.Name != "Fatalf" && selector.Sel.Name != "Errorf" {
			return true
		}
		format, ok := call.Args[0].(*ast.BasicLit)
		if !ok || format.Kind != token.STRING || strings.Count(format.Value, "%") < 3 {
			return true
		}

		caseValues := 0
		hasActual := false
		for _, argument := range call.Args[1:] {
			if containsCaseSelector(argument, caseVariable) {
				caseValues++
			}
			if containsAnyIdent(argument, actualNames) || len(findCalls(argument, "Classify")) > 0 {
				hasActual = true
			}
		}
		found = caseValues >= 2 && hasActual
		return !found
	})
	return found
}

func checkBenchmark(t *testing.T, function *ast.FuncDecl) {
	t.Helper()

	benchmarkParameter := parameterName(function)
	var validLoop bool
	ast.Inspect(function.Body, func(node ast.Node) bool {
		loop, ok := node.(*ast.ForStmt)
		if !ok || validLoop {
			return !validLoop
		}
		variable := benchmarkLoopVariable(loop)
		if variable == "" || !benchmarkCondition(loop.Cond, variable, benchmarkParameter) || !increments(loop.Post, variable) {
			return true
		}
		for _, call := range findCalls(loop.Body, "Classify") {
			if len(call.Args) == 1 {
				literal, ok := call.Args[0].(*ast.BasicLit)
				validLoop = ok && literal.Kind == token.INT && literal.Value == "42"
			}
		}
		return !validLoop
	})
	if !validLoop {
		t.Error("BenchmarkClassify must call Classify(42) inside a for loop from 0 up to b.N")
	}
}

func benchmarkLoopVariable(loop *ast.ForStmt) string {
	assignment, ok := loop.Init.(*ast.AssignStmt)
	if !ok || len(assignment.Lhs) != 1 || len(assignment.Rhs) != 1 {
		return ""
	}
	name, ok := assignment.Lhs[0].(*ast.Ident)
	start, startOK := assignment.Rhs[0].(*ast.BasicLit)
	if !ok || !startOK || start.Kind != token.INT || start.Value != "0" {
		return ""
	}
	return name.Name
}

func benchmarkCondition(expression ast.Expr, variable, benchmarkParameter string) bool {
	comparison, ok := expression.(*ast.BinaryExpr)
	if !ok || comparison.Op != token.LSS && comparison.Op != token.NEQ {
		return false
	}
	name, ok := comparison.X.(*ast.Ident)
	selector, selectorOK := comparison.Y.(*ast.SelectorExpr)
	if !selectorOK {
		return false
	}
	receiver, receiverOK := selector.X.(*ast.Ident)
	return ok && receiverOK && name.Name == variable && receiver.Name == benchmarkParameter && selector.Sel.Name == "N"
}

func increments(statement ast.Stmt, variable string) bool {
	increment, ok := statement.(*ast.IncDecStmt)
	name, nameOK := increment.X.(*ast.Ident)
	return ok && nameOK && increment.Tok == token.INC && name.Name == variable
}

func containsIdent(node ast.Node, name string) bool {
	found := false
	ast.Inspect(node, func(node ast.Node) bool {
		identifier, ok := node.(*ast.Ident)
		if ok && identifier.Name == name {
			found = true
			return false
		}
		return !found
	})
	return found
}

func containsAnyIdent(node ast.Node, names map[string]bool) bool {
	for name := range names {
		if containsIdent(node, name) {
			return true
		}
	}
	return false
}

func parameterName(function *ast.FuncDecl) string {
	if function.Type.Params == nil || len(function.Type.Params.List) == 0 || len(function.Type.Params.List[0].Names) == 0 {
		return ""
	}
	return function.Type.Params.List[0].Names[0].Name
}

func isCaseSelector(expression ast.Expr, caseVariable string) bool {
	selector, ok := expression.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	receiver, receiverOK := selector.X.(*ast.Ident)
	return receiverOK && receiver.Name == caseVariable
}

func containsCaseSelector(node ast.Node, caseVariable string) bool {
	found := false
	ast.Inspect(node, func(node ast.Node) bool {
		selector, ok := node.(*ast.SelectorExpr)
		if ok && isCaseSelector(selector, caseVariable) {
			found = true
			return false
		}
		return !found
	})
	return found
}
