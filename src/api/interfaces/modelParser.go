package interfaces

import (
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type ModelParser[tM TypeModel] interface {
	ParseModel(*http.Request) tM
}

type MemberParser struct{}

func (m MemberParser) ParseModel(r *http.Request) models.Member {
	member := models.Member{}
	member.Name = r.FormValue("name")
	member.DNI = r.FormValue("dni")
	return member
}

type ParentParser struct{}

func (p ParentParser) ParseModel(r *http.Request) models.Parent {
	parent := models.Parent{}
	parent.Name = r.PathValue("name")
	parent.Rel = r.PathValue("rel")
	return parent
}

type EnterpriseParser struct{}

func (p EnterpriseParser) ParseModel(r *http.Request) models.Enterprise {
	enterprise := models.Enterprise{}
	enterprise.Name = r.PathValue("name")
	enterprise.Address = r.PathValue("address")
	return enterprise
}
