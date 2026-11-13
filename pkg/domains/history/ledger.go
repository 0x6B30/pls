package history

import "time"

const (
	maxTotalVisits = 10_000
	agingFactor    = 0.9
	minVisits      = 1.0
)

func Visit(entries []Entry, path string, at time.Time) []Entry {
	for i := range entries {
		if entries[i].Path != path {
			continue
		}
		entries[i].Visits++
		entries[i].LastVisit = at
		return entries
	}
	return append(entries, Entry{Path: path, Visits: 1, LastVisit: at})
}

// Age keeps the history bounded: once total visits pass the budget, every
// entry shrinks and the ones that fade below one visit are forgotten.
func Age(entries []Entry) []Entry {
	if totalVisits(entries) <= maxTotalVisits {
		return entries
	}

	kept := make([]Entry, 0, len(entries))
	for _, entry := range entries {
		entry.Visits *= agingFactor
		if entry.Visits >= minVisits {
			kept = append(kept, entry)
		}
	}
	return kept
}

func totalVisits(entries []Entry) float64 {
	var total float64
	for _, entry := range entries {
		total += entry.Visits
	}
	return total
}
