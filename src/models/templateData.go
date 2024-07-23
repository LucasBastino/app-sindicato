package models

type TemplateData struct {
	Path        string
	Member      Member
	Members     []Member
	Parent      Parent
	Parents     []Parent
	Enterprise  Enterprise
	Enterprises []Enterprise
	ErrorMap    map[string]string
}
