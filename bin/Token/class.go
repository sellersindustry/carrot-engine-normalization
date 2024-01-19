package Token

type Class int8
type Subclass int8


const (
	Control        Class = iota
	Space          Class = iota
	Date           Class = iota
	Time           Class = iota
	Hashtag        Class = iota
	Email          Class = iota
	Phone          Class = iota
	URL            Class = iota
	Number         Class = iota
	RomanNumeral   Class = iota
	Word           Class = iota
	Symbol         Class = iota
	Punctuation    Class = iota
)


const (
	None          Subclass = iota
	NumberPlural  Subclass = iota
	NumberOrdinal Subclass = iota
)

