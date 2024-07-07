package api

import (
	"database/sql"
	"fmt"
	"net/http"

	interfaces "github.com/LucasBastino/app-sindicato/src/api/interfaces/model.go"
)

func imprimirCaller(m interfaces.IModel) {
	fmt.Println(m)
}

func parserCaller(parser ModelParser, r *http.Request) IModel {
	return parser.parseModel(r)
}

func insertInDBCaller(m IModel, DB *sql.DB) {
	m.InsertInDB(DB)
}

func renderTemplateCaller(m IModel, w http.ResponseWriter, path string) {
	m.RenderTemplate(w, path)
}

func deleteFromDBCaller(m IModel, DB *sql.DB) {
	m.DeleteFromDB(DB)
}

func updateInDBCaller(m IModel, idModel int, DB *sql.DB) {
	m.UpdateInDB(idModel, DB)
}

func searcherCaller(searcher ModelSearcher, r *http.Request, DB *sql.DB) []*IModel {
	return searcher.searchModel(r, DB)
}
