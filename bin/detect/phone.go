package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)


func Phone(buffer string) *token.Model {
	r, _  := regexp.Compile(`^([\+][\s]?[0-9]{1,3}[\s]?)?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4}`);
	match := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.Phone);
	}
	return nil;
}
