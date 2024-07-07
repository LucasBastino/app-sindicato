package api

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) createEnterprise(w http.ResponseWriter, r *http.Request) {
	var enterprise models.Enterprise
	enterprise.Name = r.FormValue("name")
	enterprise.Address = r.FormValue("address")
	// parseEnterprise()

	insert, err := c.DB.Query(fmt.Sprintf("INSERT INTO EnterpriseTable (Name, Address) VALUES ('%s', '%s')", enterprise.Name, enterprise.Address))
	fmt.Println("error inserting data to database")
	if err != nil {
		DBError{"INSERT ENTERPRISE"}.Error(err)
	}
	defer insert.Close()

	path := "src/views/files/enterpriseFile.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, enterprise)

}

func (c *Controller) searchEnterprise(w http.ResponseWriter, r *http.Request) {
	searchKey := r.FormValue("search-key")
	var enterprises []models.Enterprise
	var enterprise models.Enterprise

	result, err := c.DB.Query(fmt.Sprintf(`SELECT * FROM EnterpriseTable WHERE Name LIKE '%%%s%%' OR Address LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		DBError{"SELECT ENTERPRISE"}.Error(err)
	}
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			ScanError{"ENTERPRISE"}.Error(err)
		}
		enterprises = append(enterprises, enterprise)
	}
	defer result.Close()

	path := "src/views/tables/enterpriseTable.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, enterprises)
}
