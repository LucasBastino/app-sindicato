package interfaces

import (
	"database/sql"
	"net/http"
)

type IModel[T TypeModel] interface {
	Imprimir()
	InsertInDB(*sql.DB)
	RenderTemplate(http.ResponseWriter, string)
	DeleteFromDB(*sql.DB)
	UpdateInDB(int, *sql.DB)
	SearchInDB(*http.Request, *sql.DB) []T
	// SearchInDB con generics falta hacer
	// hay que borrar el memberSearcher
}
