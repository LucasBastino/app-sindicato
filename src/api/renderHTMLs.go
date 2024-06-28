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
	tmpl, err := template.ParseFiles("src/views/forms/createMemberForm.html")
	if err != nil {
		fmt.Println("error parsing file createMemberForm.html")
	}
	tmpl.Execute(w, nil) // le paso un member vacio, no se puede pasar nil
}

func (c *Controller) renderMemberFile(w http.ResponseWriter, r *http.Request) {
	var memberToEdit models.Member

	// capto el param id de la URL
	IdMember := r.PathValue("IdMember")

	// busco el miembro por id en la base de datos
	result, err := c.DB.Query(fmt.Sprintf("SELECT * FROM MemberTable WHERE IdMember = '%s'", IdMember))
	if err != nil {
		fmt.Println("error searching member")
	}

	// creo una instancia de member a la que le atribuyo la informacion de la DB
	for result.Next() {
		err := result.Scan(&memberToEdit.IdMember, &memberToEdit.Name, &memberToEdit.DNI)
		if err != nil {
			fmt.Println("error scanning result")
			panic(err)
		}
	}

	// creo el template, le paso sus funciones y lo ejecuto
	tmpl, err := template.ParseFiles("src/views/files/memberFile.html")
	if err != nil {
		fmt.Println("error parsing file editMemberForm.html")
	}
	tmpl.Execute(w, memberToEdit)
}

func (c *Controller) renderCreateParentForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("src/views/forms/createParentForm.html")
	if err != nil {
		fmt.Println("error parsing file createParentForm.html")
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func (c *Controller) renderParentFile(w http.ResponseWriter, r *http.Request) {
	IdParent := r.PathValue("IdParent")
	result, err := c.DB.Query(fmt.Sprintf("SELECT IdParent, Name, Rel, IdMember FROM ParentTable WHERE IdParent = '%s'", IdParent))
	if err != nil {
		fmt.Println("error searching parent from database")
		panic(err)
	}

	var parent models.Parent
	for result.Next() {
		err = result.Scan(&parent.IdParent, &parent.Name, &parent.Rel, &parent.IdMember)
		if err != nil {
			fmt.Println("error scanning data")
			panic(err)
		}
	}

	tmpl, err := template.ParseFiles("src/views/files/parentFile.html")
	if err != nil {
		fmt.Println("error parsing file parentFile.html")
		panic(err)
	}
	tmpl.Execute(w, parent)

}

func (c *Controller) renderCreateEnterpriseForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("src/views/forms/createEnterpriseForm.html")
	if err != nil {
		fmt.Println("error parsing file createEnterpriseForm.html")
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func (c *Controller) renderEnterpriseTable(w http.ResponseWriter, r *http.Request) {
	var enterprises []models.Enterprise
	var enterprise models.Enterprise

	result, err := c.DB.Query("SELECT * FROM EnterpriseTable")
	if err != nil {
		fmt.Println("error getting data from EnterpriseTable")
		panic(err)
	}

	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			fmt.Println("error scanning data")
			panic(err)
		}
		enterprises = append(enterprises, enterprise)
	}
	tmpl, err := template.ParseFiles("src/views/tables/enterpriseTable.html")
	if err != nil {
		fmt.Println("error parsing file enterpriseTable.html")
		panic(err)
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
		fmt.Println("error getting data from database")
		panic(err)
	}

	var enterprise models.Enterprise
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			fmt.Println("error scanning data")
			panic(err)
		}
	}

	tmpl, err := template.ParseFiles("src/views/files/enterpriseFile.html")
	if err != nil {
		fmt.Println("error parsing file enterpriseFile.html")
		panic(err)
	}

	tmpl.Execute(w, enterprise)

}
