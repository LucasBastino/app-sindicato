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

	path := "src/views/tables/memberTable.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, members)
}

func (c *Controller) renderParentTable(w http.ResponseWriter, r *http.Request) {
	IdMember := r.PathValue("IdMember")
	result, err := c.DB.Query(fmt.Sprintf("SELECT Name, Rel, IdParent FROM ParentTable WHERE IdMember = '%s'", IdMember))
	if err != nil {
		fmt.Println("error selecting data from database")
		panic(err)
	}

	var parent models.Parent
	var parents []models.Parent
	for result.Next() {
		err = result.Scan(&parent.Name, &parent.Rel, &parent.IdParent)
		if err != nil {
			fmt.Println("error scanning data")
			panic(err)
		}
		parents = append(parents, parent)
	}

	tmpl, err := template.ParseFiles("src/views/tables/parentTable.html")
	if err != nil {
		fmt.Println("error parsing file parentTable.html")
		panic(err)
	}
	tmpl.Execute(w, parents)

}

func (c *Controller) renderAllParentsTable(w http.ResponseWriter, r *http.Request) {
	var parents []models.Parent
	var parent models.Parent

	result, err := c.DB.Query("SELECT Name, Rel FROM ParentTable")
	if err != nil {
		fmt.Println("error searching parents")
		panic(err)
	}

	for result.Next() {
		err = result.Scan(&parent.Name, &parent.Rel)
		if err != nil {
			fmt.Println("error scanning data")
			panic(err)
		}
		parents = append(parents, parent)
	}

	tmpl, err := template.ParseFiles("src/views/tables/allParentsTable.html")
	if err != nil {
		fmt.Println("error parsing file allParentsTable.html")
		panic(err)
	}
	tmpl.Execute(w, parents)
}
