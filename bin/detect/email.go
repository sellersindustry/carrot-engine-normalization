package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)


func Email(buffer string) *token.Model {
	r, _  := regexp.Compile(`^[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,64}`);
	match := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.Email);
	}
	return nil;
}
