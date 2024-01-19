package token


func NewGeneral(orginal string, class Class) *Model {
	return &Model{class, None, []string{}, orginal, false};
}

