# pls

Create, clone, or jump to a folder and `cd` into it, in one command.

```sh
pls a/b/c                          # mkdir -p a/b/c && cd a/b/c
pls git@github.com:me/tool.git     # git clone ... && cd tool   (just cd if already cloned)
pls https://github.com/me/tool x   # clone into ./x
pls api                            # jump to the most used/recent visited folder named like "api"
pls -n api                         # always create ./api, never jump
```

## Install

```sh
./install.sh            # detects your shell from $SHELL
./install.sh bash       # or pick one: zsh | bash
PLS_VERSION=v0.1.0 ./install.sh
```

The script runs `go install github.com/0x6B30/pls/cmd/pls-cli@latest` and adds this to your rc file:

```sh
eval "$(pls-cli init zsh)"
```

## How it works

A program can't change its parent shell's directory, so the work is split:

- `pls-cli`, the Go binary, decides where to go and prints the absolute path.
- `pls`, a small shell function from `pls-cli init`, runs the binary and `cd`s into what it printed. A cd hook calls `pls-cli add "$PWD"` on every directory change, which is how partial-name jumps learn.

How a target is resolved: [pkg/services/resolver/README.md](pkg/services/resolver/README.md).

## Layout

```
cmd/pls-cli/          CLI entrypoint: wiring and subcommands (jump, add, init)
pkg/domains/history   visit model, frecency, aging rules
pkg/utils/giturl      recognising git URLs, deriving the repo folder name
pkg/infra/dirs        filesystem checks and mkdir
pkg/infra/gitcli      runs git clone
pkg/infra/historystore  visit history on disk
pkg/services/matcher  picks the best visited folder for a partial name
pkg/services/resolver turns a target into a folder (clone / existing / history / create)
pkg/adapters/shellinit  shell wrapper scripts (zsh, bash)
```

Each package exposes an interface, its models and errors, a `New` constructor and a `Mock`. Everything else is private.

## Configuration

| Variable           | Default                                  |
| ------------------ | ---------------------------------------- |
| `PLS_HISTORY_FILE` | `$XDG_DATA_HOME/pls/history.tsv`, else `~/.local/share/pls/history.tsv` |

## Development

```sh
go test ./...
go build -o /tmp/pls-cli ./cmd/pls-cli
```

External dependencies are tracked in [DEPENDENCIES.md](DEPENDENCIES.md).
