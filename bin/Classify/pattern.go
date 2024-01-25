package Classify

import (
	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


var REGEX_QUOTE      = "/^\"|'$/";
var REGEX_CURRENCY   = "/^[$£]$/";
var REGEX_OPERATIONS = "/^[\\+\\-\\^\\*\\/\\=x]$/";


var PATTERNS = []*Pattern {
	{
		// Number Plural - Before (Number)
		Current:       string(Token.Number),
		Suffix:        []string{ "s" },
		SetSubclassTo: Token.NumberPlural,
	}, {
		// Number Plural - After (s)
		Current:       "s",
		Prefix:        []string{ string(Token.Number) },
		SetSubclassTo: Token.Unit,
		SetIsSilentTo: true,
	}, {
		// Number Ordinal - Before (Number)
		Current:       string(Token.Number),
		Suffix:        []string{ "/^(st|nd|rd|th)$/" },
		SetSubclassTo: Token.NumberOrdinal,
	}, {
		// Number Ordinal - After (st, nd, rd, th)
		Current:       "/^(st|nd|rd|th)$/",
		Prefix:        []string{ string(Token.Number) },
		SetSubclassTo: Token.Unit,
		SetIsSilentTo: true,
	}, {
		// Number Currency - Symbol
		Current:       REGEX_CURRENCY,
		SetSubclassTo: Token.Unit,
		Suffix:        []string{ string(Token.Number) },
		SetIsSilentTo: true,
	}, {
		// Number Currency Range - Number 1st
		Current:       string(Token.Number),
		Prefix:        []string{ REGEX_CURRENCY },
		Suffix:        []string{ IGNORE_SPACES, "-", string(Token.Number) },
		SetSubclassTo: Token.None,
	}, {
		// Number Currency - Number
		Current:       string(Token.Number),
		Prefix:        []string{ REGEX_CURRENCY },
		SetSubclassTo: Token.NumberCurrency,
	}, {
		// Number Currency Range - Number 2nd
		Current:       string(Token.Number),
		Prefix:        []string{ IGNORE_SPACES, REGEX_CURRENCY, string(Token.Number), "-" },
		SetSubclassTo: Token.NumberCurrency,
	}, {
		// Number Nominal - Address, Phone, Large Number
		Current:       "/^[0-9]{,5}$/",
		SetSubclassTo: Token.NumberNominal,
	}, {
		// Number Year
		Current:       "/^[0-9]{4}$/",
		SetSubclassTo: Token.NumberYear,
	}, {
		// Quote Short - Start (1 Word)
		Current:       REGEX_QUOTE,
		Suffix: 	   []string{ IGNORE_SPACES, string(Token.Word), REGEX_QUOTE },
		SetSubclassTo: Token.QuoteStartShort,
	}, {
		// Quote Short - Start (2 Word)
		Current:       REGEX_QUOTE,
		Suffix: 	   []string{ IGNORE_SPACES, string(Token.Word), string(Token.Word), REGEX_QUOTE },
		SetSubclassTo: Token.QuoteStartShort,
	}, {
		// Quote Short - End (1 Word)
		Current:       REGEX_QUOTE,
		Prefix: 	   []string{ IGNORE_SPACES, REGEX_QUOTE, string(Token.Word) },
		SetSubclassTo: Token.QuoteEnd,
		SetIsSilentTo: true,
	}, {
		// Quote Short - End (2 Word)
		Current:       REGEX_QUOTE,
		Prefix: 	   []string{ IGNORE_SPACES, REGEX_QUOTE, string(Token.Word), string(Token.Word) },
		SetSubclassTo: Token.QuoteEnd,
		SetIsSilentTo: true,
	}, {
		// Quote - Start
		Current:       `/^\"$/`,
		Suffix: 	   []string{ string(Token.Word) },
		SetSubclassTo: Token.QuoteStart,
	}, {
		// Quote - End
		Current:       "/^\"|'$/",
		SetSubclassTo: Token.QuoteEnd,
	}, {
		// Range
		Current:       "-",
		Prefix:        []string{ IGNORE_SPACES, string(Token.Number) },
		Suffix:        []string{ IGNORE_SPACES, string(Token.Number) },
		ScanBefore:    &PatternScan{
			NotExists: []string{ REGEX_OPERATIONS },
			Range: 2,
			IgnoreSpaces: true,
		},
		ScanAfter:     &PatternScan{
			NotExists: []string{ REGEX_OPERATIONS },
			Range: 2,
			IgnoreSpaces: true,
		},
		SetSubclassTo: Token.Range,
	}, {
		// Math Operation
		Current:       REGEX_OPERATIONS,
		Prefix:        []string{ IGNORE_SPACES, string(Token.Number) },
		Suffix:        []string{ IGNORE_SPACES, string(Token.Number) },
		SetSubclassTo: Token.MathOperation,
	}, {
		// Number Prefixes
		Current:       "/^[+-]$/",
		Suffix:        []string{ IGNORE_SPACES, string(Token.Number) },
		SetSubclassTo: Token.MathPrefix,
	}, {
		// Number Currency Prefixes
		Current:       "/^[+-]$/",
		Suffix:        []string{ IGNORE_SPACES, REGEX_CURRENCY, string(Token.Number) },
		SetSubclassTo: Token.MathPrefix,
	}, {
		// Per Number
		Current:       "/",
		Prefix:        []string{ IGNORE_SPACES, string(Token.Number) },
		Suffix:        []string{ string(Token.Word) },
		SetSubclassTo: Token.Per,
	}, {
		// Per Per Number
		Current:       "/",
		Prefix:        []string{ string(Token.Per), string(Token.Word) },
		Suffix:        []string{ string(Token.Word) },
		SetSubclassTo: Token.Per,
	}, {
		// Number Currency Scale
		Current:       "/^(k|m|b|t)$/",
		Prefix:        []string{ IGNORE_SPACES, string(Token.NumberCurrency) },
		SetSubclassTo: Token.Scale,
	}, {
		Current:       string(Token.RomanNumeral),
		Prefix:        []string{ "/^[A-Z]([a-zA-Z])+$/i", string(Token.Space) },
		SetSubclassTo: Token.RomanNumeralPossessive,
	//! abbreviations
	}, {
		// Units
		Current:       `/^` + Utility.RegexWordListOr(Utility.GetWordsetBoth("./bin/wordsets/units.txt")) + `$/i`,
		Prefix:        []string{ IGNORE_SPACES, string(Token.Number) },
		SetSubclassTo: Token.Unit,
	}, {
		// Non-Silent Symbols
		Current:       "/^[^A-Za-z@#$&-=]$/",
		SetSubclassTo: Token.None,
		SetIsSilentTo: true,
	},
}

