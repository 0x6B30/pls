package matcher

import (
	"errors"
	"testing"
	"time"

	"github.com/0x6B30/pls/pkg/domains/history"
	"github.com/0x6B30/pls/pkg/infra/dirs"
	"github.com/0x6B30/pls/pkg/infra/historystore"
)

var now = time.Unix(1_700_000_000, 0)

func newTestMatcher(entries []history.Entry, missing ...string) Matcher {
	gone := map[string]bool{}
	for _, path := range missing {
		gone[path] = true
	}
	return New(
		&historystore.Mock{EntriesFunc: func() ([]history.Entry, error) { return entries, nil }},
		&dirs.Mock{ExistsFunc: func(path string) bool { return !gone[path] }},
		func() time.Time { return now },
	)
}

func TestBestPicksHighestFrecency(t *testing.T) {
	m := newTestMatcher([]history.Entry{
		{Path: "/old/api", Visits: 10, LastVisit: now.Add(-30 * 24 * time.Hour)},
		{Path: "/fresh/api", Visits: 2, LastVisit: now.Add(-time.Minute)},
	})

	got, err := m.Best("api", "/")

	if err != nil || got != "/fresh/api" {
		t.Fatalf("got %q, %v; want /fresh/api", got, err)
	}
}

func TestBestSkipsCurrentAndDeletedFolders(t *testing.T) {
	m := newTestMatcher([]history.Entry{
		{Path: "/here/api", Visits: 9, LastVisit: now},
		{Path: "/deleted/api", Visits: 5, LastVisit: now},
		{Path: "/other/api", Visits: 1, LastVisit: now},
	}, "/deleted/api")

	got, err := m.Best("api", "/here/api/")

	if err != nil || got != "/other/api" {
		t.Fatalf("got %q, %v; want /other/api", got, err)
	}
}

func TestBestWithoutMatch(t *testing.T) {
	m := newTestMatcher([]history.Entry{{Path: "/a/web", Visits: 1, LastVisit: now}})

	if _, err := m.Best("api", "/"); !errors.Is(err, ErrNoMatch) {
		t.Fatalf("got %v, want ErrNoMatch", err)
	}
}

func TestBestPropagatesHistoryErrors(t *testing.T) {
	boom := errors.New("boom")
	m := New(
		&historystore.Mock{EntriesFunc: func() ([]history.Entry, error) { return nil, boom }},
		&dirs.Mock{},
		time.Now,
	)

	if _, err := m.Best("api", "/"); !errors.Is(err, boom) {
		t.Fatalf("got %v, want boom", err)
	}
}
