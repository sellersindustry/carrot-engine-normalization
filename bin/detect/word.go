package detect

import (
	"github.com/sellersindustry/normalization-tts/bin/token"
	"github.com/sellersindustry/normalization-tts/bin/utility"
)


func word(buffer string) *token.Model {
	return utility.DetectReturnFirstNonNullToken(buffer,
		wordGeneral,
	);
}


func wordGeneral(buffer string) *token.Model {
	regexp := `^[a-zA-Z]+(\-[a-zA-Z]+)*(\'[a-zA-Z]+)?`;
	return utility.DetectTokenGeneral(buffer, regexp, token.Word);
}

