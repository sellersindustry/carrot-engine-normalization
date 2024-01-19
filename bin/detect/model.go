package detect

import "github.com/sellersindustry/normalization-tts/bin/token"


type Pattern struct {
	Regexp   string
	Class    token.Class
}

