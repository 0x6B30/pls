package historystore

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestEntriesOnMissingFileIsEmpty(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "missing.tsv"))

	entries, err := store.Entries()

	if err != nil || len(entries) != 0 {
		t.Fatalf("got %v, %v; want empty history", entries, err)
	}
}

func TestRecordRoundTrips(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "nested", "history.tsv"))
	at := time.Unix(1_700_000_000, 0)

	for range 2 {
		if err := store.Record("/home/me/project/", at); err != nil {
			t.Fatalf("record: %v", err)
		}
	}
	entries, err := store.Entries()

	if err != nil || len(entries) != 1 {
		t.Fatalf("got %v, %v; want one entry", entries, err)
	}
	got := entries[0]
	if got.Path != "/home/me/project" || got.Visits != 2 || !got.LastVisit.Equal(at) {
		t.Fatalf("unexpected entry %+v", got)
	}
}

func TestRecordRejectsUnsafePaths(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "history.tsv"))

	for _, path := range []string{"relative/dir", "/with\ttab", "/with\nnewline"} {
		if err := store.Record(path, time.Now()); !errors.Is(err, ErrInvalidPath) {
			t.Errorf("Record(%q) = %v, want ErrInvalidPath", path, err)
		}
	}
}

func TestHistoryFileIsPrivate(t *testing.T) {
	file := filepath.Join(t.TempDir(), "history.tsv")
	if err := New(file).Record("/a", time.Now()); err != nil {
		t.Fatalf("record: %v", err)
	}

	info, err := os.Stat(file)
	if err != nil {
		t.Fatalf("stat: %v", err)
	}
	if perm := info.Mode().Perm(); perm != 0o600 {
		t.Fatalf("history file permissions = %o, want 600", perm)
	}
}

func TestDecodeSkipsMalformedLines(t *testing.T) {
	file := filepath.Join(t.TempDir(), "history.tsv")
	content := "/good\t3\t100\ngarbage\n/nan\tNaN\t1\n/neg\t-1\t1\n/badtime\t1\tx\n"
	if err := os.WriteFile(file, []byte(content), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}

	entries, err := New(file).Entries()

	if err != nil || len(entries) != 1 || entries[0].Path != "/good" {
		t.Fatalf("got %+v, %v; want only /good", entries, err)
	}
}
