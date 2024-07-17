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

func (newParent Parent) InsertModel(DB *sql.DB) {
	insert, err := DB.Query(fmt.Sprintf("INSERT INTO ParentTable (Name, Rel) VALUES ('%s','%s')", newParent.Name, newParent.Rel))
	if err != nil {
		// DBError{"INSERT Parent"}.Error(err)
		fmt.Println("error insertando en la DB")
	}
	defer insert.Close()
}

func (p Parent) DeleteModel(DB *sql.DB) {
	delete, err := DB.Query(fmt.Sprintf("DELETE FROM ParentTable WHERE IdParent = '%v'", p.IdParent))
	if err != nil {
		// DBError{"DELETE Parent"}.Error(err)
		fmt.Println("error deleting Parent")
	}
	defer delete.Close()

}

func (p Parent) EditModel(IdParent int, DB *sql.DB) {
	update, err := DB.Query(fmt.Sprintf("UPDATE ParentTable SET Name = '%s', Rel = '%s' WHERE IdParent = '%v'", p.Name, p.Rel, IdParent))
	if err != nil {
		// DBError{"UPDATE Parent"}.Error(err)
		fmt.Println("error updating Parent")
		panic(err)
	}
	update.Close()
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

func (p Parent) SearchOneModelById(r *http.Request, DB *sql.DB) Parent {
	IdParent := p.GetIdModel(r)
	result, err := DB.Query(fmt.Sprintf("SELECT IdParent, Name, Rel FROM ParentTable WHERE IdParent = '%v'", IdParent))
	if err != nil {
		fmt.Println("error searching parent by id")
	}

	var parent Parent
	for result.Next() {
		err = result.Scan(&parent.IdParent, &parent.Name, &parent.Rel)
		if err != nil {
			fmt.Println("error scanning parent")
		}
	}
	return parent
}

func (p Parent) SearchModelsByKey(r *http.Request, DB *sql.DB) []Parent {
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

func (p Parent) SearchAllModels(DB *sql.DB) []Parent {
	parent := Parent{}
	parents := []Parent{}
	result, err := DB.Query("SELECT IdParent, Name, Rel FROM parentTable")
	if err != nil {
		fmt.Println("error searching all parents")
	}
	for result.Next() {
		err = result.Scan(&parent.IdParent, &parent.Name, &parent.Rel)
		if err != nil {
			fmt.Println("error scanning data from parent")
		}
		parents = append(parents, parent)
	}
	return parents
}

func (p Parent) RenderFileTemplate(w http.ResponseWriter, path string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, p)
}

func (p Parent) RenderTableTemplate(w http.ResponseWriter, path string, modelList []Parent) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, modelList)
}
