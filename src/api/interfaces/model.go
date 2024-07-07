package api

import (
	"database/sql"
	"net/http"
)

type IModel interface {
	Imprimir()
	InsertInDB(*sql.DB)
	RenderTemplate(http.ResponseWriter, string)
	DeleteFromDB(*sql.DB)
	UpdateInDB(int, *sql.DB)
}
