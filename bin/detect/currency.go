package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)


func Currency(buffer string) *token.Model {
	r, _    := regexp.Compile(`^((\$)|(€)|(ƒ)|(¥)|(JP¥)|(CN¥))(\s)*[0-9]+(,[0-9]+)*(\.[0-9]+)?`);
	match   := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.Currency);
	}
	return nil;
}
