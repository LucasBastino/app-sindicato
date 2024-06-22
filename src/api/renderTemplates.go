package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) renderMemberTable(w http.ResponseWriter, r *http.Request) {
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

	tmpl := returnHtmlTemplate("src/views/tables/memberTable.html")
	tmpl.Execute(w, members)
}

func (c *Controller) renderParentTable(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entro en renderparenttable")
	IdMember := r.PathValue("IdMember")
	result, err := c.DB.Query(fmt.Sprintf("SELECT Name, Rel FROM ParentTable WHERE IdMember = '%s'", IdMember))
	if err != nil {
		fmt.Println("error selecting data from database")
		panic(err)
	}

	fmt.Println("ya hizo el query")
	var parent models.Parent
	var parents []models.Parent
	for result.Next() {
		err = result.Scan(&parent.Name, &parent.Rel)
		if err != nil {
			fmt.Println("error scanning data")
			panic(err)
		}
		parents = append(parents, parent)
	}
	fmt.Println("paso los resultados a parents")
	fmt.Println(parents)
	tmpl, err := template.ParseFiles("src/views/tables/parentMemberTable.html")
	if err != nil {
		fmt.Println("error parsing file parentMemberTable.html")
		panic(err)
	}
	tmpl.Execute(w, parents)

}
