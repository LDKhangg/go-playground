# Learning Roadmap

Work from top to bottom. Exercises teach one concept in isolation; project milestones apply it to the task API.

## Core Curriculum

- [ ] 01 Basics: variables, constants, types, conditions, and loops
- [ ] 02 Functions: parameters, multiple returns, scope, and basic pointers
- [ ] 03 Collections: arrays, slices, maps, range, and safe copies
- [ ] 04 Structs and methods: custom types, constructors, receivers, and composition
- [ ] 05 Interfaces and errors: small interfaces, sentinel errors, wrapping, and inspection
- [ ] 06 Testing: table-driven tests, subtests, coverage, and benchmarks
- [ ] 07 Concurrency: goroutines, wait groups, mutexes, contexts, and race detection

## Task API Milestones

- [x] Understand the baseline health, list, and create routes
- [ ] Add complete CRUD operations and domain validation
- [ ] Expand handler coverage with `httptest`
- [ ] Add configuration, context propagation, and graceful shutdown
- [ ] Place persistence behind a small store interface
- [ ] Add database storage and migrations
- [ ] Add middleware, structured logging, and observability
- [ ] Package with Docker and deploy the tested service

## Definition of Done

A chapter or project milestone is complete when its tests pass, `make check` passes, the README progress table is current, a learning-log reflection exists, one coherent commit records the work, and the commit is pushed to GitHub.
