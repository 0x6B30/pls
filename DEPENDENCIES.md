# External dependencies

Every third-party Go module imported by this repo is listed here so it can be checked for CVEs (for example with `govulncheck ./...`).

| Module | Version | Used by | Why |
| ------ | ------- | ------- | --- |
| _none_ | | | Standard library only |

Runtime tools called by the binary:

| Tool  | Used by          | Why                    |
| ----- | ---------------- | ---------------------- |
| `git` | `pkg/infra/gitcli` | cloning repositories |
