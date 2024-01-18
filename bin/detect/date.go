package detect

import (
	"github.com/sellersindustry/normalization-tts/bin/token"
	"github.com/sellersindustry/normalization-tts/bin/utility"
)


func date(buffer string) *token.Model {
	regexp := `^[0-9]{1,4}[\.\/-][0-9]{1,2}[\.\/-]\d{1,4}\b`;
	return utility.DetectTokenGeneral(buffer, regexp, token.Date);
}
