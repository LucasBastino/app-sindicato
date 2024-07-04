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
}

type ModelMaker interface {
	makeModel(*http.Request) IModel
}

type MemberMaker struct{}

func (maker MemberMaker) makeModel(r *http.Request) IModel {
	member := models.Member{}
	member.Name = r.FormValue("name")
	member.DNI = r.FormValue("dni")
	return member
}
