package Wordify

import (
	"strconv"
	"strings"
)

func Date(text string) string {
	parts := strings.FieldsFunc(text, func(r rune) bool {
		return r == '-' || r == '.' || r == '/'
	})

	var normalizedDate []string

	for i, part := range parts {
		if _, err := strconv.Atoi(part); err == nil {
			switch i {
			case 2:
				if len(part) == 4 {
					normalizedDate = append(normalizedDate, NumberYear(part))
				} else if len(part) == 2 && strings.HasPrefix(part, "0") {
					normalizedDate = append(normalizedDate, "o", Number(part[1:]))
				} else {
					normalizedDate = append(normalizedDate, Number(part))
				}
			default:
				normalizedDate = append(normalizedDate, Number(part))
			}
		} else {
			normalizedDate = append(normalizedDate, part)
		}
	}

	return strings.TrimSpace(strings.Join(normalizedDate, " "))
}
