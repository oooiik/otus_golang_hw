package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	if len(s) == 0 {
		return nil
	}

	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", "")

	otchot := make(map[string]int)
	for _, slova := range strings.Fields(s) {
		if len(slova) == 0 {
			continue
		}
		otchot[slova]++
	}

	keys := make([]string, 0, 10)

	for key := range otchot {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if otchot[keys[i]] == otchot[keys[j]] {
			return strings.Compare(keys[i], keys[j]) <= 0
		}
		return otchot[keys[i]] > otchot[keys[j]]
	})

	if len(keys) <= 10 {
		return keys
	}

	return keys[:10]
}
