package detect

import (
	"github.com/sellersindustry/normalization-tts/bin/token"
	"github.com/sellersindustry/normalization-tts/bin/utility"
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
	_token := utility.DetectReturnFirstNonNullToken(buffer,
		space,
		control,
		hashtag,
		email,
		url,
		phone,
		date,
		time,
		number,
		word,
	);
	if (_token == nil) {
		return token.NewGeneral(buffer[0:1], token.Symbol);
	} else {
		return _token
	}
}

