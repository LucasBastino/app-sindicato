package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) renderHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/views/home.html"))
	tmpl.Execute(w, nil)
}

func (c *Controller) renderCreateMemberForm(w http.ResponseWriter, req *http.Request) {
	// tmpl := createTemplate("src/views/createMemberForm.html")
	// execTemplate(w, nil, tmpl, "createMemberForm.html")

	// tmpl, _ := template.New("createMemberForm.html").ParseFiles("src/views/createMemberForm.html")
	// tmpl.Execute(w, nil)

	tmpl, _ := template.ParseFiles("src/views/createMemberForm.html", "src/views/footer.html")
	// el primero siempre es el main template, los demas se usan como componentes
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

	type arguments struct {
		siono   bool
		members []models.Member
	}

	tmpl := returnHtmlTemplate("src/views/memberList.html")
	// tmpl.Execute(w, arguments{false, members})
	tmpl.Execute(w, members)
}

func (c *Controller) renderEditMemberForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("estoy en edit ")
	tmpl := returnHtmlTemplate("src/views/memberForm.html")
	tmpl.Execute(w, nil)
}
