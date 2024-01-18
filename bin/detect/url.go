package detect

import (
	"github.com/sellersindustry/normalization-tts/bin/token"
	"github.com/sellersindustry/normalization-tts/bin/utility"
)


func url(buffer string) *token.Model {
	regexp := `^([A-Za-z0-9]+:\/\/)?([A-Za-z0-9]+(\.[A-Za-z0-9]+)+)(([\/\?])([\S]*[0-9A-Za-z\/])?)?`;
	return utility.DetectTokenGeneral(buffer, regexp, token.URL);
}
