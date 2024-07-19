package interfaces

import (
	"database/sql"
	"net/http"
)

type IModel[M TypeModel] interface {
	InsertModel(*sql.DB)
	DeleteModel(*sql.DB)
	EditModel(int, *sql.DB)
	GetIdModel(*http.Request) int
	SearchOneModelById(*http.Request, *sql.DB) M
	SearchModels(*http.Request, *sql.DB) []M
	RenderFileTemplate(http.ResponseWriter, string)
	RenderTableTemplate(http.ResponseWriter, string, []M)
	RenderCreateModelForm(http.ResponseWriter, string)
	ValidateFields(*http.Request) bool
}
