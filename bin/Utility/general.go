package Utility

import "strings"


func KeywordMapping(buffer string, keys []string, values []string, caseSensitive bool) string {
	for index, key := range keys {
		if caseSensitive {
			if strings.EqualFold(key, buffer) {
				return values[index];
			}
		} else {
			if key == buffer {
				return values[index];
			}
		}
	}
	return buffer;
}

