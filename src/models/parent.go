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
	insert, err := DB.Query(fmt.Sprintf("INSERT INTO ParentTable (Name, Rel, IdMember) VALUES ('%s','%s', '%d')", newParent.Name, newParent.Rel, newParent.IdMember))
	if err != nil {
		// DBError{"INSERT Parent"}.Error(err)
		fmt.Println("error inserting parent")
		panic(err)
	}
	defer insert.Close()
}

func (p Parent) DeleteModel(DB *sql.DB) {
	delete, err := DB.Query(fmt.Sprintf("DELETE FROM ParentTable WHERE IdParent = '%v'", p.IdParent))
	if err != nil {
		// DBError{"DELETE Parent"}.Error(err)
		fmt.Println("error deleting parent")
		panic(err)
	}
	defer delete.Close()

}

func (p Parent) EditModel(IdParent int, DB *sql.DB) {
	update, err := DB.Query(fmt.Sprintf("UPDATE ParentTable SET Name = '%s', Rel = '%s' WHERE IdParent = '%v'", p.Name, p.Rel, IdParent))
	if err != nil {
		// DBError{"UPDATE Parent"}.Error(err)
		fmt.Println("error updating parent")
		panic(err)
	}
	defer update.Close()
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
	result, err := DB.Query(fmt.Sprintf("SELECT IdParent, Name, Rel, IdMember FROM ParentTable WHERE IdParent = '%v'", IdParent))
	if err != nil {
		fmt.Println("error searching parent by id")
		panic(err)
	}

	var parent Parent
	for result.Next() {
		err = result.Scan(&parent.IdParent, &parent.Name, &parent.Rel, &parent.IdMember)
		if err != nil {
			fmt.Println("error scanning parent")
			panic(err)
		}
	}
	defer result.Close()
	return parent
}

func (p Parent) SearchModels(r *http.Request, DB *sql.DB) []Parent {
	searchKey := r.FormValue("search-key")
	var parents []Parent
	var parent Parent

	result, err := DB.Query(fmt.Sprintf(`SELECT IdParent, Name, Rel FROM ParentTable WHERE Name LIKE '%%%s%%' OR Rel LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		fmt.Println("error searching Parent in DB")
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&parent.IdParent, &parent.Name, &parent.Rel)
		if err != nil {
			fmt.Println("error scanning Parent")
			panic(err)
		}
		parents = append(parents, parent)
	}
	defer result.Close()
	return parents
}

func (p Parent) RenderFileTemplate(w http.ResponseWriter, path string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
		panic(err)
	}
	tmpl.Execute(w, p)
}

func (p Parent) RenderTableTemplate(w http.ResponseWriter, path string, modelList []Parent) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
		panic(err)
	}
	tmpl.Execute(w, modelList)
}

func (p Parent) RenderCreateModelForm(w http.ResponseWriter, path string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("error parsing file", path)
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func (m Parent) ValidateFields(r *http.Request) bool {
	return true
}
