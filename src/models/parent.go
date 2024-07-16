package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Parent struct {
	IdParent int
	Name     string
	Rel      string
	IdMember int
}

func (m Parent) Imprimir() {
	fmt.Println(m)
}

func (newParent Parent) InsertInDB(DB *sql.DB) {
	insert, err := DB.Query(fmt.Sprintf("INSERT INTO ParentTable (Name, Rel) VALUES ('%s','%s')", newParent.Name, newParent.Rel))
	if err != nil {
		// DBError{"INSERT Parent"}.Error(err)
		fmt.Println("error insertando en la DB")
	}
	defer insert.Close()
}

func (m Parent) RenderFileTemplate(w http.ResponseWriter, path string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, m)
}

func (m Parent) RenderTableTemplate(w http.ResponseWriter, path string, modelList []Parent) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, modelList)
}

func (m Parent) DeleteFromDB(DB *sql.DB) {
	delete, err := DB.Query(fmt.Sprintf("DELETE FROM ParentTable WHERE IdParent = '%v'", m.IdParent))
	if err != nil {
		// DBError{"DELETE Parent"}.Error(err)
		fmt.Println("error deleting Parent")
	}
	defer delete.Close()

}

func (m Parent) UpdateInDB(IdParent int, DB *sql.DB) {
	update, err := DB.Query(fmt.Sprintf("UPDATE ParentTable SET Name = '%s', Rel = '%s' WHERE IdParent = '%v'", m.Name, m.Rel, IdParent))
	if err != nil {
		// DBError{"UPDATE Parent"}.Error(err)
		fmt.Println("error updating Parent")
		panic(err)
	}
	update.Close()
}

func (m Parent) SearchInDB(r *http.Request, DB *sql.DB) []Parent {
	searchKey := r.FormValue("search-key")
	var parents []Parent
	var parent Parent

	result, err := DB.Query(fmt.Sprintf(`SELECT * FROM ParentTable WHERE Name LIKE '%%%s%%' OR Rel LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		fmt.Println("error searching Parent in DB")
	}
	for result.Next() {
		err = result.Scan(&parent.IdParent, &parent.Name, &parent.Rel)
		if err != nil {
			fmt.Println("error scanning Parent")
		}
		parents = append(parents, parent)
	}
	defer result.Close()
	return parents
}

func (p Parent) GetIdModel(r *http.Request) int {
	IdParentStr := r.PathValue("IdParent")
	IdParent, err := strconv.Atoi(IdParentStr)
	if err != nil {
		fmt.Println("error converting type")
		panic(err)
	}
	return IdParent
}
