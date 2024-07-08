package api

import (
	"database/sql"
	"fmt"
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
)

func imprimirCaller(m i.IModel) {
	fmt.Println(m)
}

func parserCaller(parser i.ModelParser, r *http.Request) i.IModel {
	return parser.ParseModel(r)
}

func insertInDBCaller(m i.IModel, DB *sql.DB) {
	m.InsertInDB(DB)
}

func renderTemplateCaller(m i.IModel, w http.ResponseWriter, path string) {
	m.RenderTemplate(w, path)
}

func deleteFromDBCaller(m i.IModel, DB *sql.DB) {
	m.DeleteFromDB(DB)
}

func updateInDBCaller(m i.IModel, idModel int, DB *sql.DB) {
	m.UpdateInDB(idModel, DB)
}

// probar sacar esto y hacerlo como antes pero con generics
// igualmente primero probar asi que ya esta hecho

func searchInDBCaller(m i.IModel, r *http.Request, DB *sql.DB) {
	m.SearchInDB(r, DB)
}
