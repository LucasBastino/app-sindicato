package api

import (
	"database/sql"
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
)

func parserCaller[M i.TypeModel](parser i.ModelParser[M], r *http.Request) M {
	return parser.ParseModel(r)
}

func insertModelCaller[M i.TypeModel](m i.IModel[M], DB *sql.DB) {
	m.InsertModel(DB)
}

func deleteModelCaller[M i.TypeModel](m i.IModel[M], DB *sql.DB) {
	m.DeleteModel(DB)
}

func editModelCaller[M i.TypeModel](m i.IModel[M], idModel int, DB *sql.DB) {
	m.EditModel(idModel, DB)
}

func getIdModelCaller[M i.TypeModel](m i.IModel[M], r *http.Request) int {
	return m.GetIdModel(r)
}

func searchOneModelByIdCaller[M i.TypeModel](m i.IModel[M], r *http.Request, DB *sql.DB) M {
	return m.SearchOneModelById(r, DB)
}

func searchModelsCaller[M i.TypeModel](m i.IModel[M], r *http.Request, DB *sql.DB) []M {
	return m.SearchModels(r, DB)
}

func renderFileTemplateCaller[M i.TypeModel](m i.IModel[M], w http.ResponseWriter, path string) {
	m.RenderFileTemplate(w, path)
}

func renderTableTemplateCaller[M i.TypeModel](m i.IModel[M], w http.ResponseWriter, path string, modelList []M) {
	m.RenderTableTemplate(w, path, modelList)
}

func renderCreateModelFormCaller[M i.TypeModel](m i.IModel[M], w http.ResponseWriter, path string) {
	m.RenderCreateModelForm(w, path)
}
