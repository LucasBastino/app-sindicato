package api

import (
	"fmt"
	"html/template"
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) createParent(w http.ResponseWriter, r *http.Request) {
	parentParser := i.ParentParser{}
	newParent := parserCaller(parentParser, r)
	insertModelCaller(newParent, c.DB)
	renderFileTemplateCaller(newParent, w, "src/views/files/parentFile.html")
}

func (c *Controller) deleteParent(w http.ResponseWriter, r *http.Request) {
	IdParent := getIdModelCaller(models.Parent{}, r)
	deleteModelCaller(models.Parent{IdParent: IdParent}, c.DB)
	c.renderParentTable(w, r)
}

func (c *Controller) editParent(w http.ResponseWriter, r *http.Request) {
	parser := i.ParentParser{}
	parentEdited := parserCaller(parser, r)
	IdParent := getIdModelCaller(models.Parent{}, r)
	parentEdited.IdParent = IdParent
	editModelCaller(parentEdited, IdParent, c.DB)
	renderFileTemplateCaller(parentEdited, w, "src/views/files/parentFile.html")
}

func (c *Controller) renderParentTable(w http.ResponseWriter, r *http.Request) {
	parents := searchModelsCaller(models.Parent{}, r, c.DB)
	renderTableTemplateCaller(models.Parent{}, w, "src/views/tables/parentTable.html", parents)
}

func (c *Controller) renderMemberParents(w http.ResponseWriter, r *http.Request) {
	IdMember := getIdModelCaller(models.Member{}, r)
	result, err := c.DB.Query(fmt.Sprintf("SELECT Name, Rel, IdParent FROM ParentTable WHERE IdMember = '%s'", IdMember))
	if err != nil {
		fmt.Println("error searching parents from db")
		panic(err)
	}

	// hacer un metodo para scan
	var parent models.Parent
	var parents []models.Parent
	for result.Next() {
		err = result.Scan(&parent.Name, &parent.Rel, &parent.IdParent)
		if err != nil {
			fmt.Println("error scanning parent")
			panic(err)
		}
		parents = append(parents, parent)
	}

	path := "src/views/tables/parentTable.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, parents)

}
