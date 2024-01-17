package detect

import (
	"github.com/sellersindustry/normalization-tts/bin/token"
)

func Process(buffer string) *[]token.Model {
	var tokens []token.Model;
	for (len(buffer) > 0) {
		nextToken := nextToken(buffer);
		tokens     = append(tokens, *nextToken);
		buffer     = buffer[len(nextToken.Original):];
	}
	return &tokens;
}

func nextToken(buffer string) *token.Model {
	nextToken := Space(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = Control(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = Hashtag(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = Email(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = URL(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = Phone(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = Date(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = Time(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = Currency(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = Currency(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = NumberRoman(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = NumberPlural(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = NumberOrdinal(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	nextToken = NumberNominal(buffer);
	if (nextToken != nil) {
		return nextToken;
	}
	// number Cardinal - 2/3 and 1,324
	// word
	return token.New(buffer, token.Word);
}

