package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)

func Time(buffer string) *token.Model {
	r, _  := regexp.Compile(`^[0-9]{1,2}:[0-9]{2}`);
	match := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.Time);
	}
	return nil;
}
