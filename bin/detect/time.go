package detect

import (
	"github.com/sellersindustry/normalization-tts/bin/token"
	"github.com/sellersindustry/normalization-tts/bin/utility"
)


func time(buffer string) *token.Model {
	regexp := `^[0-9]{1,2}:[0-9]{2}\b`;
	return utility.DetectTokenGeneral(buffer, regexp, token.Time);
}
