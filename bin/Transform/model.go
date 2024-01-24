package Transform

import "github.com/sellersindustry/normalization-tts/bin/Token"

type Pattern struct {
	Class       Token.Class
	Subclass    Token.Subclass
	Function    func (index int, tokens *[]Token.Model) string
}

