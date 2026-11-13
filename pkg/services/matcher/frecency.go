package matcher

import (
	"cmp"
	"path/filepath"
	"slices"
	"time"

	"github.com/0x6B30/pls/pkg/domains/history"
	"github.com/0x6B30/pls/pkg/infra/dirs"
	"github.com/0x6B30/pls/pkg/infra/historystore"
)

type frecencyMatcher struct {
	history historystore.Reader
	dirs    dirs.Checker
	now     func() time.Time
}

func (m *frecencyMatcher) Best(query, exclude string) (string, error) {
	entries, err := m.history.Entries()
	if err != nil {
		return "", err
	}

	candidates := matching(entries, query, filepath.Clean(exclude))
	sortByFrecency(candidates, m.now())

	for _, candidate := range candidates {
		if m.dirs.Exists(candidate.Path) {
			return candidate.Path, nil
		}
	}
	return "", ErrNoMatch
}

func matching(entries []history.Entry, query, exclude string) []history.Entry {
	var found []history.Entry
	for _, entry := range entries {
		if entry.Path != exclude && entry.Matches(query) {
			found = append(found, entry)
		}
	}
	return found
}

func sortByFrecency(entries []history.Entry, now time.Time) {
	slices.SortStableFunc(entries, func(a, b history.Entry) int {
		return cmp.Compare(b.Frecency(now), a.Frecency(now))
	})
}
