package Utility

import (
	"strings"
	"unicode"
)


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


func FilterNonEmptyStrings(input []string) []string {
	var result []string
	for _, str := range input {
		if len(str) > 0 {
			result = append(result, str)
		}
	}
	return result
}


func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}


func Decamelize(word string) string {
	if !ContainsOnlyLetters(word) || strings.ToUpper(word) == word {
		return word
	}
	var result strings.Builder
	for i, letter := range word {
		if i > 0 && unicode.IsUpper(letter) && unicode.IsLetter(rune(word[i-1])) {
			result.WriteRune(' ')
		}
		result.WriteRune(letter)
	}
	return result.String()
}


func ContainsOnlyLetters(word string) bool {
	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			return false
		}
	}
	return true
}

