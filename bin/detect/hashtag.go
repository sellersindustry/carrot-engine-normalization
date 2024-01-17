package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)

func Hashtag(buffer string) *token.Model {
	r, _  := regexp.Compile(`^#[a-zA-Z0-9-_]*[a-zA-Z0-9]`);
	match := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.Hashtag);
	}
	return nil;
}
