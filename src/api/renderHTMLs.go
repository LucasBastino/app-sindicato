package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type Info struct {
	Action string
	Member models.Member
}

var funcMap = template.FuncMap{"ShowIfEdit": ShowIfEdit}

func (c *Controller) renderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/views/index.html", "src/views/footer.html"))
	tmpl.Execute(w, nil)
}

func (c *Controller) renderCreateMemberForm(w http.ResponseWriter, req *http.Request) {
	// creo el template para crear un afiliado y lo ejecuto
	tmpl, err := template.New("memberForm.html").Funcs(funcMap).ParseFiles("src/views/memberForm.html")
	if err != nil {
		fmt.Println("error parsing file memberForm.html")
	}
	// tmpl, _ := template.ParseFiles("src/views/memberForm.html", "src/views/footer.html")
	// el primero siempre es el main template, los demas se usan como componentes
	tmpl.Execute(w, Info{"create", models.Member{}}) // le paso un member vacio, no se puede pasar nil
}

func (c *Controller) renderEditMemberForm(w http.ResponseWriter, r *http.Request) {

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
	tmpl, err := template.New("memberForm.html").Funcs(funcMap).ParseFiles("src/views/memberForm.html")
	if err != nil {
		fmt.Println("error parsing file memberForm.html")
	}
	tmpl.Execute(w, Info{"edit", memberToEdit})
}
