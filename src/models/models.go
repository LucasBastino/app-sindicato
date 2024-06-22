package models

type Member struct {
	IdMember int
	Name     string
	DNI      string
}

type Parent struct {
	IdParent int
	Name     string
	Rel      string
	IdMember int
}
