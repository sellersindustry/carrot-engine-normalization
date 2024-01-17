package token

type Class int64

const (
	Control       Class = iota
	Space         Class = iota
	Word          Class = iota
	Date          Class = iota
	Time          Class = iota
	Currency      Class = iota
	Hashtag       Class = iota
	Email         Class = iota
	Phone         Class = iota
	URL           Class = iota
	NumberPlural  Class = iota
	NumberOrdinal Class = iota
	NumberRoman   Class = iota
	NumberNominal Class = iota
)

func New(orginal string, class Class) *Model {
	return &Model{class, []string{}, orginal, ""}
}

type Model struct {
	Class          Class
	Text           []string
	Original       string
	Categorization string
}
