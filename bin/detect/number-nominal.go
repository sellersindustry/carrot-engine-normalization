package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)


// Nominal Numbers - Each Digit Spoken Seperatly
func NumberNominal(buffer string) *token.Model {
	r, _    := regexp.Compile(`^[0-9]{5,}`);
	match   := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.NumberNominal);
	}
	return nil;
}
