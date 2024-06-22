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
	tmpl, err := template.ParseFiles("/src/views/forms/createMemberForm.html")
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
