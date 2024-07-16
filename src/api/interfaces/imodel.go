package interfaces

import (
	"database/sql"
	"net/http"
)

type IModel[M TypeModel] interface {
	Imprimir()
	InsertInDB(*sql.DB)
	RenderFileTemplate(http.ResponseWriter, string)
	RenderTableTemplate(http.ResponseWriter, string, []M)
	DeleteFromDB(*sql.DB)
	UpdateInDB(int, *sql.DB)
	SearchInDB(*http.Request, *sql.DB) []M
	GetIdModel(*http.Request) int
}
