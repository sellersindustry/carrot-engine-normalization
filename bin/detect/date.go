package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)

func Date(buffer string) *token.Model {
	r, _  := regexp.Compile(`^[0-9]{1,4}[\.\/-][0-9]{1,2}[\.\/-]\d{1,4}`);
	match := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.Date);
	}
	return nil;
}
