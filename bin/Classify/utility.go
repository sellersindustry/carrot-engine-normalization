package Classify

import (
	"regexp"
	"strings"

	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


func isTokenMatch(index int, tokens *[]Token.Model, pattern *Pattern) bool {
	if !matchToken(index, tokens, &(*pattern).Current) {
		return false
	}
	if !matchAdjecentTokens(index, tokens, &(*pattern).Prefix, true) {
		return false
	}
	if !matchAdjecentTokens(index, tokens, &(*pattern).Suffix, false) {
		return false
	}
	//! IMPLEMENT SCAN BEFORE/AFTER
	return true
}


func matchAdjecentTokens(matchIndex int, tokens *[]Token.Model, prefixes *[]string, isBackwards bool) bool {
	if prefixes == nil {
		return true
	}
	ignoreSpaces := true
	offsetIndex := 1
	prefixIndex := 0
	for prefixIndex < len(*prefixes) {
		var index int
		if isBackwards {
			index = matchIndex - offsetIndex
		} else {
			index = matchIndex + offsetIndex
		}
		if index < 0 || index > len(*tokens)-1 {
			return false
		}
		if ignoreSpaces && (*tokens)[index].Class == Token.Space {
			offsetIndex++
			continue
		}
		if !matchToken(index, tokens, &[]string{(*prefixes)[prefixIndex]}) {
			return false
		}
		offsetIndex += 1
		prefixIndex += 1
	}
	return true
}


func matchToken(index int, tokens *[]Token.Model, keywords *[]string) bool {
	if keywords == nil {
		return true
	}
	buffer := strings.ToLower(*&(*tokens)[index].Original)
	for _, keyword := range *keywords {
		if (keyword == CLASS_NUMBER && Token.Number == *&(*tokens)[index].Class) {
			return true
		}
		if (keywordIsRegex(&keyword)) {
			if (keywordGetRegex(&keyword).Match([]byte(buffer))) {
				return true;
			}
		}
		if buffer == keyword {
			return true
		}
	}
	return false
}


func keywordIsRegex(keyword *string) bool {
	return keyword != nil && strings.HasPrefix(*keyword, "/") && strings.HasSuffix(*keyword, "/");
}


func keywordGetRegex(keyword *string) *regexp.Regexp {
	pattern := strings.TrimPrefix(strings.TrimSuffix(*keyword, "/"), "/");
	return Utility.CompileRegex(pattern);
}

