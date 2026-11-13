package history

import (
	"testing"
	"time"
)

func TestFrecencyFavoursRecentVisits(t *testing.T) {
	now := time.Now()
	recent := Entry{Visits: 1, LastVisit: now.Add(-time.Minute)}
	old := Entry{Visits: 1, LastVisit: now.Add(-30 * 24 * time.Hour)}

	if recent.Frecency(now) <= old.Frecency(now) {
		t.Fatalf("recent entry should outrank old one")
	}
}

func TestMatchesOnlyLastComponentIgnoringCase(t *testing.T) {
	entry := Entry{Path: "/home/me/Work/API-Server"}

	if !entry.Matches("api") {
		t.Fatalf("expected match on last component")
	}
	if entry.Matches("work") {
		t.Fatalf("parent folders must not match")
	}
}

func TestVisitBumpsExistingEntry(t *testing.T) {
	at := time.Unix(100, 0)
	entries := []Entry{{Path: "/a", Visits: 2, LastVisit: time.Unix(1, 0)}}

	entries = Visit(entries, "/a", at)

	if len(entries) != 1 || entries[0].Visits != 3 || !entries[0].LastVisit.Equal(at) {
		t.Fatalf("unexpected entries: %+v", entries)
	}
}

func TestVisitAppendsNewEntry(t *testing.T) {
	entries := Visit(nil, "/a", time.Unix(1, 0))

	if len(entries) != 1 || entries[0].Visits != 1 {
		t.Fatalf("unexpected entries: %+v", entries)
	}
}

func TestAgeLeavesSmallHistoryAlone(t *testing.T) {
	entries := []Entry{{Path: "/a", Visits: 5}}

	if got := Age(entries); got[0].Visits != 5 {
		t.Fatalf("expected untouched entry, got %+v", got)
	}
}

func TestAgeShrinksAndForgetsFadedEntries(t *testing.T) {
	entries := []Entry{
		{Path: "/busy", Visits: maxTotalVisits},
		{Path: "/faded", Visits: 1},
	}

	got := Age(entries)

	if len(got) != 1 || got[0].Path != "/busy" || got[0].Visits >= maxTotalVisits {
		t.Fatalf("unexpected aged entries: %+v", got)
	}
}
