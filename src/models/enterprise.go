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

func (e Enterprise) Imprimir() {
	fmt.Println(e)
}

func (newEnterprise Enterprise) InsertInDB(DB *sql.DB) {
	insert, err := DB.Query(fmt.Sprintf("INSERT INTO EnterpriseTable (Name, Address) VALUES ('%s','%s')", newEnterprise.Name, newEnterprise.Address))
	if err != nil {
		// DBError{"INSERT Enterprise"}.Error(err)
		fmt.Println("error insertando en la DB")
	}
	defer insert.Close()
}

func (e Enterprise) RenderFileTemplate(w http.ResponseWriter, path string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, e)
}

func (e Enterprise) RenderTableTemplate(w http.ResponseWriter, path string, modelList []Enterprise) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, modelList)
}

func (e Enterprise) DeleteFromDB(DB *sql.DB) {
	delete, err := DB.Query(fmt.Sprintf("DELETE FROM EnterpriseTable WHERE IdEnterprise = '%v'", e.IdEnterprise))
	if err != nil {
		// DBError{"DELETE Enterprise"}.Error(err)
		fmt.Println("error deleting Enterprise")
	}
	defer delete.Close()

}

func (e Enterprise) UpdateInDB(IdEnterprise int, DB *sql.DB) {
	update, err := DB.Query(fmt.Sprintf("UPDATE EnterpriseTable SET Name = '%s', Address = '%s' WHERE IdEnterprise = '%v'", e.Name, e.Address, IdEnterprise))
	if err != nil {
		// DBError{"UPDATE Enterprise"}.Error(err)
		fmt.Println("error updating Enterprise")
		panic(err)
	}
	update.Close()
}

func (e Enterprise) SearchInDB(r *http.Request, DB *sql.DB) []Enterprise {
	searchKey := r.FormValue("search-key")
	var enterprises []Enterprise
	var enterprise Enterprise

	result, err := DB.Query(fmt.Sprintf(`SELECT * FROM EnterpriseTable WHERE Name LIKE '%%%s%%' OR Address LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		fmt.Println("error searching Enterprise in DB")
	}
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			fmt.Println("error scanning Enterprise")
		}
		enterprises = append(enterprises, enterprise)
	}
	defer result.Close()
	return enterprises
}

func (e Enterprise) SearchAllModels(DB *sql.DB) []Enterprise {
	enterprise := Enterprise{}
	enterprises := []Enterprise{}
	result, err := DB.Query("SELECT IdEnterprise, Name, Address FROM enterpriseTable")
	if err != nil {
		fmt.Println("error searching all enterprises")
	}
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			fmt.Println("error scanning data from enterprise")
		}
		enterprises = append(enterprises, enterprise)
	}
	return enterprises
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
