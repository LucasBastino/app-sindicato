package models

type Member struct {
	IdMember int
	Name     string
	DNI      string
}

type Parent struct {
	Name     string
	Rel      string
	IdMember int
}
