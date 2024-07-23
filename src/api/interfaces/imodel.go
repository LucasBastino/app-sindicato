package interfaces

import (
	"database/sql"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type IModel[M models.TypeModel] interface {
	InsertModel(*sql.DB)
	DeleteModel(*sql.DB)
	EditModel(int, *sql.DB)
	GetIdModel(*http.Request) int
	SearchOneModelById(*http.Request, *sql.DB) M
	SearchModels(*http.Request, *sql.DB) []M
	RenderFileTemplate(http.ResponseWriter, models.TemplateData)
	RenderTableTemplate(http.ResponseWriter, models.TemplateData)
	RenderCreateModelForm(http.ResponseWriter, models.TemplateData)
	ValidateFields(*http.Request) map[string]string
	CreateTemplateData(M, []M, string, map[string]string) models.TemplateData
}
