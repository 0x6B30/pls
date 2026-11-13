# historystore

Keeps the visit history in a single TSV file, one folder per line:

```
<absolute path>\t<visit count>\t<last visit, unix seconds>
```

- The file is created with mode `0600` and its folder with `0700`.
- Every write goes to a temp file that is then renamed over the old one, so readers never see a half-written file.
- Lines that fail to parse are skipped, not reported, so a damaged file never blocks a `cd`.
- Paths that aren't absolute, or that contain tabs or newlines, are refused with `ErrInvalidPath`.
- When two shells write at the same moment, the last write wins and one visit can be lost. That's fine for a ranking signal.

Size stays bounded through `history.Age` (see `pkg/domains/history`).
