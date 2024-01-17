package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)


// Ordinal Number - 1st, 1,324th
func NumberOrdinal(buffer string) *token.Model {
	r, _    := regexp.Compile(`^(([0-9])+(\,[0-9]{3})*(st|ST|nd|ND|rd|RD|th|TH))`);
	match   := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.New(buffer[match[0]:match[1]], token.NumberOrdinal);
	}
	return nil;
}
