package token

type Class int8
type Subclass int8


const (
	Control       Class = iota
	Space         Class = iota
	Date          Class = iota
	Time          Class = iota
	Hashtag       Class = iota
	Email         Class = iota
	Phone         Class = iota
	URL           Class = iota
	Word          Class = iota
	Number        Class = iota
	Symbol        Class = iota
)


const (
	General        Subclass = iota
	NumberRoman    Subclass = iota
	NumberPlural   Subclass = iota
	NumberYear     Subclass = iota
	NumberCurrency Subclass = iota
	NumberFraction Subclass = iota
	NumberNominal  Subclass = iota
	NumberOrdinal  Subclass = iota
)

