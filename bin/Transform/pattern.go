package Transform

import (
	"strings"

	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
	"github.com/sellersindustry/normalization-tts/bin/Wordify"
)

var PATTERNS = []*Pattern{
	{
		Subclass: Token.NumberPlural,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.NumberPlural((*tokens)[index].Original)
		},
	}, {
		Subclass: Token.NumberOrdinal,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.NumberOrdinal((*tokens)[index].Original)
		},
	}, {
		Subclass: Token.NumberNominal,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.NumberNominal((*tokens)[index].Original)
		},
	}, {
		Subclass: Token.NumberCurrency,
		Function: func(index int, tokens *[]Token.Model) string {
			if index != 0 {
				if (*tokens)[index-1].Original == "$" {
					return Wordify.NumberCurrency((*tokens)[index].Original, "$")
				}
				if (*tokens)[index-1].Original == "€" {
					return Wordify.NumberCurrency((*tokens)[index].Original, "€")
				}
			}
			return Wordify.NumberCurrency((*tokens)[index].Original, "")
		},
	}, {
		Subclass: Token.NumberYear,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.NumberYear((*tokens)[index].Original)
		},
	}, {
		Subclass: Token.Unit,
		Function: func(index int, tokens *[]Token.Model) string {
			keys, values := Utility.GetWordset("./bin/wordsets/units.txt")
			return Utility.KeywordMapping((*tokens)[index].Original, keys, values, true)
		},
	}, {
		Subclass: Token.Scale,
		Function: func(index int, tokens *[]Token.Model) string {
			keys, values := Utility.GetWordset("./bin/wordsets/scales.txt")
			return Utility.KeywordMapping((*tokens)[index].Original, keys, values, false)
		},
	}, {
		Subclass: Token.Range,
		Function: func(index int, tokens *[]Token.Model) string {
			return "to"
		},
	}, {
		Subclass: Token.Per,
		Function: func(index int, tokens *[]Token.Model) string {
			return "per"
		},
	}, {
		Subclass: Token.QuoteStart,
		Function: func(index int, tokens *[]Token.Model) string {
			return "quote"
		},
	}, {
		Subclass: Token.QuoteStartShort,
		Function: func(index int, tokens *[]Token.Model) string {
			return "quote on quote"
		},
	}, {
		Subclass: Token.QuoteEnd,
		Function: func(index int, tokens *[]Token.Model) string {
			return "end quote"
		},
	}, {
		Subclass: Token.MathPrefix,
		Function: func(index int, tokens *[]Token.Model) string {
			if (*tokens)[index].Original == "-" {
				return "negative"
			}
			return "positive"
		},
	}, {
		Subclass: Token.MathOperation,
		Function: func(index int, tokens *[]Token.Model) string {
			keys, values := Utility.GetWordset("./bin/wordsets/operations.txt")
			return Utility.KeywordMapping((*tokens)[index].Original, keys, values, false)
		},
	}, {
		Class: Token.Date,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.Date((*tokens)[index].Original)
		},
	}, {
		Class: Token.Time,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.Time((*tokens)[index].Original)
		},
	}, {
		Class: Token.Number,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.Number((*tokens)[index].Original)
		},
	}, {
		Class: Token.Phone,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.Phone((*tokens)[index].Original)
		},
	}, {
		Subclass: Token.RomanNumeralPossessive,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.RomanNumeralPossessive((*tokens)[index].Original)
		},
	}, {
		Subclass: Token.RomanNumeral,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.RomanNumeral((*tokens)[index].Original)
		},
	}, {
		Subclass: Token.Initialism,
		Function: func(index int, tokens *[]Token.Model) string {
			text := (*tokens)[index].Original
			if strings.HasSuffix((*tokens)[index].Original, "s") {
				text = strings.TrimSuffix(text, "s")
			}
			text = strings.ReplaceAll(text, ".", "")

			buffer := strings.Join(strings.Split(text, ""), " ")
			if strings.HasSuffix((*tokens)[index].Original, "s") {
				buffer += "s"
			}
			return buffer
		},
	}, {
		Class: Token.Special,
		Function: func(index int, tokens *[]Token.Model) string {
			return Wordify.Special((*tokens)[index].Original)
		},
	}, {
		Subclass: Token.NewParagraph,
		Function: func(index int, tokens *[]Token.Model) string {
			return "{{PAUSE}}"
		},
	}, {
		Subclass: Token.Punctuation,
		Function: func(index int, tokens *[]Token.Model) string {
			if (*tokens)[index].Original == "(" {
				return ","
			}
			if (*tokens)[index].Original == ")" {
				return ","
			}
			return (*tokens)[index].Original
		},
	}, {
		Class: Token.Symbol,
		Function: func(index int, tokens *[]Token.Model) string {
			keys, values := Utility.GetWordset("./bin/wordsets/symbols-sentence.txt")
			return Utility.KeywordMapping((*tokens)[index].Original, keys, values, false)
		},
	},
}
