package interfaces

import (
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type ModelParser interface {
	ParseModel(*http.Request) IModel
}

type MemberParser struct{}

func (m MemberParser) ParseModel(r *http.Request) IModel {
	member := models.Member{}
	member.Name = r.FormValue("name")
	member.DNI = r.FormValue("dni")
	return member
}
