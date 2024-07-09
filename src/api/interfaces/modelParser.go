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
