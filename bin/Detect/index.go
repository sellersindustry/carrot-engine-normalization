package Detect

import (
	"strings"

	"github.com/sellersindustry/normalization-tts/bin/Token"
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
		token := extractByPattern(buffer, pattern);
		if (token != nil) {
			return token
		}
	}
	return Token.NewGeneral(buffer[0:1], Token.Symbol)
}


func extractByPattern(buffer string, pattern *Pattern) *Token.Model {
	if (pattern.DetectWords != nil) {
		for _, word := range pattern.DetectWords {
			if strings.EqualFold(buffer, word) {
				return Token.NewGeneral(buffer, pattern.Class)
			}
		}
	}
	if (pattern.DetectRegexp != nil) {
		match := pattern.DetectRegexp.FindStringIndex(buffer)
		if (len(match) > 0) {
			text := buffer[match[0]:match[1]]
			if (len(text) > 0) {
				return Token.NewGeneral(text, pattern.Class)
			}
		}
	}
	return nil;
}

