# 04 - Structs and Methods

## Goal

Model state with a struct and change it through methods.

## Concepts

Structs, constructors, unexported fields, value receivers, pointer receivers, and encapsulation.

## Exercise

Make `NewTask` preserve the title, make `Complete` mutate the task, and make `IsComplete` report current state.

## Acceptance Criteria

- A new task preserves its title and starts incomplete.
- Calling `Complete` changes that same task to complete.

## Hints

A method that mutates a value needs a pointer receiver. A read-only method can use a value receiver for this small struct.

## Commands

`go test -tags exercise ./exercises/04-structs-methods/...`

## Reflection Prompts

Why would a value receiver fail to persist the completion change? What does an unexported field protect?
