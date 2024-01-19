package detect

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)


func getNextToken(buffer string, patterns []*Pattern) *token.Model {
	for _, pattern := range patterns {
		token := extractByPattern(buffer, pattern.Regexp, pattern.Class);
		if (token != nil) {
			return token
		}
	}
	return token.NewGeneral(buffer[0:1], token.Symbol)
}


func extractByPattern(buffer string, regex string, class token.Class) *token.Model {
	pattern := compileRegex(regex)
	match   := pattern.FindStringIndex(buffer)
	if (len(match) > 0) {
		text := buffer[match[0]:match[1]]
		if (len(text) > 0) {
			return token.NewGeneral(text, class)
		}
	}
	return nil;
}


func compileRegex(regex string) *regexp.Regexp {
	r, err := regexp.Compile(regex)
	if (err != nil) {
		fmt.Printf("Error: Failed to compile regular expression: \"%s\"\n", regex);
		os.Exit(1);
		return nil;
	}
	return r;
}

