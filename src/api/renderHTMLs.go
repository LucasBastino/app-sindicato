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
	tmpl := createTemplate("src/views/createMemberForm.html")
	execTemplate(w, nil, tmpl, "createMemberForm.html")
}

func (c *Controller) renderMemberList(w http.ResponseWriter, r *http.Request) {
	result, err := c.DB.Query("SELECT Name, DNI FROM MemberTable")
	if err != nil {
		fmt.Println("error obtaining data from database")
		log.Panic(err)
	}

	var members []models.Member
	for result.Next() {
		member := models.Member{}
		err := result.Scan(&member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning data")
			log.Panic(err)
		}
		members = append(members, member)
	}

	tmpl := createTemplate("src/views/memberList.html")
	execTemplate(w, members, tmpl, "memberList.html")
}
