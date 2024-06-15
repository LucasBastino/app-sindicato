package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type Info struct {
	Action string
	Member models.Member
}

var funcMap = template.FuncMap{"ShowOrNot": ShowOrNot}

func (c *Controller) renderHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/views/home.html"))
	tmpl.Execute(w, nil)
}

func (c *Controller) renderMembers1(w http.ResponseWriter, r *http.Request) {
	result, err := c.DB.Query("SELECT IdMember, Name, DNI FROM MemberTable")
	if err != nil {
		fmt.Println("error obtaining data from database")
		log.Panic(err)
	}

	var members []models.Member
	for result.Next() {
		member := models.Member{}
		err := result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning data")
			log.Panic(err)
		}
		members = append(members, member)
	}

	tmpl, err := template.ParseFiles("src/views/memberList.html", "src/views/footer.html")
	if err != nil {
		fmt.Println("error parsing files")
		panic(err)
	}

	tmpl.Execute(w, members)
}

func (c *Controller) renderMembers2(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("src/views/members2.html", "src/views/footer.html")
	if err != nil {
		fmt.Println("error parsing file memberList2")
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func (c *Controller) renderMemberList(w http.ResponseWriter, r *http.Request) {
	result, err := c.DB.Query("SELECT IdMember, Name, DNI FROM MemberTable")
	if err != nil {
		fmt.Println("error obtaining data from database")
		log.Panic(err)
	}

	var members []models.Member
	for result.Next() {
		member := models.Member{}
		err := result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning data")
			log.Panic(err)
		}
		members = append(members, member)
	}

	tmpl := returnHtmlTemplate("src/views/memberList.html")
	tmpl.Execute(w, members)
}

func (c *Controller) renderCreateMemberForm(w http.ResponseWriter, req *http.Request) {
	// creo el template para crear un afiliado y lo ejecuto
	tmpl, err := template.New("memberForm.html").Funcs(funcMap).ParseFiles("src/views/memberForm.html")
	if err != nil {
		fmt.Println("error parsing file memberForm.html")
	}
	// tmpl, _ := template.ParseFiles("src/views/memberForm.html", "src/views/footer.html")
	// el primero siempre es el main template, los demas se usan como componentes
	tmpl.Execute(w, Info{"create", models.Member{}}) // le paso un miembro vacio
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

func ShowOrNot(info Info, field string) string {
	if info.Action == "edit" {
		switch field {
		case "Name":
			return info.Member.Name
		case "DNI":
			return info.Member.DNI
		default:
			return "error no field"
		}
	} else if info.Action == "create" {
		return ""
	} else {
		return "big error"
	}
}
