# 01 - Basics

## Goal

Implement age-based ticket pricing with Go's basic types and control flow.

## Concepts

Variables, constants, integers, comparisons, `if`, and `switch`.

## Exercise

Implement `TicketPrice`. Ages below 13 cost 5, ages from 13 through 64 cost 12, and ages 65 or above cost 7.

## Acceptance Criteria

- `TicketPrice(8)` returns `5`.
- `TicketPrice(30)` returns `12`.
- `TicketPrice(70)` returns `7`.

## Hints

Name the three prices as constants. Check boundaries in an order that makes every age belong to one group.

## Commands

`go test -tags exercise ./exercises/01-basics/...`

## Reflection Prompts

Why does the order of age checks matter? When would a `switch` read more clearly than chained `if` statements?
