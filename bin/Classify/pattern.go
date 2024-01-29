package Classify

import (
	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


var REGEX_QUOTE       = "/^\"|'$/";
var REGEX_ROMAN_NUM   = "/^(([MDCLXVI]*[MDCLXV]+[MDCLXVI]*)|(III)|(II))$/";
var REGEX_CURRENCY    = "/^[$£]$/";
var REGEX_OPERATIONS  = "/^[\\+\\-\\^\\*\\/\\=x]$/";
var REGEX_PUNCTUATION = "/^[\\.,\\?\\!\\:\\;]$/";


var PATTERNS = []*Pattern {
	{
		// Number Plural - Before (Number)
		CurrentByClass:     Token.Number,
		HasSuffix:          []string{ "s" },
		SetSubclassTo:      Token.NumberPlural,
	}, {
		// Number Plural - After (s)
		CurrentByRegexp:    Utility.CompileRegex(`/^(s)$/`),
		HasPrefix:          []string{ string(Token.Number) },
		SetSubclassTo:      Token.Unit,
		SetIsInactiveTo:    true,
	}, {
		// Number Ordinal - Before (Number)
		CurrentByClass:     Token.Number,
		HasSuffix:          []string{ "/^(st|nd|rd|th)$/" },
		SetSubclassTo:      Token.NumberOrdinal,
	}, {
		// Number Ordinal - After (st, nd, rd, th)
		CurrentByRegexp:    Utility.CompileRegex(`/^(st|nd|rd|th)$/`),
		HasPrefix:          []string{ string(Token.Number) },
		SetSubclassTo:      Token.Unit,
		SetIsInactiveTo:    true,
	}, {
		// Number Currency - Symbol
		CurrentByRegexp:    Utility.CompileRegex(REGEX_CURRENCY),
		SetSubclassTo:      Token.Unit,
		HasSuffix:          []string{ string(Token.Number) },
		SetIsInactiveTo:    true,
	}, {
		// Number Currency Range - Number 1st
		CurrentByClass:     Token.Number,
		HasPrefix:          []string{ REGEX_CURRENCY },
		HasSuffix:          []string{ IGNORE_SPACES, "-", string(Token.Number) },
		SetSubclassTo:      Token.None,
	}, {
		// Number Currency - Number
		CurrentByClass:     Token.Number,
		HasPrefix:          []string{ REGEX_CURRENCY },
		SetSubclassTo:      Token.NumberCurrency,
	}, {
		// Number Currency Range - Number 2nd
		CurrentByClass:     Token.Number,
		HasPrefix:          []string{ IGNORE_SPACES, REGEX_CURRENCY, string(Token.Number), "-" },
		SetSubclassTo:      Token.NumberCurrency,
	}, {
		// Number Nominal - Address, Phone, Large Number
		CurrentByRegexp:    Utility.CompileRegex("/^[0-9]{5,}$/"),
		SetSubclassTo:      Token.NumberNominal,
	}, {
		// Number Year
		CurrentByRegexp:    Utility.CompileRegex("/^[0-9]{4}$/"),
		SetSubclassTo:      Token.NumberYear,
	}, {
		// Quote Short - Start (1-2 Words)
		CurrentByRegexp:    Utility.CompileRegex(REGEX_QUOTE),
		HasSuffix: 	        []string{ string(Token.Word) },
		ScanAfter: &PatternScan{
			Exists: []string{ REGEX_QUOTE },
			Range: 3,
			IgnoreSpaces: true,
		},
		SetSubclassTo:      Token.QuoteStartShort,
	}, {
		// Quote Short - End (1-2 Words)
		CurrentByRegexp:    Utility.CompileRegex(REGEX_QUOTE),
		ScanBefore: &PatternScan{
			Exists: []string{ string(Token.QuoteStartShort) },
			Range: 3,
			IgnoreSpaces: true,
		},
		SetIsInactiveTo:    true,
	}, {
		// Quote - Start
		CurrentByRegexp:    Utility.CompileRegex(REGEX_QUOTE),
		HasSuffix: 	        []string{ string(Token.Word) },
		SetSubclassTo:      Token.QuoteStart,
	}, {
		// Quote - End (1-2 Words)
		CurrentByRegexp:    Utility.CompileRegex(REGEX_QUOTE),
		ScanBefore: &PatternScan{
			Exists: []string{ string(Token.QuoteStart) },
			Range: 100,
			IgnoreSpaces: true,
		},
		SetSubclassTo:      Token.QuoteEnd,
	}, {
		// Quote - End
		CurrentByRegexp:    Utility.CompileRegex(REGEX_QUOTE),
		ScanBefore: &PatternScan{
			Exists: []string{ string(Token.QuoteStartShort) },
			Range: 100,
			IgnoreSpaces: true,
		},
		SetSubclassTo:      Token.QuoteEnd,
	}, {
		// Range
		CurrentByWords:     []string{ "-" },
		HasPrefix:          []string{ IGNORE_SPACES, string(Token.Number) },
		HasSuffix:          []string{ IGNORE_SPACES, string(Token.Number) },
		ScanBefore: &PatternScan{
			NotExists: []string{ REGEX_OPERATIONS },
			Range: 2,
			IgnoreSpaces:   true,
		},
		ScanAfter: &PatternScan{
			NotExists: []string{ REGEX_OPERATIONS },
			Range: 2,
			IgnoreSpaces: true,
		},
		SetSubclassTo:      Token.Range,
	}, {
		// Math Operation
		CurrentByWords:     []string{ "-" },
		HasPrefix:          []string{ IGNORE_SPACES, string(Token.Number) },
		HasSuffix:          []string{ IGNORE_SPACES, string(Token.Number) },
		SetSubclassTo:      Token.MathOperation,
	}, {
		// Number Prefixes
		CurrentByWords:     []string{ "-", "+"},
		HasSuffix:          []string{ IGNORE_SPACES, string(Token.Number) },
		SetSubclassTo:      Token.MathPrefix,
	}, {
		// Number Currency Prefixes
		CurrentByWords:     []string{ "-", "+" },
		HasSuffix:          []string{ IGNORE_SPACES, REGEX_CURRENCY, string(Token.Number) },
		SetSubclassTo:      Token.MathPrefix,
	}, {
		// Per Number
		CurrentByWords:     []string{ "/" },
		HasPrefix:          []string{ IGNORE_SPACES, string(Token.Number) },
		HasSuffix:          []string{ string(Token.Word) },
		SetSubclassTo:      Token.Per,
	}, {
		// Per Per Number
		CurrentByWords:     []string{ "/" },
		HasPrefix:          []string{ string(Token.Per) },
		HasSuffix:          []string{ string(Token.Word) },
		SetSubclassTo:      Token.Per,
	}, {
		// Number Currency Scale
		CurrentByWords:     []string{ "k", "m", "b", "t" },
		HasPrefix:          []string{ IGNORE_SPACES, string(Token.NumberCurrency) },
		SetSubclassTo:      Token.Scale,
	}, {
		// Roman Numeral Possessive
		CurrentByRegexp:    Utility.CompileRegex(REGEX_ROMAN_NUM),
		HasPrefix:          []string{ "/^[A-Z]([a-zA-Z])+$/i", string(Token.Space) },
		SetSubclassTo:      Token.RomanNumeralPossessive,
	}, {
		// Roman Numeral
		CurrentByRegexp:    Utility.CompileRegex(REGEX_ROMAN_NUM),
		SetSubclassTo:      Token.RomanNumeral,
	}, {
		// Units
		CurrentByWords:     Utility.GetWordsetBoth("./bin/wordsets/units.txt"),
		HasPrefix:          []string{ IGNORE_SPACES, string(Token.Number) },
		SetSubclassTo:      Token.Unit,
	}, {
		// New Pharagraph
		CurrentByRegexp:    Utility.CompileRegex(`/^\s*\n\s*$/`),
		HasPrefix:          []string{ string(Token.Punctuation) },
		SetSubclassTo:      Token.NewParagraph,
	}, {
		// Punctuation
		CurrentByRegexp:    Utility.CompileRegex(REGEX_PUNCTUATION),
		SetSubclassTo:      Token.Punctuation,
		HasPrefix:          []string{ `/^[^\s]+$/` },
		HasSuffix:          []string{ string(Token.Space) },
	}, {
		// Punctuation - End of Text
		CurrentByRegexp:    Utility.CompileRegex(REGEX_PUNCTUATION),
		SetSubclassTo:      Token.Punctuation,
		HasPrefix:          []string{ `/^[^\s]+$/` },
		HasSuffix:          []string{ string(Token.Termination) },
	}, {
		// Silent Symbols
		CurrentByRegexp:    Utility.CompileRegex(`/^[^.@#$%&+=~0-9\/\sa-zA-Z]$/`),
		SetSubclassTo:      Token.None,
		SetIsInactiveTo:    true,
	}, {
		// Initialism
		CurrentByRegexp:    Utility.CompileRegex(`/^[A-Z]{2,}$/`),
		Filter: func (text string) bool {
			return Utility.IsInitialism(text);
		},
		SetSubclassTo:      Token.Initialism,
	},
	//! abbreviations
	//! PERIODS IN BETWEEN ACRYONMS
}

