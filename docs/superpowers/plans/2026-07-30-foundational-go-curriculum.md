# Foundational Go Curriculum Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Turn the existing solution-free exercises into complete English-language lessons for foundational Go syntax, while removing non-actionable Neovim diagnostics from the learning workflow.

**Architecture:** Each of the seven chapter README files becomes a self-contained lesson, retaining its starter code and challenge tests. Neovim keeps compiler and type diagnostics but applies `exercise` build tags only for this workspace and moves linting to an explicit command. Every changed tracked file is committed and pushed separately.

**Tech Stack:** Go 1.26 standard library, gopls, Neovim 0.12 Lua LSP API, nvim-lint, Git, GitHub.

## Global Constraints

- All learning prose and examples are English.
- Retain the existing seven chapters; do not add chapters or backend milestones.
- Preserve all starter code and challenge tests without implementation answers.
- Challenge tests continue to use the `exercise` build tag until a learner completes that chapter.
- Keep gopls compiler and type diagnostics; disable only style/static-analysis noise.
- Apply `-tags=exercise` only when gopls starts in `/home/kane/Dev/go-playground`.
- Commit and push each changed tracked file separately to that file's `origin/main`.
- Never stage the pre-existing untracked workflow documents unrelated to this plan.

---

### Task 1: Configure gopls for Learning Exercises

**Files:**
- Modify: `/home/kane/dotfiles-vim/.config/nvim/lua/plugins/lsp.lua`
- Modify: `/home/kane/.config/nvim/lua/plugins/lsp.lua`

**Interfaces:**
- Produces: a `gopls` configuration that enables gopls, preserves `unusedparams` and `gofumpt`, disables `staticcheck`, and sets `buildFlags = { "-tags=exercise" }` only for the Go Playground workspace.

- [ ] **Step 1: Update the canonical dotfiles configuration**

Replace the plugin configuration with a `gopls` server entry containing the following behavior:

```lua
local playground_root = vim.fs.normalize(vim.fn.expand("~/Dev/go-playground"))

vim.lsp.config("gopls", {
  before_init = function(_, config)
    if vim.fs.normalize(config.root_dir) == playground_root then
      config.settings = vim.tbl_deep_extend("force", config.settings or {}, {
        gopls = { buildFlags = { "-tags=exercise" } },
      })
    end
  end,
  settings = {
    gopls = {
      analyses = { unusedparams = true },
      staticcheck = false,
      gofumpt = true,
    },
  },
})
```

Keep `gopls` in Mason's `ensure_installed` list and in `vim.lsp.enable`.

- [ ] **Step 2: Apply the same configuration to the active Neovim file**

Make `/home/kane/.config/nvim/lua/plugins/lsp.lua` byte-for-byte equivalent to the canonical dotfiles version for the `gopls` configuration. Preserve its Lua, Python, and TypeScript servers.

- [ ] **Step 3: Verify the Lua configuration and gopls behavior**

Run:

```bash
nvim --headless '+lua require("lazy").load({plugins={"nvim-lspconfig"}})' '+qall'
nvim --headless exercises/04-structs-methods/exercise_test.go '+lua vim.wait(3000)' '+lua vim.print(vim.diagnostic.get(0))' '+qall'
```

Expected: Neovim exits without Lua errors; the tagged test buffer does not report that it is excluded due to build tags.

- [ ] **Step 4: Commit and push the canonical file only**

Run in `/home/kane/dotfiles-vim`:

```bash
git add .config/nvim/lua/plugins/lsp.lua
git commit -m "fix: tune gopls for learning exercises"
git push origin main
```

### Task 2: Stop Automatic Go Lint Noise

**Files:**
- Modify: `/home/kane/dotfiles-vim/.config/nvim/lua/plugins/lint.lua`
- Modify: `/home/kane/.config/nvim/lua/plugins/lint.lua`

**Interfaces:**
- Produces: automatic linting for JavaScript, TypeScript, and Python; Go remains available through the existing explicit `<Space>tl` keymap but is excluded from automatic `nvim-lint` runs.

- [ ] **Step 1: Update the canonical lint configuration**

Add the Go mapping to the existing JavaScript, TypeScript, and Python mappings, then add an autocmd for `BufEnter`, `BufWritePost`, and `InsertLeave`. The callback skips Go and runs `lint.try_lint()` for the other mapped filetypes.

```lua
callback = function()
  if vim.bo.filetype ~= "go" then
    lint.try_lint()
  end
end,
```

- [ ] **Step 2: Apply the same lint configuration to the active Neovim file**

Make `/home/kane/.config/nvim/lua/plugins/lint.lua` equivalent to the canonical file: it retains automatic JavaScript, TypeScript, and Python linting but excludes Go from the autocmd.

- [ ] **Step 3: Verify the plugin file parses**

Run:

```bash
nvim --headless '+lua assert(loadfile(vim.fn.expand("~/.config/nvim/lua/plugins/lint.lua")))' '+qall'
```

Expected: exit code 0.

- [ ] **Step 4: Commit and push the canonical file only**

Run in `/home/kane/dotfiles-vim`:

```bash
git add .config/nvim/lua/plugins/lint.lua
git commit -m "fix: run Go linting on demand"
git push origin main
```

### Task 3: Update the Root Learning Dashboard

**Files:**
- Modify: `README.md`

**Interfaces:**
- Consumes: the seven chapter paths and the opt-in test workflow.
- Produces: a dashboard that tells learners each chapter README is a full lesson and distinguishes baseline from active-exercise commands.

- [ ] **Step 1: Add a Chapter Format section**

After the learning path, add `## Chapter Format` stating that every chapter provides a Goal, Concepts, Syntax Primer, Mental Model, Annotated Examples, Common Diagnostics, Exercise, Acceptance Criteria, Hints, Verify, and Reflection Prompts. State that no completed solution is included.

- [ ] **Step 2: Clarify the daily workflow**

Add `go test -tags exercise ./exercises/<chapter>/...` as the active-exercise command and explain that `go test ./...` deliberately excludes unfinished tagged tests. Link to `exercises/README.md` for build-tag details.

- [ ] **Step 3: Verify links and baseline tests**

Run:

```bash
go test ./...
git diff --check -- README.md
```

Expected: Go tests pass and the diff check emits no output.

- [ ] **Step 4: Commit and push this file**

```bash
git add README.md
git commit -m "docs: explain the Go lesson format"
git push origin main
```

### Task 4: Document the Shared Exercise Workflow

**Files:**
- Modify: `exercises/README.md`

**Interfaces:**
- Produces: the canonical explanation for the `exercise` build tag, intentionally failing starter code, and diagnostic categories.

- [ ] **Step 1: Add diagnostic guidance**

Add `## Reading Diagnostics` explaining that compiler/type errors prevent a program from building, static-analysis warnings are suggestions, and tagged tests are intentionally excluded unless the `exercise` tag is supplied. Include the exact command:

```bash
go test -tags exercise ./exercises/<chapter>/...
```

- [ ] **Step 2: Clarify the start and finish workflow**

State that an active chapter test can fail because its acceptance criteria are not implemented. Preserve the current instruction to remove `//go:build exercise` only after a chapter is complete.

- [ ] **Step 3: Verify the guide references every chapter**

Run:

```bash
for chapter in exercises/[0-9][0-9]-*/; do test -f "${chapter}README.md"; done
git diff --check -- exercises/README.md
```

Expected: exit code 0 with no diff-check output.

- [ ] **Step 4: Commit and push this file**

```bash
git add exercises/README.md
git commit -m "docs: explain exercise diagnostics"
git push origin main
```

### Task 5: Expand Lessons 01 and 02

**Files:**
- Modify: `exercises/01-basics/README.md`
- Modify: `exercises/02-functions/README.md`

**Interfaces:**
- Consumes: `TicketPrice`, `TotalTicketPrice`, `Divide`, `ErrDivideByZero`, and `Double` starter APIs.
- Produces: syntax lessons for basic syntax and function/pointer/error syntax.

- [ ] **Step 1: Rewrite Chapter 01 README**

Add all required lesson sections. Its Syntax Primer must show `var`, `:=`, `const`, `int`, `string`, zero values, explicit conversion, comparisons, `if`, `switch`, and `for range`. Its annotated example must use ticket prices and iteration. Its Common Diagnostics section must explain `declared and not used`, mismatched types, and missing braces. Do not reveal the `TicketPrice` implementation.

- [ ] **Step 2: Verify Chapter 01**

Run:

```bash
go test -tags exercise ./exercises/01-basics/...
git diff --check -- exercises/01-basics/README.md
```

Expected: the test passes and the diff check emits no output.

- [ ] **Step 3: Commit and push Chapter 01**

```bash
git add exercises/01-basics/README.md
git commit -m "docs: expand basics syntax lesson"
git push origin main
```

- [ ] **Step 4: Rewrite Chapter 02 README**

Add all required lesson sections. Its Syntax Primer must show function declarations, grouped parameters, multiple results, `error`, `return`, `errors.New`, pointers, `&`, and `*`. Its annotated example must use `Divide` and `Double`. Its Common Diagnostics section must explain assigning one result to multiple values, dereferencing a non-pointer, and ignoring an error. Do not reveal the solution.

- [ ] **Step 5: Verify and commit Chapter 02**

Run:

```bash
go test -tags exercise ./exercises/02-functions/...
git diff --check -- exercises/02-functions/README.md
git add exercises/02-functions/README.md
git commit -m "docs: expand functions syntax lesson"
git push origin main
```

Expected: test passes and the commit contains only the Chapter 02 README.

### Task 6: Expand Lessons 03 and 04

**Files:**
- Modify: `exercises/03-collections/README.md`
- Modify: `exercises/04-structs-methods/README.md`

**Interfaces:**
- Consumes: `UniqueWords`, `SumArray`, `Task`, `Project`, and their existing methods.
- Produces: syntax lessons for collections and mutable struct state.

- [ ] **Step 1: Rewrite Chapter 03 README and publish it**

Add all required lesson sections. Cover array length as part of its type, slices as descriptors over backing arrays, `make`, map lookup with its boolean result, `range`, `append`, `len`, nil versus empty slices, and copying into independent storage. Explain the `assignment copies lock value`-style class of ownership warning only as a later-topic note; focus Common Diagnostics on indexing out of range and assignment type mismatches.

Verify with `go test -tags exercise ./exercises/03-collections/...` and `git diff --check -- exercises/03-collections/README.md`, then commit and push only that README with `docs: expand collections syntax lesson`.

- [ ] **Step 2: Rewrite Chapter 04 README and publish it**

Add all required lesson sections. Cover struct declarations and literals, fields, constructors as ordinary functions, method receivers, value copying, pointer receivers, `&`, composition, and unexported fields. In Common Diagnostics, explain the four visible starter warnings: package documentation as style, and unused fields as static analysis until `NewTask`, `Complete`, `IsComplete`, and project methods use state. State that these are not syntax errors.

Verify that the tagged test still reports incomplete starter behavior with `! go test -tags exercise ./exercises/04-structs-methods/...`, then run `git diff --check -- exercises/04-structs-methods/README.md`. Commit and push only that README with `docs: expand structs and methods lesson`.

### Task 7: Expand Lessons 05 and 06

**Files:**
- Modify: `exercises/05-interfaces-errors/README.md`
- Modify: `exercises/06-testing/README.md`

**Interfaces:**
- Consumes: `TitleValidator`, `ValidateTitle`, `ErrEmptyTitle`, and `Classify`.
- Produces: syntax lessons for behavior-oriented design/errors and Go's standard testing API.

- [ ] **Step 1: Rewrite Chapter 05 README and publish it**

Add all required lesson sections. Cover interface declarations, implicit implementation, accepting an interface, the `error` interface, sentinel errors, `nil`, `fmt.Errorf`, `%w`, and `errors.Is`. Explain that an interface is satisfied by method set rather than an explicit declaration. Common Diagnostics must distinguish `nil` errors from non-nil wrapped errors and explain interface method mismatch messages.

Verify that the tagged test still reports incomplete starter behavior with `! go test -tags exercise ./exercises/05-interfaces-errors/...`, then run `git diff --check -- exercises/05-interfaces-errors/README.md`. Commit and push only that README with `docs: expand interfaces and errors lesson`.

- [ ] **Step 2: Rewrite Chapter 06 README and publish it**

Add all required lesson sections. Cover `_test.go`, `testing.T`, `testing.B`, `TestXxx`, `BenchmarkXxx`, table structs, `t.Run`, `t.Fatalf`, `b.N`, `b.ResetTimer`, and coverage. Explain this chapter's exception: starter production code is complete, while the learner writes the test and benchmark. Common Diagnostics must explain why a test command can pass with no tests and why an incorrectly named test is not discovered.

Verify with `go test -tags exercise -cover ./exercises/06-testing/...`, `go test -tags exercise -bench . ./exercises/06-testing/...`, and `git diff --check -- exercises/06-testing/README.md`; then commit and push only that README with `docs: expand testing syntax lesson`.

### Task 8: Expand Lesson 07

**Files:**
- Modify: `exercises/07-concurrency/README.md`

**Interfaces:**
- Consumes: `Counter`, `(*Counter).Increment`, `Counter.Value`, and `Sum`.
- Produces: the final foundational lesson covering synchronization, channels, and cancellation.

- [ ] **Step 1: Rewrite Chapter 07 README**

Add all required lesson sections. Cover `go`, goroutines, unbuffered and receive-only channels, channel close behavior, `range` over a channel, `select`, `sync.Mutex`, `Lock`, `Unlock`, `defer`, `context.Context`, `ctx.Done`, and the race detector. The annotated examples must illustrate synchronization without implementing `Counter` or `Sum`. Common Diagnostics must explain race reports, deadlock symptoms, send-on-closed-channel, and why an unused or unsynchronized field in a starter is not syntax failure.

- [ ] **Step 2: Verify and publish Chapter 07**

Run:

```bash
! go test -race -tags exercise ./exercises/07-concurrency/...
git diff --check -- exercises/07-concurrency/README.md
git add exercises/07-concurrency/README.md
git commit -m "docs: expand concurrency syntax lesson"
git push origin main
```

Expected: the command reports incomplete starter behavior, the diff check emits no output, and the commit contains only the Chapter 07 README.

### Task 9: Final Curriculum Verification

**Files:**
- Verify: all changed Go Playground and dotfiles files

**Interfaces:**
- Produces: a pushed granular Git history and a learning repository whose documentation and diagnostics agree.

- [ ] **Step 1: Verify every lesson section**

Run:

```bash
for chapter in exercises/[0-9][0-9]-*/README.md; do
  for heading in "Syntax Primer" "Mental Model" "Annotated Examples" "Common Diagnostics" "Verify"; do
    grep -F "## ${heading}" "$chapter" >/dev/null || exit 1
  done
done
```

Expected: exit code 0.

- [ ] **Step 2: Verify baseline and active exercises**

Run:

```bash
go test ./...
go test -tags exercise ./exercises/01-basics/...
go test -tags exercise ./exercises/02-functions/...
go test -tags exercise ./exercises/03-collections/...
! go test -tags exercise ./exercises/04-structs-methods/...
! go test -tags exercise ./exercises/05-interfaces-errors/...
go test -tags exercise ./exercises/06-testing/...
! go test -race -tags exercise ./exercises/07-concurrency/...
```

Expected: baseline plus Chapters 01, 02, 03, and 06 exit 0. Chapters 04, 05, and 07 exit nonzero because their starter APIs are deliberately incomplete; their failures must name unmet exercise expectations rather than syntax or package-loading errors.

- [ ] **Step 3: Verify histories and remote state**

Run in both repositories:

```bash
git status --short
git diff --check
git log --oneline -20
git ls-remote origin refs/heads/main
```

Expected: no intended changes remain unstaged, every changed tracked file has an individual commit, and each local `main` SHA matches the remote SHA.
