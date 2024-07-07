package api

import (
	"database/sql"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type IModel interface {
	Imprimir()
	InsertInDB(*sql.DB)
	RenderTemplate(http.ResponseWriter, string)
	DeleteFromDB(*sql.DB)
	UpdateInDB(int, *sql.DB)
	SearchInDB(*http.Request, *sql.DB) []models.Member
}

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
