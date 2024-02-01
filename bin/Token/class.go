package Token

type Class string
type Subclass string

const (
	Control     Class = "{{CLASS:CONTROL}}"
	Space       Class = "{{CLASS:SPACE}}"
	Date        Class = "{{CLASS:DATE}}"
	Time        Class = "{{CLASS:TIME}}"
	Special     Class = "{{CLASS:SPECIAL}}"
	Phone       Class = "{{CLASS:PHONE}}"
	Number      Class = "{{CLASS:NUMBER}}"
	Word        Class = "{{CLASS:WORD}}"
	Symbol      Class = "{{CLASS:SYMBOL}}"
	Termination Class = "{{CLASS:TERMINATION}}"
)

const (
	None                   Subclass = "{{SUBCLASS:NONE}}"
	NewParagraph           Subclass = "{{SUBCLASS:NEW_PARAGRAPH}}"
	NumberPlural           Subclass = "{{SUBCLASS:NUMBER_PLURAL}}"
	NumberOrdinal          Subclass = "{{SUBCLASS:NUMBER_ORDINAL}}"
	NumberNominal          Subclass = "{{SUBCLASS:NUMBER_NOMINAL}}"
	NumberCurrency         Subclass = "{{SUBCLASS:NUMBER_CURRENCY}}"
	NumberYear             Subclass = "{{SUBCLASS:NUMBER_YEAR}}"
	RomanNumeral           Subclass = "{{SUBCLASS:ROMAN_NUMERAL}}"
	RomanNumeralPossessive Subclass = "{{SUBCLASS:ROMAN_NUMERAL_POSSESSIVE}}"
	Punctuation            Subclass = "{{SUBCLASS:PUNCTUATION}}"
	Initialism             Subclass = "{{SUBCLASS:INITIALISM}}"
	Unit                   Subclass = "{{SUBCLASS:UNIT}}"
	Scale                  Subclass = "{{SUBCLASS:SCALE}}"
	Range                  Subclass = "{{SUBCLASS:RANGE}}"
	Per                    Subclass = "{{SUBCLASS:PER}}"
	QuoteStart             Subclass = "{{SUBCLASS:QUOTE_START}}"
	QuoteStartShort        Subclass = "{{SUBCLASS:QUOTE_SHORT}}"
	QuoteEnd               Subclass = "{{SUBCLASS:QUOTE_END}}"
	MathPrefix             Subclass = "{{SUBCLASS:MATH_PREFIX}}"
	MathOperation          Subclass = "{{SUBCLASS:MATH_OPERATION}}"
)
