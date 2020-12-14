package logger

import "strings"

func logFieldName(prefixes []string) string {
	filtered := make([]string, 0, len(prefixes))
	for _, p := range prefixes {
		if p != "" {
			filtered = append(filtered, p)
		}
	}

	return strings.Join(filtered, "_")
}
