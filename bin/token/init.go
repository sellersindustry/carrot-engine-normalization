package token


func NewGeneral(orginal string, class Class) *Model {
	return &Model{class, General, []string{}, orginal}
}


func NewSubclass(orginal string, class Class, subclass Subclass) *Model {
	return &Model{class, subclass, []string{}, orginal}
}

