# 04 - Structs and Methods

## Goal

Model state with structs, compose one type from another, and change state through methods.

## Concepts

Structs, constructors, unexported fields, value receivers, pointer receivers, composition, delegation, and encapsulation.

## Exercise

Make `NewTask` preserve the title, make `Complete` mutate the task, and make `IsComplete` report current state. Then make `Project` compose a current `Task`: initialize it in `NewProject` and delegate the project's completion methods to that task.

## Acceptance Criteria

- A new task preserves its title and starts incomplete.
- Calling `Complete` changes that same task to complete.
- A new project contains a current task with the supplied title.
- Completing a project's current task is visible through `IsCurrentComplete`.

## Hints

A method that mutates a value needs a pointer receiver. A read-only method can use a value receiver for these small structs. Reuse `Task` methods instead of duplicating its state transitions in `Project`.

## Commands

`go test -tags exercise ./exercises/04-structs-methods/...`

## Reflection Prompts

Why would a value receiver fail to persist the completion change? How does composition let `Project` reuse `Task` behavior?
