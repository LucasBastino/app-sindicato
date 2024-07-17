package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) renderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	tmpl.Execute(w, nil)
}

func (c *Controller) renderCreateMemberForm(w http.ResponseWriter, req *http.Request) {
	// creo el template para crear un afiliado y lo ejecuto
	path := "src/views/forms/createMemberForm.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, nil) // le paso un member vacio, no se puede pasar nil
}

func (c *Controller) renderCreateParentForm(w http.ResponseWriter, r *http.Request) {
	path := "src/views/forms/createParentForm.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, nil)
}

func (c *Controller) renderParentFile(w http.ResponseWriter, r *http.Request) {
	IdParent := r.PathValue("IdParent")
	result, err := c.DB.Query(fmt.Sprintf("SELECT IdParent, Name, Rel, IdMember FROM ParentTable WHERE IdParent = '%s'", IdParent))
	if err != nil {
		DBError{"SELECT PARENT"}.Error(err)
	}

	var parent models.Parent
	for result.Next() {
		err = result.Scan(&parent.IdParent, &parent.Name, &parent.Rel, &parent.IdMember)
		if err != nil {
			ScanError{"PARENT"}.Error(err)
		}
	}

	path := "src/views/files/parentFile.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, parent)

}

func (c *Controller) renderCreateEnterpriseForm(w http.ResponseWriter, r *http.Request) {
	path := "src/views/forms/createEnterpriseForm.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, nil)
}

func (c *Controller) renderEnterpriseTable(w http.ResponseWriter, r *http.Request) {
	var enterprises []models.Enterprise
	var enterprise models.Enterprise

	result, err := c.DB.Query("SELECT * FROM EnterpriseTable")
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
	path := "src/views/tables/enterpriseTable.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, enterprises)
}

// es el mismo procedimiento para members empresas y familiares
// buscar los datos en la bd, scanearlos y despues ejecutar un template
// hay que hacer alguna funcion para simplificar como hacia hdeleon

func (c *Controller) renderEnterpriseFile(w http.ResponseWriter, r *http.Request) {
	IdEnterprise := r.PathValue("IdEnterprise")

	result, err := c.DB.Query(fmt.Sprintf("SELECT * FROM EnterpriseTable WHERE IdEnterprise = '%s'", IdEnterprise))
	if err != nil {
		DBError{"SELECT ENTERPRISE"}.Error(err)
	}

	var enterprise models.Enterprise
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			ScanError{"ENTERPRISE"}.Error(err)
		}
	}

	path := "src/views/files/enterpriseFile.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}

	tmpl.Execute(w, enterprise)

}
