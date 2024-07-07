package api

import (
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type ModelParser interface {
	parseModel(*http.Request) IModel
}

type MemberParser struct{}

func (m MemberParser) parseModel(r *http.Request) IModel {
	member := models.Member{}
	member.Name = r.FormValue("name")
	member.DNI = r.FormValue("dni")
	return member
}
