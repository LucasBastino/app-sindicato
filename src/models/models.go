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

type Enterprise struct {
	IdEnterprise int
	Name         string
	Address      string
}

type DBError struct {
	Statement string
	Model     string
}
