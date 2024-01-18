package detect

import (
	"github.com/sellersindustry/normalization-tts/bin/token"
	"github.com/sellersindustry/normalization-tts/bin/utility"
)


func space(buffer string) *token.Model {
	regexp := `^[\n\s\t]+`;
	return utility.DetectTokenGeneral(buffer, regexp, token.Space);
}
