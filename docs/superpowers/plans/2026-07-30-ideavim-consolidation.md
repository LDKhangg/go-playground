# IdeaVim Consolidation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Store the merged IdeaVim configuration as the only file in the XDG location that IntelliJ reads directly.

**Architecture:** `~/.config/ideavim` is a standalone Git repository. Its `ideavimrc` is moved from the active `~/.ideavimrc` and extended only with non-conflicting legacy aliases. The old dotfiles repositories cease tracking IdeaVim configuration while retaining their Neovim and desktop files.

**Tech Stack:** IdeaVim Vimscript, Git, POSIX shell utilities.

## Global Constraints

- Use `~/.config/ideavim/ideavimrc` as the only physical IdeaVim configuration file.
- Do not create a symlink, copy, or sync script for IdeaVim configuration.
- Preserve the active IntelliJ configuration as the merge base.
- Keep the Neovim configuration in `~/dotfiles-vim` intact.
- Do not add the legacy `<leader>fg` mapping because it conflicts with the active Search Everywhere mapping.

---

### Task 1: Create The XDG IdeaVim Repository And Merged Configuration

**Files:**
- Create: `~/.config/ideavim/.git/`
- Create: `~/.config/ideavim/ideavimrc`
- Delete: `~/.ideavimrc`

**Interfaces:**
- Consumes: the active `~/.ideavimrc` configuration.
- Produces: `~/.config/ideavim/ideavimrc`, which IdeaVim discovers via XDG.

- [ ] **Step 1: Confirm the destination does not already hold an IdeaVim configuration**

Run:

```bash
test ! -e "$HOME/.config/ideavim/ideavimrc"
```

Expected: exit code `0`.

- [ ] **Step 2: Create the repository and move the active configuration**

Run:

```bash
mkdir -p "$HOME/.config/ideavim"
git -C "$HOME/.config/ideavim" init
mv "$HOME/.ideavimrc" "$HOME/.config/ideavim/ideavimrc"
```

Expected: `~/.ideavimrc` no longer exists, and the XDG directory is a Git repository.

- [ ] **Step 3: Add the non-conflicting legacy aliases**

Add the following Vimscript blocks to `~/.config/ideavim/ideavimrc` in their matching mapping groups:

```vim
" ---- legacy-compatible aliases ----
nmap <leader>e  <Action>(ActivateProjectToolWindow)
nmap <leader>sf <Action>(GotoFile)
nmap <leader>sg <Action>(SearchEverywhere)
nmap <leader>ss <Action>(GotoSymbol)
nmap <leader>fm <Action>(ReformatCode)
nmap <leader>oi <Action>(OptimizeImports)
nmap <leader>d  <Action>(ShowErrorDescription)

" ---- legacy-compatible split aliases ----
nmap <leader>sv <C-w>v
nmap <leader>sh <C-w>s
nmap <leader>se <C-w>=
nmap <leader>sx <C-w>c
```

Do not add `nmap <leader>fg <Action>(FindInPath)`: the active configuration retains `<leader>fg` for Search Everywhere and `<leader>fp` for Find in Path.

- [ ] **Step 4: Verify the XDG config contains the merged mapping set**

Run:

```bash
test ! -e "$HOME/.ideavimrc"
test -f "$HOME/.config/ideavim/ideavimrc"
git -C "$HOME/.config/ideavim" rev-parse --is-inside-work-tree
rg -n '^nmap <leader>(e|sf|sg|ss|fm|oi|d|sv|sh|se|sx) ' "$HOME/.config/ideavim/ideavimrc"
```

Expected: the first two checks exit `0`, Git prints `true`, and `rg` reports all eleven aliases.

- [ ] **Step 5: Create the initial local repository commit when explicitly requested**

Run:

```bash
git -C "$HOME/.config/ideavim" add ideavimrc
git -C "$HOME/.config/ideavim" commit -m "feat: consolidate IdeaVim configuration"
```

Expected: one commit tracking only `ideavimrc`.

### Task 2: Remove Duplicate IdeaVim Management

**Files:**
- Delete: `~/dotfiles-vim/.ideavimrc`
- Modify: `~/dotfiles-vim/scripts/install.sh:25,29-31`
- Modify: `~/dotfiles-vim/scripts/sync-from-home.sh:25,29-31`
- Delete: `~/dotfiles-ldkhangg/home/.ideavimrc`

**Interfaces:**
- Consumes: the XDG repository from Task 1.
- Produces: no remaining IdeaVim configuration file in either older dotfiles repository.

- [ ] **Step 1: Remove the duplicate tracked config files**

Run:

```bash
git -C "$HOME/dotfiles-vim" rm .ideavimrc
git -C "$HOME/dotfiles-ldkhangg" rm -f home/.ideavimrc
```

Expected: both files are staged as deleted; no Neovim file is changed.

- [ ] **Step 2: Remove only IdeaVim copy/sync commands from the Vim repository scripts**

Delete these exact blocks from both `~/dotfiles-vim/scripts/install.sh` and `~/dotfiles-vim/scripts/sync-from-home.sh`:

```bash
copy_file "$REPO_ROOT/.ideavimrc" "$HOME_DIR/.ideavimrc"
copy_file \
  "$REPO_ROOT/.config/JetBrains/IntelliJIdea2026.1/options/vim_settings.xml" \
  "$HOME_DIR/.config/JetBrains/IntelliJIdea2026.1/options/vim_settings.xml"
```

For `sync-from-home.sh`, use the reverse copy direction in the removed block:

```bash
copy_file "$HOME_DIR/.ideavimrc" "$REPO_ROOT/.ideavimrc"
copy_file \
  "$HOME_DIR/.config/JetBrains/IntelliJIdea2026.1/options/vim_settings.xml" \
  "$REPO_ROOT/.config/JetBrains/IntelliJIdea2026.1/options/vim_settings.xml"
```

Keep the Neovim, `editor.xml`, and keymap copy commands unchanged.

- [ ] **Step 3: Verify scripts remain valid and duplicates are gone**

Run:

```bash
bash -n "$HOME/dotfiles-vim/scripts/install.sh"
bash -n "$HOME/dotfiles-vim/scripts/sync-from-home.sh"
test ! -e "$HOME/dotfiles-vim/.ideavimrc"
test ! -e "$HOME/dotfiles-ldkhangg/home/.ideavimrc"
```

Expected: all commands exit `0`.

- [ ] **Step 4: Commit cleanup in each existing repository when explicitly requested**

Run:

```bash
git -C "$HOME/dotfiles-vim" add -u
git -C "$HOME/dotfiles-vim" commit -m "chore: remove IdeaVim configuration"
git -C "$HOME/dotfiles-ldkhangg" add -u home/.ideavimrc
git -C "$HOME/dotfiles-ldkhangg" commit -m "chore: move IdeaVim configuration"
```

Expected: each commit only includes its intended IdeaVim cleanup.

### Task 3: Validate IdeaVim Discovery

**Files:**
- Verify: `~/.config/ideavim/ideavimrc`

**Interfaces:**
- Consumes: XDG configuration created in Task 1.
- Produces: confirmation that IntelliJ loads the intended file.

- [ ] **Step 1: Reload the configuration in IntelliJ**

In IntelliJ, use `:source ~/.config/ideavim/ideavimrc` in the IdeaVim command line, then restart the IDE once to exercise automatic XDG discovery.

Expected: no IdeaVim error notification.

- [ ] **Step 2: Exercise representative merged mappings**

In an open project, test `<Space>ff`, `<Space>fp`, `<Space>fg`, `<Space>e`, `<Space>sv`, and `<Space>dr`.

Expected: Goto File, Find in Path, Search Everywhere, Project tool window, vertical split, and Run respectively.

- [ ] **Step 3: Verify Git scope**

Run:

```bash
git -C "$HOME/.config/ideavim" status --short
git -C "$HOME/dotfiles-vim" status --short
git -C "$HOME/dotfiles-ldkhangg" status --short
```

Expected: only the intended IdeaVim additions/deletions and pre-existing unrelated changes are present.
