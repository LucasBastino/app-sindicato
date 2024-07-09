package api

import (
	"database/sql"
	"fmt"
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
)

func imprimirCaller[tM i.TypeModel](m i.IModel[tM]) {
	fmt.Println(m)

}

func parserCaller[M i.TypeModel](parser i.ModelParser[M], r *http.Request) M {
	return parser.ParseModel(r)
}

func insertInDBCaller[M i.TypeModel](m i.IModel[M], DB *sql.DB) {
	m.InsertInDB(DB)
}

func renderFileTemplateCaller[M i.TypeModel](m i.IModel[M], w http.ResponseWriter, path string) {
	m.RenderFileTemplate(w, path)
}

func renderTableTemplateCaller[M i.TypeModel](m i.IModel[M], w http.ResponseWriter, path string, modelList []M) {
	m.RenderTableTemplate(w, path, modelList)
}

func deleteFromDBCaller[M i.TypeModel](m i.IModel[M], DB *sql.DB) {
	m.DeleteFromDB(DB)
}

func updateInDBCaller[M i.TypeModel](m i.IModel[M], idModel int, DB *sql.DB) {
	m.UpdateInDB(idModel, DB)
}

func searchInDBCaller[M i.TypeModel](m i.IModel[M], r *http.Request, DB *sql.DB) []M {
	return m.SearchInDB(r, DB)
}

// func searcherCaller[M i.TypeModel](searcher i.ModelSearcher[M], r *http.Request, DB *sql.DB) []M {
// 	return searcher.SearchModel(r, DB)
// }
