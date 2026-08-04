# Neovim Cheatsheet Expansion Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make the `Super+F3` popup an accurate, detailed guide to the active Neovim setup and restore `<leader>tn` as next-tab in Go buffers.

**Architecture:** Keep the existing Kitty plus `less` popup and manually maintain its categorized rows. Make one minimal Lua change to remove the buffer-local collision; derive plugin-local keys from the installed Neo-tree, Telescope, and Blink sources so the documentation matches this machine.

**Tech Stack:** Neovim Lua, Bash, Kitty, `less`, lazy.nvim plugins, Go CLI

## Global Constraints

- Keep `~/.config/hypr/UserScripts/CheatsheetNvim.sh` as the `Super+F3` entry point.
- Keep `<leader>tr` as nearest Go test and all other existing Go test/DAP mappings.
- Add no new Go run/build/vet/lint keymaps.
- Do not change the IdeaVim cheatsheet or unrelated Neovim behavior.
- Do not create a git commit unless the user explicitly requests one.

---

### Task 1: Restore Next-Tab Behavior in Go Buffers

**Files:**
- Modify: `/home/kane/.config/nvim/lua/config/go.lua:101-102`

**Interfaces:**
- Consumes: global `<leader>tn` mapping from `/home/kane/.config/nvim/lua/config/keymap.lua`
- Produces: `<leader>tn` remains `tabnext` in Go buffers; `<leader>tr` remains the nearest-test runner

- [ ] **Step 1: Capture the current conflicting mappings**

Run:

```bash
nvim --headless +'edit /tmp/cheatsheet-map.go' +'lua print(vim.inspect(vim.fn.maparg("<leader>tn", "n", false, true)))' +'lua print(vim.inspect(vim.fn.maparg("<leader>tr", "n", false, true)))' +qa
```

Expected before the change: both mappings are buffer-local Go test mappings.

- [ ] **Step 2: Remove only the duplicate nearest-test mapping**

Change:

```lua
set_runner_key(opts, "<leader>tn", "Go test nearest test", run_nearest_test)
set_runner_key(opts, "<leader>tr", "Go test nearest test", run_nearest_test)
```

to:

```lua
set_runner_key(opts, "<leader>tr", "Go test nearest test", run_nearest_test)
```

- [ ] **Step 3: Verify startup and active mappings**

Run:

```bash
nvim --headless +'edit /tmp/cheatsheet-map.go' +'lua local tn=vim.fn.maparg("<leader>tn", "n", false, true); assert(tn.buffer == 0 and tn.rhs:find("tabn"), "<leader>tn is not global tabnext")' +'lua local tr=vim.fn.maparg("<leader>tr", "n", false, true); assert(tr.buffer == 1 and tr.desc == "Go test nearest test", "<leader>tr is not the Go nearest-test mapping")' +qa
```

Expected: exit status `0` with no assertion or startup errors.

### Task 2: Expand and Correct the Popup Reference

**Files:**
- Modify: `/home/kane/.config/hypr/UserScripts/CheatsheetNvim.sh:12-101`

**Interfaces:**
- Consumes: active custom mappings from `~/.config/nvim/lua/config/*.lua`, installed plugin defaults, and Go CLI commands
- Produces: scrollable ANSI-colored reference opened by `Super+F3`

- [ ] **Step 1: Add a compact note helper and truthful popup controls**

Keep `line`, `section`, and `key`, then add:

```bash
note() { printf "  ${DIM}%-24s  %s${R}\n" "$1" "$2"; }
```

Use a header that says `q to close`; add rows for `Up/Down`, `PgUp/PgDn`, `g/G`, and `/text` as `less` controls. Do not advertise Escape as a close key because `less` exits with `q`.

- [ ] **Step 2: Reorganize the reference around actual Neovim objects**

Add a `BUFFER / WINDOW / TAB` section with these exact concepts and mappings:

```text
Buffer                    File loaded in memory; Bufferline shows these
Window / split            Viewport displaying a buffer
Tab page                  Workspace containing one or more windows
Tab / Shift-Tab           Next / previous buffer
<leader>x                 Delete current buffer
<leader>sv / sh           Vertical / horizontal split
<leader>se / sx           Equalize / close current split
Ctrl-h/j/k/l              Move across Nvim splits or tmux panes
<leader>to / tx           New / close tab page
<leader>tn / tp           Next / previous tab page
<leader>tf                Current buffer in a new tab page
gt / gT / {count}gt       Next / previous / numbered tab page
```

Retain the core movement, editing, search, mark, macro, save, quit, and formatting entries, while eliminating duplicate rows.

- [ ] **Step 3: Document plugin operation using active defaults**

Add categorized rows for:

```text
Neo-tree: <leader>e, q, Enter, Space, S, s, t, a, A, d, r, y/x/p, H, R, /, Backspace, ., ?
Telescope launch: <leader>ff, <leader>fg
Telescope picker: Ctrl-n/p, Enter, Ctrl-x/v/t, Ctrl-u/d, Tab/Shift-Tab, Ctrl-q, Ctrl-/, Ctrl-c/Esc
LSP: gd, gD, gR, gi, gt, K, <leader>ca/rn/d/D, [d/]d
Blink: Ctrl-Space, Ctrl-n/p, Up/Down, Enter/Ctrl-y, Ctrl-e, Ctrl-b/f, Ctrl-k, Tab/Shift-Tab
Formatting/lint: <leader>fm, <leader>mp, format-on-save, lint-on-events
Auto-session: :AutoSession save/restore/search/toggle
Maintenance: :Lazy, :Lazy sync, :Mason, :LspInfo, :ConformInfo, :checkhealth, :TSUpdate
```

Label keys used inside Neo-tree or Telescope so they are not confused with global mappings.

- [ ] **Step 4: Document the configured Go workflow and terminal commands**

List these existing mappings exactly:

```text
<leader>tr / tt / tT / tl
<leader>td / tD
<leader>dc / db / do / di / dO / dl / du
<leader>ih / it / im
```

Add command rows covering:

```bash
go run .
go build ./...
go build -o bin/app .
go test ./...
go test -run '^TestName$' ./...
go test -run 'TestName/Subtest' ./...
go test -count=1 ./...
go test -race ./...
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
go test -bench=. -benchmem ./...
go test -c ./path/to/package
go vet ./...
gofmt -w .
gofmt -d .
goimports -w .
gofumpt -w .
go mod tidy
go mod download
go mod verify
go list ./...
go list -m all
go get module@version
go install package@version
golangci-lint run ./...
golangci-lint run --fix ./...
GOOS=linux GOARCH=amd64 go build -o app-linux .
```

Use descriptions that distinguish compile, run, test, static analysis, formatting, modules, installation, and cross-compilation.

- [ ] **Step 5: Validate Bash syntax and rendered text**

Run:

```bash
bash -n /home/kane/.config/hypr/UserScripts/CheatsheetNvim.sh
```

Expected: exit status `0` with no output.

Extract and run the popup body without Kitty, then confirm every section renders and no shell quoting is broken:

```bash
timeout 2s env PAGER=cat /home/kane/.config/hypr/UserScripts/CheatsheetNvim.sh
```

Expected: Kitty may fail without a graphical socket in headless execution; the final interactive check is opening `Super+F3`, scrolling from top to bottom, and closing with `q`.

### Task 3: Final Cross-Check

**Files:**
- Verify: `/home/kane/.config/nvim/lua/config/go.lua`
- Verify: `/home/kane/.config/hypr/UserScripts/CheatsheetNvim.sh`

**Interfaces:**
- Consumes: outputs of Tasks 1 and 2
- Produces: evidence that documentation and active behavior agree

- [ ] **Step 1: Search for stale conflicting documentation**

Confirm there is exactly one `<leader>tn` entry in the popup and that it describes next tab, while nearest Go test is only `<leader>tr`.

- [ ] **Step 2: Run all automated checks together**

Run the headless Neovim mapping assertions from Task 1 and `bash -n` from Task 2.

Expected: every command exits `0`.

- [ ] **Step 3: Inspect the final diff without committing**

Review only the intended Lua and Bash changes plus the approved design/plan documents. Do not modify or revert unrelated worktree changes.
