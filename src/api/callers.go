package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

func imprimirCaller(m IModel) {
	fmt.Println(m)
}

func insertInDBCaller(m IModel, DB *sql.DB) {
	m.InsertInDB(DB)
}

func renderTemplateCaller(m IModel, w http.ResponseWriter, path string) {
	m.RenderTemplate(w, path)
}

func makerCaller(maker ModelMaker, r *http.Request) IModel {
	return maker.makeModel(r)
}
