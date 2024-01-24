package Token


func NewGeneral(orginal string, class Class) *Model {
	return &Model{class, None, "", orginal, false};
}

