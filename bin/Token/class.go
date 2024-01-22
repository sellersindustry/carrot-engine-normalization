package Token

type Class string
type Subclass string


const (
	Control        Class = "{{CLASS:CONTROL}}"
	Space          Class = "{{CLASS:SPACE}}"
	Date           Class = "{{CLASS:DATE}}"
	Time           Class = "{{CLASS:TIME}}"
	Hashtag        Class = "{{CLASS:HASHTAG}}"
	Email          Class = "{{CLASS:EMAIL}}"
	Phone          Class = "{{CLASS:PHONE}}"
	URL            Class = "{{CLASS:URL}}"
	Number         Class = "{{CLASS:NUMBER}}"
	RomanNumeral   Class = "{{CLASS:ROMAN_NUMERAL}}"
	Word           Class = "{{CLASS:WORD}}"
	Symbol         Class = "{{CLASS:SYMBOL}}"
	Punctuation    Class = "{{CLASS:PUNCTUATION}}"
)


const (
	None            Subclass = "{{SUBCLASS:NONE}}"
	NumberPlural    Subclass = "{{SUBCLASS:NUMBER_PLURAL}}"
	NumberOrdinal   Subclass = "{{SUBCLASS:NUMBER_ORDINAL}}"
	NumberNominal   Subclass = "{{SUBCLASS:NUMBER_NOMINAL}}"
	NumberCurrency  Subclass = "{{SUBCLASS:NUMBER_CURRENCY}}"
	NumberYear      Subclass = "{{SUBCLASS:NUMBER_YEAR}}"
	Unit            Subclass = "{{SUBCLASS:UNIT}}"
	Scale           Subclass = "{{SUBCLASS:SCALE}}"
	Range           Subclass = "{{SUBCLASS:RANGE}}"
	Per             Subclass = "{{SUBCLASS:PER}}"
	QuoteStart      Subclass = "{{SUBCLASS:QUOTE_START}}"
	QuoteStartShort Subclass = "{{SUBCLASS:QUOTE_SHORT}}"
	QuoteEnd	    Subclass = "{{SUBCLASS:QUOTE_END}}"
	MathPrefix      Subclass = "{{SUBCLASS:MATH_PREFIX}}"
	MathOperation   Subclass = "{{SUBCLASS:MATH_OPERATION}}"
)

