package history

import (
	"path/filepath"
	"strings"
	"time"
)

type Entry struct {
	Path      string
	Visits    float64
	LastVisit time.Time
}

func (e Entry) Frecency(now time.Time) float64 {
	age := now.Sub(e.LastVisit)
	switch {
	case age < time.Hour:
		return e.Visits * 4
	case age < 24*time.Hour:
		return e.Visits * 2
	case age < 7*24*time.Hour:
		return e.Visits / 2
	default:
		return e.Visits / 4
	}
}

// Matches compares against the last path component only, so "api" finds
// ~/work/api but not ~/api-docs/notes.
func (e Entry) Matches(query string) bool {
	name := strings.ToLower(filepath.Base(e.Path))
	return strings.Contains(name, strings.ToLower(query))
}
