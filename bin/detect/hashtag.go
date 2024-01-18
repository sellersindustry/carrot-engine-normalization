package detect

import (
	"github.com/sellersindustry/normalization-tts/bin/token"
	"github.com/sellersindustry/normalization-tts/bin/utility"
)


func hashtag(buffer string) *token.Model {
	regexp := `^#[a-zA-Z0-9-_]*[a-zA-Z0-9]`;
	return utility.DetectTokenGeneral(buffer, regexp, token.Hashtag);
}
