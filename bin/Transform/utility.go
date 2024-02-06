package Transform

import (
	"github.com/sellersindustry/normalization-tts/bin/Token"
)



func isBefore(index int, tokens *[]Token.Model, distance int, targets []string) (bool, string) {
	for i := index; i >= 0 && index - i <= distance; i-- {
		for _, target := range targets {
			if (*tokens)[i].Original == target {
				return true, (*tokens)[i].Original
			}
		}
	}
	return false, ""
}


func isAfter(index int, tokens *[]Token.Model, distance int, targets []string) (bool, string) {
	for i := index; i < len(*tokens) && i - index <= distance; i++ {
		for _, target := range targets {
			if (*tokens)[i].Original == target {
				return true, (*tokens)[i].Original
			}
		}
	}
	return false, ""
}

