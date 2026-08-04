# IdeaVim Consolidation Design

## Goal

Keep one physical IdeaVim configuration file in the standard XDG location that
IdeaVim reads directly, while separating it from the existing Neovim dotfiles.

## Repository And Configuration Location

- `~/.config/ideavim/` becomes a standalone Git repository.
- `~/.config/ideavim/ideavimrc` is the only IdeaVim configuration file.
- Remove `~/.ideavimrc`; IdeaVim loads the XDG file directly.
- No symlinks, copied configuration files, or sync scripts are used.

## Merge Policy

The current `~/.ideavimrc` is the base because it is the configuration
currently used by IntelliJ and contains the most complete mapping set.

Add non-conflicting aliases from `~/dotfiles-vim/.ideavimrc`:

- File/navigation aliases: `<leader>e`, `<leader>sf`, `<leader>sg`, and
  `<leader>ss`.
- Formatting aliases: `<leader>fm` and `<leader>oi`.
- Window aliases: `<leader>sv`, `<leader>sh`, `<leader>se`, and
  `<leader>sx`.
- Diagnostic alias: `<leader>d`.

Do not add the old `<leader>fg` mapping because the current configuration uses
that key for Search Everywhere; Find in Path remains available at
`<leader>fp`.

## Existing Repositories

- Remove the duplicate `.ideavimrc` and IdeaVim copy/sync steps from
  `~/dotfiles-vim`, while preserving its Neovim configuration.
- Remove `home/.ideavimrc` from `~/dotfiles-ldkhangg`; its Hyprland
  IdeaVim cheatsheet stays there because it is desktop integration rather than
  IdeaVim configuration.

## Verification

- Confirm `~/.ideavimrc` no longer exists.
- Confirm `~/.config/ideavim/ideavimrc` exists and is the only IdeaVim config
  tracked by Git.
- Check every retained mapping is unique and every legacy alias resolves to the
  intended IdeaVim/IntelliJ action.
- Reload IdeaVim in IntelliJ and verify the configuration is loaded without
  errors.
