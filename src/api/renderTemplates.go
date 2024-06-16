package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

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
