package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)

func Space(buffer string) *token.Model {
	r, _  := regexp.Compile(`^[\n\s\t]+`);
	match := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.Space);
	}
	return nil;
}
