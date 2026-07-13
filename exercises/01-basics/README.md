# 01 - Basics

## Goal

Implement age-based ticket pricing and use a loop to total several tickets.

## Concepts

Variables, constants, integers, comparisons, `if`, `switch`, and `for` loops.

## Exercise

Implement `TicketPrice`. Ages below 13 cost 5, ages from 13 through 64 cost 12, and ages 65 or above cost 7. Then implement `TotalTicketPrice` by looping over every supplied age and adding its ticket price.

## Acceptance Criteria

- Ages 12 and below cost `5`.
- Ages 13 through 64 cost `12`.
- Ages 65 and above cost `7`.
- `TotalTicketPrice([]int{8, 30, 70})` returns `24`.

## Hints

Name the three prices as constants. Check boundaries in an order that makes every age belong to one group. A `for range` loop can visit every age in the slice.

## Commands

`go test -tags exercise ./exercises/01-basics/...`

## Reflection Prompts

Why does the order of age checks matter? What value should the total loop produce for an empty slice?
