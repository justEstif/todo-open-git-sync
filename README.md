# todo-open-git-sync

Minimal external sync plugin example for `todo-open`.

It provides one plugin binary: `todoopen-plugin-sync-git`.

## Install globally (mise)

Use mise's Go backend to install the command from module source:

```bash
mise use -g go:github.com/justEstif/todo-open-git-sync/cmd/todoopen-plugin-sync-git@latest
```

That installs `todoopen-plugin-sync-git` onto your PATH via mise shims.

## Register in a workspace

Edit `.todoopen/meta.json`:

```json
{
  "workspace_version": 1,
  "schema_version": "todo.open.task.v1",
  "enabled_sync_adapters": ["noop", "git"],
  "adapter_plugins": [
    {"name": "git", "kind": "sync", "command": "todoopen-plugin-sync-git"}
  ]
}
```

## Verify

```bash
todoopen adapters --workspace /path/to/workspace --json
```

Expected: a `sync` adapter named `git` with `source: plugin`.

---

This is a scaffold for plugin wiring and handshake behavior, not a production git sync engine.
