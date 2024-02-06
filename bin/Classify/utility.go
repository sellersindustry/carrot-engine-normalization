package Classify

import (
	"regexp"
	"strings"

	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


func isTokenMatch(index int, tokens *[]Token.Model, pattern *Pattern) bool {
	if !detectCurrent(index, tokens, pattern) {
		return false
	}
	if !matchAdjecentTokens(index, tokens, &(*pattern).HasPrefix, true) {
		return false
	}
	if !matchAdjecentTokens(index, tokens, &(*pattern).HasSuffix, false) {
		return false
	}
	if !matchRangeTokens(index, tokens, (*pattern).ScanBefore, true) {
		return false
	}
	if !matchRangeTokens(index, tokens, (*pattern).ScanAfter, false) {
		return false
	}
	return true
}


func detectCurrent(index int, tokens *[]Token.Model, pattern *Pattern) bool {
	if detectCurrentBy(index, tokens, pattern) {
		if (*pattern).Filter != nil {
			return (*pattern).Filter((*tokens)[index].Original)
		}
		return true
	} else {
		return false
	}
}


func detectCurrentBy(index int, tokens *[]Token.Model, pattern *Pattern) bool {
	if pattern.CurrentByRegexp != nil {
		return pattern.CurrentByRegexp.Match([]byte(*&(*tokens)[index].Original));
	} else if string(pattern.CurrentByClass) != "" {
		return pattern.CurrentByClass == (*tokens)[index].Class
	} else if string(pattern.CurrentBySubclass) != "" {
		return pattern.CurrentBySubclass == (*tokens)[index].Subclass
	} else if pattern.CurrentByWords != nil {
		return Utility.Contains(pattern.CurrentByWords, (*tokens)[index].Original)
	}
	return false;
}


var IGNORE_SPACES = "{{IGNORE_SPACES}}"

func matchAdjecentTokens(tokenIndex int, tokens *[]Token.Model, prefixes *[]string, isPrefix bool) bool {
	if prefixes == nil {
		return true
	}
	if (len(*prefixes) == 0) {
		return true;
	}
	ignoreSpaces := (*prefixes)[0] == IGNORE_SPACES;
	offsetIndex  := 1
	prefixIndex  := 0
	for prefixIndex < len(*prefixes) {
		var _tokenIndex int
		var _prefixIndex int
		if isPrefix {
			_tokenIndex  = tokenIndex - offsetIndex
			_prefixIndex = len(*prefixes) - prefixIndex - 1
		} else {
			_tokenIndex  = tokenIndex + offsetIndex
			_prefixIndex = prefixIndex
		}
		if ((*prefixes)[_prefixIndex] == IGNORE_SPACES) {
			prefixIndex += 1;
			continue;
		}
		if _tokenIndex < 0 || _tokenIndex > len(*tokens)-1 {
			return false
		}
		if ignoreSpaces && (*tokens)[_tokenIndex].Class == Token.Space {
			offsetIndex += 1;
			continue
		}
		if !matchToken(_tokenIndex, tokens, &[]string{(*prefixes)[_prefixIndex]}) {
			return false
		}
		offsetIndex += 1
		prefixIndex += 1
	}
	return true
}


func matchRangeTokens(tokenIndex int, tokens *[]Token.Model, pattern *PatternScan, isPrefix bool) bool {
	if (pattern == nil) {
		return true;
	}
	ignoreSpaces := pattern.IgnoreSpaces;
	offsetIndex  := 1
	rangeIndex   := 0
	for rangeIndex < (*pattern).Range {
		var _tokenIndex int
		if isPrefix {
			_tokenIndex = tokenIndex - offsetIndex
		} else {
			_tokenIndex = tokenIndex + offsetIndex
		}
		if _tokenIndex < 0 || _tokenIndex > len(*tokens)-1 {
			if len((*pattern).Exists) == 0 {
				return true;
			}
			return false
		}
		if ignoreSpaces && (*tokens)[_tokenIndex].Class == Token.Space {
			offsetIndex += 1;
			continue
		}
		if &(*pattern).NotExists != nil {
			if matchToken(_tokenIndex, tokens, &(*pattern).NotExists) {
				return false
			}
		}
		if &(*pattern).Exists != nil {
			if matchToken(_tokenIndex, tokens, &(*pattern).Exists) {
				return true;
			}
		}
		rangeIndex  += 1;
		offsetIndex += 1;
	}
	if len((*pattern).Exists) == 0 {
		return true;
	}
	return false;
}



func matchToken(index int, tokens *[]Token.Model, keywords *[]string) bool {
	if keywords == nil {
		return true
	}
	buffer := strings.ToLower(*&(*tokens)[index].Original)
	for _, keyword := range *keywords {
		if (isKeywordClass(&keyword)) {
			return (*tokens)[index].Class == Token.Class(keyword)
		}
		if (isKeywordSubclass(&keyword)) {
			return (*tokens)[index].Subclass == Token.Subclass(keyword)
		}
		if (isKeywordRegex(&keyword)) {
			if (strings.HasSuffix(keyword, "/i")) {
				if (keywordGetRegex(&keyword).Match([]byte(*&(*tokens)[index].Original))) {
					return true;
				}
			} else {
				if (keywordGetRegex(&keyword).Match([]byte(buffer))) {
					return true;
				}
			}
		}
		if buffer == keyword {
			return true
		}
	}
	return false
}


func isKeywordRegex(keyword *string) bool {
	return keyword != nil && len(*keyword) > 2 && strings.HasPrefix(*keyword, "/") && (strings.HasSuffix(*keyword, "/") || strings.HasSuffix(*keyword, "/i"));
}


func keywordGetRegex(keyword *string) *regexp.Regexp {
	pattern := strings.TrimSuffix(strings.TrimSuffix(*keyword, "/i"), "/") + "/";
	return Utility.CompileRegex(pattern);
}


func isKeywordClass(keyword *string) bool {
	return keyword != nil && strings.HasPrefix(*keyword, "{{CLASS:") && strings.HasSuffix(*keyword, "}}");
}


func isKeywordSubclass(keyword *string) bool {
	return keyword != nil && strings.HasPrefix(*keyword, "{{SUBCLASS:") && strings.HasSuffix(*keyword, "}}");
}
