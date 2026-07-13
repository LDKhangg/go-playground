//go:build exercise

package testingpractice

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
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

	var hasRange, hasSubtest, callsClassify bool
	names := map[string]bool{"negative": false, "zero": false, "positive": false}
	ast.Inspect(function.Body, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.RangeStmt:
			hasRange = true
		case *ast.CallExpr:
			if identifier, ok := node.Fun.(*ast.Ident); ok && identifier.Name == "Classify" {
				callsClassify = true
			}
			if selector, ok := node.Fun.(*ast.SelectorExpr); ok && selector.Sel.Name == "Run" {
				hasSubtest = true
			}
		case *ast.BasicLit:
			if node.Kind == token.STRING {
				value, err := strconv.Unquote(node.Value)
				if err == nil {
					names[value] = true
				}
			}
		}
		return true
	})

	if !hasRange {
		t.Error("TestClassify must range over a table of cases")
	}
	if !hasSubtest {
		t.Error("TestClassify must run named subtests with t.Run")
	}
	if !callsClassify {
		t.Error("TestClassify must call Classify")
	}
	for _, name := range []string{"negative", "zero", "positive"} {
		if !names[name] {
			t.Errorf("TestClassify table must include a %q named case", name)
		}
	}
}

func checkBenchmark(t *testing.T, function *ast.FuncDecl) {
	t.Helper()

	var hasLoop, usesN, callsClassify bool
	ast.Inspect(function.Body, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.ForStmt:
			hasLoop = true
		case *ast.SelectorExpr:
			if node.Sel.Name == "N" {
				usesN = true
			}
		case *ast.CallExpr:
			if identifier, ok := node.Fun.(*ast.Ident); ok && identifier.Name == "Classify" {
				callsClassify = true
			}
		}
		return true
	})

	if !hasLoop || !usesN || !callsClassify {
		t.Error("BenchmarkClassify must call Classify in a loop controlled by b.N")
	}
}
