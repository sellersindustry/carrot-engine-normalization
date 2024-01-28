package Detect

import (
	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


func Process(buffer string) *[]Token.Model {
	var tokens []Token.Model
	for (len(buffer) > 0) {
		nextToken := getNextToken(buffer, PATTERNS);
		tokens     = append(tokens, *nextToken);
		buffer     = buffer[len(nextToken.Original):];
	}
	tokens = append(tokens, *Token.NewGeneral("", Token.Termination));
	return &tokens
}


func getNextToken(buffer string, patterns []*Pattern) *Token.Model {
	for _, pattern := range patterns {
		token := extractByPattern(buffer, pattern.Regexp, pattern.Class);
		if (token != nil) {
			return token
		}
	}
	return Token.NewGeneral(buffer[0:1], Token.Symbol)
}


func extractByPattern(buffer string, regex string, class Token.Class) *Token.Model {
	pattern := Utility.CompileRegex(regex);
	match   := pattern.FindStringIndex(buffer)
	if (len(match) > 0) {
		text := buffer[match[0]:match[1]]
		if (len(text) > 0) {
			return Token.NewGeneral(text, class)
		}
	}
	return nil;
}

