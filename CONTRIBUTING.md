# Contributing Learning Work

## Prerequisites

- Go 1.26 or newer
- Git
- Make

## Workflow

1. Pull the latest `main` branch.
2. Create a focused branch such as `learn/collections` or `feat/task-delete`.
3. Run the relevant test before editing.
4. Make the smallest change that satisfies one requirement.
5. Run `make check`.
6. Update the root progress table and learning log when finishing a milestone.
7. Commit and push the branch.

## Completing a Chapter

Remove the `exercise` build constraint from a solved chapter's test, run `make check`, mark the chapter complete in `README.md`, and add a reflection to `docs/learning-log.md`.

## Commit Examples

```text
learn: complete slices exercise
test: cover invalid task payloads
feat: add task deletion
docs: record concurrency notes
```

Keep each commit limited to one exercise, one behavior change, or one documentation milestone.
