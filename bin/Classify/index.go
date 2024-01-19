package Classify

import (
	"github.com/sellersindustry/normalization-tts/bin/Token"
)


func Process(tokens *[]Token.Model) {
	for index := range *tokens {
		classify(index, tokens);
	}
}


func classify(index int, tokens *[]Token.Model) {
	for _, pattern := range PATTERNS {
		if (classifyByPattern(index, tokens, pattern)) {
			return
		}
	}
}


func classifyByPattern(index int, tokens *[]Token.Model, pattern *Pattern) bool {
	if (&(*tokens)[index].IsSilent != nil && (*tokens)[index].IsSilent == true) {
		return false;
	}
	if (!isTokenMatch(index, tokens, pattern)) {
		return false;
	}
	if (&pattern.SetSubclassTo != nil) {
		(*tokens)[index].Subclass = pattern.SetSubclassTo;
	}
	if (&pattern.SetIsSilentTo != nil) {
		(*tokens)[index].IsSilent = pattern.SetIsSilentTo;
	}
	return true;
}


