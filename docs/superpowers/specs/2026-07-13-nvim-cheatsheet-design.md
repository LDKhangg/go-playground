# Neovim Cheatsheet Expansion Design

## Goal

Expand the existing `Super+F3` Kitty popup into a practical Vietnamese reference for the current Neovim configuration. The displayed shortcuts must match the real keymaps, clearly explain buffer/window/tab behavior, cover the installed plugins, and include a complete Go command reference.

## Scope

- Keep the current external popup implemented by `~/.config/hypr/UserScripts/CheatsheetNvim.sh`.
- Keep the current visual style and scrollable `less` interface.
- Fix the Go-local `<leader>tn` conflict so the global next-tab mapping works in Go buffers.
- Keep `<leader>tr` as the shortcut for running the nearest Go test.
- Keep all other existing Go test and DAP mappings.
- Add Go terminal commands as documentation only; do not add new run/build/vet/lint keymaps.
- Do not change unrelated plugins, styling, or IdeaVim cheatsheet behavior.

## Cheatsheet Structure

The popup will be organized in task order:

1. Mental model and popup controls: distinguish buffer, window/split, and tab; explain scrolling and closing the popup.
2. Core movement and editing: retain the useful Vim motions and editing operators already shown.
3. Buffers, windows, and tabs: show how to switch, create, close, and move between each object.
4. Neo-tree: toggle/focus/close the tree, navigate folders, open files in the current window/splits/tabs, and perform common file operations.
5. Telescope: launch configured pickers and operate inside a picker, including selection, preview scrolling, split/tab opening, help, and closing.
6. LSP and diagnostics: document every configured mapping.
7. Completion: document the active Blink default preset behavior and how to inspect mappings if plugin defaults change.
8. Formatting and linting: include both configured formatting mappings and explain automatic linting.
9. Go test and debug mappings: list the buffer-local runner and DAP shortcuts exactly as configured.
10. Go terminal commands: cover run, build, test variants, race detection, coverage, benchmarks, vet, formatting, module maintenance, dependency inspection, install, and golangci-lint.
11. Plugin and configuration maintenance: include Which-key discovery, Lazy management, Mason tooling, and useful health/status commands where they are available.

## Keymap Correction

Remove the Go buffer-local `<leader>tn` mapping from `lua/config/go.lua`. It currently overrides the global `<leader>tn` next-tab mapping whenever a Go file is active. `<leader>tr` remains the sole nearest-test mapping, while `<leader>tt`, `<leader>tT`, `<leader>tl`, and the existing DAP mappings remain unchanged.

## Accuracy Rules

- Configured mappings are copied from the current Lua files, not invented.
- Plugin-local mappings use the installed plugin defaults and are labelled as actions performed inside that plugin UI.
- Similar concepts are explicitly separated: Bufferline entries are buffers, Vim tab pages are tabs, and splits are windows.
- Commands that are documentation-only are placed in the Go terminal command section and are not presented as shortcuts.
- The popup header will advertise only close controls that actually work with `less`.

## Verification

- Run a headless Neovim startup check after editing the Lua keymap.
- Inspect active normal-mode mappings in a Go buffer to confirm `<leader>tn` resolves to `tabnext` and `<leader>tr` remains the nearest-test runner.
- Run `bash -n` on the cheatsheet script.
- Render the script in a terminal-compatible environment and inspect the output for quoting, wrapping, section order, and command accuracy.
