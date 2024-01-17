package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)


func URL(buffer string) *token.Model {
	r, _    := regexp.Compile(`^([A-Za-z0-9]+:\/\/)?([A-Za-z0-9]+(\.[A-Za-z0-9]+)+)(([\/\?])([\S]*[0-9A-Za-z\/])?)?`);
	match   := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.URL);
	}
	return nil;
}
