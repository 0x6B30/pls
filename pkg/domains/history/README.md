# history

Rules for visited folders. Pure logic with no I/O.

**Frecency** = visits × a recency weight:

| Last visit    | Weight |
| ------------- | ------ |
| < 1 hour      | ×4     |
| < 1 day       | ×2     |
| < 1 week      | ×0.5   |
| older         | ×0.25  |

**Matching** checks only the last path component, ignoring case: `api` matches `~/work/API-server` but not `~/api/docs`.

**Aging**: once total visits pass 10,000, every count is multiplied by 0.9, and entries that drop below 1 are removed. This keeps the history file small without a hard entry limit.
