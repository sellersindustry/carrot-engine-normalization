package Transform

import (
	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


func Process(tokens *[]Token.Model) {
	for index := range *tokens {
		transform(index, tokens)
	}
}


func transform(index int, tokens *[]Token.Model) {
	for _, pattern := range PATTERNS {
		if transformByPattern(index, tokens, pattern) {
			return
		}
	}
	if (*tokens)[index].Subclass == Token.None {
		(*tokens)[index].Text = Utility.Decamelize((*tokens)[index].Original);
	} else {
		(*tokens)[index].Text = (*tokens)[index].Original
	}
}


func transformByPattern(index int, tokens *[]Token.Model, pattern *Pattern) bool {
	if (&(*tokens)[index].IsInactive != nil && (*tokens)[index].IsInactive == true && (*tokens)[index].Class != Token.Space) {
		(*tokens)[index].Text = "";
		return true;
	}
	if (pattern.Class != "" && pattern.Class == (*tokens)[index].Class) {
		(*tokens)[index].Text = (pattern.Function)(index, tokens);
		return true;
	}
	if (pattern.Subclass != "" && pattern.Subclass == (*tokens)[index].Subclass) {
		(*tokens)[index].Text = (pattern.Function)(index, tokens);
		return true;
	}
	return false;
}

