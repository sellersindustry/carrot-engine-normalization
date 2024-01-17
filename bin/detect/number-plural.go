package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)


// Plural Number - 80s, 1970s, 1s, 10s
func NumberPlural(buffer string) *token.Model {
	r, _    := regexp.Compile(`^(([0-9])+(s|S))`);
	match   := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.NumberPlural);
	}
	return nil;
}
