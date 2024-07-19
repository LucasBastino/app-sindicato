package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Enterprise struct {
	IdEnterprise int
	Name         string
	Address      string
}

func (newEnterprise Enterprise) InsertModel(DB *sql.DB) {
	insert, err := DB.Query(fmt.Sprintf("INSERT INTO EnterpriseTable (Name, Address) VALUES ('%s','%s')", newEnterprise.Name, newEnterprise.Address))
	if err != nil {
		// DBError{"INSERT Enterprise"}.Error(err)
		fmt.Println("error insertando en la DB")
		panic(err)
	}
	defer insert.Close()
}

func (e Enterprise) DeleteModel(DB *sql.DB) {
	delete, err := DB.Query(fmt.Sprintf("DELETE FROM EnterpriseTable WHERE IdEnterprise = '%v'", e.IdEnterprise))
	if err != nil {
		// DBError{"DELETE Enterprise"}.Error(err)
		fmt.Println("error deleting Enterprise")
		panic(err)
	}
	defer delete.Close()

}

func (e Enterprise) EditModel(IdEnterprise int, DB *sql.DB) {
	update, err := DB.Query(fmt.Sprintf("UPDATE EnterpriseTable SET Name = '%s', Address = '%s' WHERE IdEnterprise = '%v'", e.Name, e.Address, IdEnterprise))
	if err != nil {
		// DBError{"UPDATE Enterprise"}.Error(err)
		fmt.Println("error updating Enterprise")
		panic(err)
	}
	defer update.Close()
}

func (e Enterprise) GetIdModel(r *http.Request) int {
	IdEnterpriseStr := r.PathValue("IdEnterprise")
	IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
	if err != nil {
		fmt.Println("error converting type")
		panic(err)
	}
	return IdEnterprise
}

func (e Enterprise) SearchOneModelById(r *http.Request, DB *sql.DB) Enterprise {
	IdEnterprise := e.GetIdModel(r)
	result, err := DB.Query(fmt.Sprintf("SELECT IdEnterprise, Name, Address FROM EnterpriseTable WHERE IdEnterprise = '%v'", IdEnterprise))
	if err != nil {
		fmt.Println("error searching enterprise by id")
		panic(err)
	}

	var enterprise Enterprise
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			fmt.Println("error scanning Enterprise")
			panic(err)
		}
	}
	defer result.Close()
	return enterprise
}

func (e Enterprise) SearchModels(r *http.Request, DB *sql.DB) []Enterprise {
	searchKey := r.FormValue("search-key")
	var enterprises []Enterprise
	var enterprise Enterprise

	result, err := DB.Query(fmt.Sprintf(`SELECT * FROM EnterpriseTable WHERE Name LIKE '%%%s%%' OR Address LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		fmt.Println("error searching Enterprise in DB")
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			fmt.Println("error scanning Enterprise")
			panic(err)
		}
		enterprises = append(enterprises, enterprise)
	}
	defer result.Close()
	return enterprises
}

func (e Enterprise) RenderFileTemplate(w http.ResponseWriter, path string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
		panic(err)
	}
	tmpl.Execute(w, e)
}

func (e Enterprise) RenderTableTemplate(w http.ResponseWriter, path string, modelList []Enterprise) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
		panic(err)
	}
	tmpl.Execute(w, modelList)
}

func (e Enterprise) RenderCreateModelForm(w http.ResponseWriter, path string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("error parsing file", path)
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func (m Enterprise) ValidateFields(r *http.Request) bool {
	return true
}
