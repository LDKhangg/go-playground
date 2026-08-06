# 10 - GORM

## Goal

Persist Go structs with GORM: define models, migrate them, and perform the
basic CRUD operations against a real database (SQLite in the tests).

## Concepts

- ORMs: structs as tables, fields as columns
- GORM model tags (`primaryKey`, `size`, `not null`)
- `AutoMigrate` to create and evolve tables
- `db.Create`, `db.First`, `db.Find`, `db.Model(...).Update`
- `gorm.ErrRecordNotFound` and `errors.Is`
- A repository function per operation, taking `*gorm.DB`

## Syntax Primer

GORM maps a struct to a table. Tags shape the columns:

```go
type Task struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"size:200;not null"`
	Done  bool
}
```

`AutoMigrate` creates the table (and later alters it) to match the struct:

```go
db.AutoMigrate(&Task{})
```

Queries return GORM's sentinel when a row is missing:

```go
var task Task
if err := db.First(&task, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	return Task{}, err
}
```

## Mental Model

GORM translates struct operations into SQL: `Create` is an `INSERT`, `First`
is a `SELECT ... LIMIT 1`, `Find` is a `SELECT`, and
`Model(&task).Update("done", true)` is an `UPDATE`. The `*gorm.DB` is the
connection handle; every call returns a chain you end with `.Error`.

## Annotated Examples

```go
func countTasks(db *gorm.DB) (int64, error) {
	var total int64
	if err := db.Model(&Task{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
```

## Common Diagnostics

- `no such table: tasks`: `AutoMigrate` never ran; migrate before any query.
- `record not found`: `First` found nothing; check with `errors.Is(err, gorm.ErrRecordNotFound)` instead of `err == nil`.
- `NOT NULL constraint failed`: a required column was empty — validate before `Create`.
- `gorm: model value required`: the query argument is a value, not a pointer (`&Task{}`).
- `undefined: gorm`: add `import "gorm.io/gorm"` and run `go mod tidy`.

## Exercise

Implement the four repository functions:

- `CreateTask(db, title)` — insert a task and return it with its assigned `ID`.
- `FindTask(db, id)` — load one task by primary key; propagate `ErrNotFound`.
- `ListTasks(db)` — return all tasks.
- `MarkDone(db, id)` — load the task, mark it done, persist the change, return the updated task.

## Acceptance Criteria

- `CreateTask` returns a task with a non-zero `ID` and the given title.
- `FindTask` round-trips a created task.
- `FindTask(999)` returns an error matching `gorm.ErrRecordNotFound`.
- `MarkDone` persists `Done = true` so a later `FindTask` sees it.
- `ListTasks` returns every created task.

## Hints

- Create: `db.Create(&task)` — pass a pointer so GORM fills `task.ID`.
- Find by id: `db.First(&task, id)` and return the `.Error`.
- Update: `db.Model(&task).Update("done", true)` after loading, or load then update.
- Every GORM chain ends with `.Error`; return it.

## Verify

```bash
gofmt -w exercises/10-gorm
go test -tags exercise ./exercises/10-gorm/...
```

This chapter depends on GORM and the SQLite driver, already declared in
`go.mod`. If dependencies ever drift, run `go mod tidy`.

## Reflection Prompts

When does an ORM help, and when is plain SQL clearer? Why do repository
functions receive `*gorm.DB` instead of calling a global? How would you
express a `WHERE done = true` query in GORM?