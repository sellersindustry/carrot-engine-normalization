package detect

import (
	"github.com/sellersindustry/normalization-tts/bin/token"
)


func Process(buffer string) *[]token.Model {
	var tokens []token.Model
	for (len(buffer) > 0) {
		nextToken := getNextToken(buffer, patterns);
		tokens     = append(tokens, *nextToken);
		buffer     = buffer[len(nextToken.Original):];
	}
	return &tokens
}

