# todo-open-git-sync

Standalone example sync plugin for `todo-open`.

This project lives outside the main `todo-open` repo and provides a minimal plugin binary named `todoopen-plugin-sync-git`.

## What it does

- Implements the runtime plugin handshake for a `sync` adapter named `git`.
- Exposes required capabilities: `pull`, `push`, `status`.
- Returns stubbed responses for `status`, `push`, and `pull` requests.

> This is an example scaffold (not a production git sync engine).

## Build

```bash
go build -o bin/todoopen-plugin-sync-git ./cmd/todoopen-plugin-sync-git
```

## Run (manual)

```bash
./bin/todoopen-plugin-sync-git
```

The plugin writes a JSON handshake line to stdout and then serves newline-delimited JSON requests on stdin.

## Use with todo-open

Add plugin registration to your workspace `.todoopen/meta.json`:

```json
{
  "workspace_version": 1,
  "schema_version": "todo.open.task.v1",
  "enabled_sync_adapters": ["noop", "git"],
  "adapter_plugins": [
    {
      "name": "git",
      "kind": "sync",
      "command": "/absolute/path/to/todo-open-git-sync/bin/todoopen-plugin-sync-git"
    }
  ]
}
```

Then inspect status with:

```bash
todoopen adapters --workspace /path/to/workspace --json
```
# todo-open-git-sync
