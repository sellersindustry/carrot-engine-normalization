package Classify

import (
	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)

var REGEX_QUOTE = "/^\"|'$/"
var REGEX_ROMAN_NUM = "/^(([MDCLXVI]*[MDCLXV]+[MDCLXVI]*)|(III)|(II))$/"
var REGEX_CURRENCY = "/^[$£€]$/"
var REGEX_OPERATIONS = "/^[\\+\\-\\^\\*\\/\\=x]$/"
var REGEX_PUNCTUATION = "/^[\\.,\\?\\!\\:\\;]$/"


var PATTERNS = []*Pattern{
	{
		// Number Plural - Before (Number)
		CurrentByClass: Token.Number,
		HasSuffix:      []string{"s"},
		SetSubclassTo:  Token.NumberPlural,
	}, {
		// Number Plural - After (s)
		CurrentByRegexp: Utility.CompileRegex(`/^(s)$/`),
		HasPrefix:       []string{string(Token.Number)},
		SetSubclassTo:   Token.Unit,
		SetIsInactiveTo: true,
	}, {
		// Number Ordinal - Before (Number)
		CurrentByClass: Token.Number,
		HasSuffix:      []string{"/^(st|nd|rd|th)$/"},
		SetSubclassTo:  Token.NumberOrdinal,
	}, {
		// Number Ordinal - After (st, nd, rd, th)
		CurrentByRegexp: Utility.CompileRegex(`/^(st|nd|rd|th)$/`),
		HasPrefix:       []string{string(Token.Number)},
		SetSubclassTo:   Token.Unit,
		SetIsInactiveTo: true,
	}, {
		// Number Ordinal - Dates
		CurrentByClass:  Token.Number,
		Filter: func(text string) bool {
			if len(text) <= 2 {
				return true;
			}
			return false;
		},
		ScanBefore:      &PatternScan{
			Exists:       []string{ "january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december", "jan", "feb", "mar", "apr", "jun", "jul", "aug", "sep", "oct", "nov", "dec" },
			Range:        2,
			IgnoreSpaces: true,
		},
		SetSubclassTo:   Token.NumberOrdinal,
	}, {
		// Number Ordinal - Dates
		CurrentByClass:  Token.Number,
		Filter: func(text string) bool {
			if len(text) <= 2 {
				return true;
			}
			return false;
		},
		ScanAfter:      &PatternScan{
			Exists:       []string{ "january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december", "jan", "feb", "mar", "apr", "jun", "jul", "aug", "sep", "oct", "nov", "dec" },
			Range:        2,
			IgnoreSpaces: true,
		},
		SetSubclassTo:   Token.NumberOrdinal,
	}, {
		// Number Currency - Symbol
		CurrentByRegexp: Utility.CompileRegex(REGEX_CURRENCY),
		SetSubclassTo:   Token.Unit,
		HasSuffix:       []string{string(Token.Number)},
		SetIsInactiveTo: true,
	}, {
		// Number Currency Range - Number 1st
		CurrentByClass: Token.Number,
		HasPrefix:      []string{REGEX_CURRENCY},
		HasSuffix:      []string{IGNORE_SPACES, "-", string(Token.Number)},
		SetSubclassTo:  Token.None,
	}, {
		// Number Currency - Number
		CurrentByClass: Token.Number,
		HasPrefix:      []string{REGEX_CURRENCY},
		SetSubclassTo:  Token.NumberCurrency,
	}, {
		// Number Currency Range - Number 2nd
		CurrentByClass: Token.Number,
		HasPrefix:      []string{IGNORE_SPACES, REGEX_CURRENCY, string(Token.Number), "-"},
		SetSubclassTo:  Token.NumberCurrency,
	}, {
		// Number Nominal - Address, Phone, Large Number
		CurrentByRegexp: Utility.CompileRegex("/^[0-9]{5,}$/"),
		SetSubclassTo:   Token.NumberNominal,
	}, {
		// Number Year (4-Digit)
		CurrentByRegexp: Utility.CompileRegex("/^[0-9]{4}$/"),
		SetSubclassTo:   Token.NumberYear,
	}, {
		// Number Year (Others)
		CurrentByClass: Token.Number,
		ScanAfter: &PatternScan{
			Exists:       []string{"ad", "bc", "bce", "ce"},
			Range:        1,
			IgnoreSpaces: true,
		},
		SetSubclassTo: Token.NumberYear,
	}, {
		// Quote Short - Start (1-2 Words)Ne
		CurrentByRegexp: Utility.CompileRegex(REGEX_QUOTE),
		HasSuffix:       []string{string(Token.Word)},
		ScanAfter: &PatternScan{
			Exists:       []string{REGEX_QUOTE},
			Range:        3,
			IgnoreSpaces: true,
		},
		SetSubclassTo: Token.QuoteStartShort,
	}, {
		// Quote Short - End (1-2 Words)
		CurrentByRegexp: Utility.CompileRegex(REGEX_QUOTE),
		ScanBefore: &PatternScan{
			Exists:       []string{string(Token.QuoteStartShort)},
			Range:        3,
			IgnoreSpaces: true,
		},
		SetIsInactiveTo: true,
	}, {
		// Quote - Start
		CurrentByRegexp: Utility.CompileRegex(REGEX_QUOTE),
		HasSuffix:       []string{string(Token.Word)},
		SetSubclassTo:   Token.QuoteStart,
	}, {
		// Quote - End (1-2 Words)
		CurrentByRegexp: Utility.CompileRegex(REGEX_QUOTE),
		ScanBefore: &PatternScan{
			Exists:       []string{string(Token.QuoteStart)},
			Range:        100,
			IgnoreSpaces: true,
		},
		SetSubclassTo: Token.QuoteEnd,
	}, {
		// Quote - End
		CurrentByRegexp: Utility.CompileRegex(REGEX_QUOTE),
		ScanBefore: &PatternScan{
			Exists:       []string{string(Token.QuoteStartShort)},
			Range:        100,
			IgnoreSpaces: true,
		},
		SetSubclassTo: Token.QuoteEnd,
	}, {
		// Range
		CurrentByWords:  []string{ "-" },
		HasPrefix:       []string{IGNORE_SPACES, string(Token.Number)},
		HasSuffix:       []string{IGNORE_SPACES, string(Token.Number)},
		ScanBefore: &PatternScan{
			NotExists:    []string{REGEX_OPERATIONS},
			Range:        2,
			IgnoreSpaces: true,
		},
		ScanAfter: &PatternScan{
			NotExists:    []string{REGEX_OPERATIONS},
			Range:        2,
			IgnoreSpaces: true,
		},
		SetSubclassTo: Token.Range,
	}, {
		// Range (per)
		CurrentByWords:  []string{ "x" },
		HasPrefix:       []string{IGNORE_SPACES, string(Token.Number)},
		HasSuffix:       []string{IGNORE_SPACES, string(Token.Number)},
		ScanBefore: &PatternScan{
			NotExists:    []string{REGEX_OPERATIONS},
			Range:        2,
			IgnoreSpaces: true,
		},
		ScanAfter: &PatternScan{
			NotExists:    []string{REGEX_OPERATIONS},
			Range:        2,
			IgnoreSpaces: true,
		},
		SetSubclassTo:   Token.Range,
	}, {
		// Range (per)
		CurrentByWords:  []string{ "-" },
		HasPrefix:       []string{IGNORE_SPACES, string(Token.Number)},
		HasSuffix:       []string{IGNORE_SPACES, string(Token.Number), "/", string(Token.Word)},
		SetSubclassTo:   Token.Range,
	}, {
		// Math Operation
		CurrentByWords:  []string{ "-", "+", "/", "*", "x", "^" },
		HasPrefix:       []string{IGNORE_SPACES, string(Token.Number)},
		HasSuffix:       []string{IGNORE_SPACES, string(Token.Number)},
		SetSubclassTo:   Token.MathOperation,
	}, {
		// Dashes
		CurrentByWords:  []string{ "-" },
		HasPrefix:       []string{string(Token.Word)},
		SetIsInactiveTo:  true,
	}, {
		// Number Prefixes
		CurrentByWords: []string{"-", "+"},
		HasSuffix:      []string{IGNORE_SPACES, string(Token.Number)},
		SetSubclassTo:  Token.MathPrefix,
	}, {
		// Number Currency Prefixes
		CurrentByWords: []string{"-", "+"},
		HasSuffix:      []string{IGNORE_SPACES, REGEX_CURRENCY, string(Token.Number)},
		SetSubclassTo:  Token.MathPrefix,
	}, {
		// Per Number
		CurrentByWords: []string{"/"},
		HasPrefix:      []string{IGNORE_SPACES, string(Token.Number)},
		HasSuffix:      []string{string(Token.Word)},
		SetSubclassTo:  Token.Per,
	}, {
		// Per Per Number
		CurrentByWords: []string{"/"},
		HasPrefix:      []string{string(Token.Per)},
		HasSuffix:      []string{string(Token.Word)},
		SetSubclassTo:  Token.Per,
	}, {
		// Number Currency Scale
		CurrentByWords: []string{"k", "thousand", "m", "million", "b", "billion", "t", "trillion"},
		HasPrefix:      []string{IGNORE_SPACES, string(Token.NumberCurrency)},
		SetIsInactiveTo: true,
	}, {
		// Roman Numeral Possessive
		CurrentByRegexp: Utility.CompileRegex(REGEX_ROMAN_NUM),
		HasPrefix:       []string{"/^[A-Z]([a-zA-Z])+$/i", string(Token.Space)},
		SetSubclassTo:   Token.RomanNumeralPossessive,
	}, {
		// Roman Numeral
		CurrentByRegexp: Utility.CompileRegex(REGEX_ROMAN_NUM),
		SetSubclassTo:   Token.RomanNumeral,
	}, {
		// Units
		CurrentByWords: Utility.GetWordsetBoth("./bin/wordsets/units.txt"),
		HasPrefix:      []string{IGNORE_SPACES, string(Token.Number)},
		SetSubclassTo:  Token.Unit,
	}, {
		// New Pharagraph
		CurrentByRegexp: Utility.CompileRegex(`/^\s*\n\s*$/`),
		HasPrefix:       []string{string(Token.Punctuation)},
		SetSubclassTo:   Token.NewParagraph,
	}, {
		// New Pharagraph (Bullet Points)
		CurrentByRegexp: Utility.CompileRegex(`/^-$/`),
		HasPrefix:       []string{`/^\s*\n\s*$/`},
		SetSubclassTo:   Token.NewParagraph,
	}, {
		// Punctuation
		CurrentByRegexp: Utility.CompileRegex(REGEX_PUNCTUATION),
		SetSubclassTo:   Token.Punctuation,
		HasPrefix:       []string{`/^[^\s]+$/`},
		HasSuffix:       []string{string(Token.Space)},
	}, {
		// Punctuation - End of Text
		CurrentByRegexp: Utility.CompileRegex(REGEX_PUNCTUATION),
		SetSubclassTo:   Token.Punctuation,
		HasPrefix:       []string{`/^[^\s]+$/`},
		HasSuffix:       []string{string(Token.Termination)},
	}, {
		// Pause for Parentheses - Start (Space)
		CurrentByClass:  Token.Space,
		HasPrefix:       []string{string(Token.Word)},
		HasSuffix:       []string{`/^\($/`},
		SetIsInactiveTo: true,
	}, {
		// Pause for Parentheses - Start (Parentheses)
		CurrentByRegexp: Utility.CompileRegex(`/^\($/`),
		HasPrefix:       []string{string(Token.Word), string(Token.Space)},
		SetSubclassTo:   Token.Punctuation,
	}, {
		// Pause for Parentheses - End (Parentheses)
		CurrentByRegexp: Utility.CompileRegex(`/^\)$/`),
		HasPrefix:       []string{string(Token.Word)},
		HasSuffix:       []string{string(Token.Space)},
		SetSubclassTo:   Token.Punctuation,
	}, {
		// Silent Symbols
		CurrentByRegexp: Utility.CompileRegex(`/^[^.@#$%&+=~0-9\/\sa-zA-Z]$/`),
		SetSubclassTo:   Token.None,
		SetIsInactiveTo: true,
	}, {
		// Initialism (w/ Periods)
		CurrentByRegexp: Utility.CompileRegex(`/^([A-Z]\.)+([A-Z])(\.)/`),
		SetSubclassTo:   Token.Initialism,
	}, {
		// Initialism
		CurrentByRegexp: Utility.CompileRegex(`/^[A-Z]{2,}(\')?(s)?$/`),
		Filter: func(text string) bool {
			return Utility.IsInitialism(text)
		},
		SetSubclassTo: Token.Initialism,
	},
	//! abbreviations
	//! PERIODS IN BETWEEN ACRYONMS
}
