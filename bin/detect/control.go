package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)

func Control(buffer string) *token.Model {
	r, _  := regexp.Compile(`^(\{\{)[\w\=\-\s]+(\}\})`);
	match := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.Control);
	}
	return nil;
}
