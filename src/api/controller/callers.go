package api

import (
	"database/sql"
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

func parserCaller[M models.TypeModel](parser i.ModelParser[M], r *http.Request) M {
	return parser.ParseModel(r)
}

func insertModelCaller[M models.TypeModel](m i.IModel[M], DB *sql.DB) {
	m.InsertModel(DB)
}

func deleteModelCaller[M models.TypeModel](m i.IModel[M], DB *sql.DB) {
	m.DeleteModel(DB)
}

func editModelCaller[M models.TypeModel](m i.IModel[M], idModel int, DB *sql.DB) {
	m.EditModel(idModel, DB)
}

func getIdModelCaller[M models.TypeModel](m i.IModel[M], r *http.Request) int {
	return m.GetIdModel(r)
}

func searchOneModelByIdCaller[M models.TypeModel](m i.IModel[M], r *http.Request, DB *sql.DB) M {
	return m.SearchOneModelById(r, DB)
}

func searchModelsCaller[M models.TypeModel](m i.IModel[M], r *http.Request, DB *sql.DB) []M {
	return m.SearchModels(r, DB)
}

func renderFileTemplateCaller[M models.TypeModel](m i.IModel[M], w http.ResponseWriter, templateData models.TemplateData) {
	m.RenderFileTemplate(w, templateData)
}

func renderTableTemplateCaller[M models.TypeModel](m i.IModel[M], w http.ResponseWriter, templateData models.TemplateData) {
	m.RenderTableTemplate(w, templateData)
}

func renderCreateModelFormCaller[M models.TypeModel](m i.IModel[M], w http.ResponseWriter, templateData models.TemplateData) {
	m.RenderCreateModelForm(w, templateData)
}

func validateFieldsCaller[M models.TypeModel](m i.IModel[M], r *http.Request) map[string]string {
	return m.ValidateFields(r)
}

func createTemplateDataCaller[M models.TypeModel](m i.IModel[M], model M, models []M, path string, errorMap map[string]string) models.TemplateData {
	return m.CreateTemplateData(model, models, path, errorMap)
}
