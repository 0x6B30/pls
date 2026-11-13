# resolver

Turns what the user typed into an absolute folder path. It tries a chain of strategies in order, and the first one that handles the request wins.

| # | Strategy       | Handles when                               | Does                                  |
| - | -------------- | ------------------------------------------ | ------------------------------------- |
| 1 | `cloneRepo`    | target starts with `git@`, `https://`, ... | reuses the folder if it exists, otherwise clones |
| 2 | `forcedCreate` | `-n` was passed                            | `mkdir -p`                            |
| 3 | `existingDir`  | target is an existing folder               | returns it                            |
| 4 | `historyMatch` | target has no `/` and a visited folder matches | returns the best match            |
| 5 | `createDir`    | always (fallback)                          | `mkdir -p`                            |

A target with a `/` in it is treated as a path, never a search, so `pls a/b/c` always creates `a/b/c` rather than jumping to some other `c`.

## Adding behaviour

1. Create a file with a type that implements `strategy`. Return `handled=false` to pass the request on.
2. Insert it into the list in `New` (`resolver.go`) at the right position.

Existing strategies don't need to change.
