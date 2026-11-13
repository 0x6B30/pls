# resolver

Turns what the user typed after `pls` into an absolute folder path. It tries a chain of strategies in order, and the first one that handles the request wins.

| # | Strategy      | Handles when                               | Does                                  |
| - | ------------- | ------------------------------------------ | ------------------------------------- |
| 1 | `cloneRepo`   | target starts with `git@`, `https://`, ... | reuses the folder if it exists, otherwise clones |
| 2 | `existingDir` | target is an existing folder               | returns it                            |
| 3 | `createDir`   | always (fallback)                          | `mkdir -p`                            |

The resolver never looks at visit history. Jumping to a visited folder is a separate command (`pls jump`) that uses `pkg/services/matcher`.

## Adding behaviour

1. Create a file with a type that implements `strategy`. Return `handled=false` to pass the request on.
2. Insert it into the list in `New` (`resolver.go`) at the right position.

Existing strategies don't need to change.
