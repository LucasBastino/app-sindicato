package api

import (
	"fmt"
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

var (
	member       models.Member
	memberParser i.MemberParser
)

func (c *Controller) createMember(w http.ResponseWriter, r *http.Request) {
	errorMap := validateFieldsCaller(member, r)
	parser := i.MemberParser{}
	member = parserCaller(parser, r)
	if len(errorMap) > 0 {
		templateData := createTemplateDataCaller(member, member, nil, "src/views/files/memberFile.html", errorMap)
		c.RenderHTML(w, templateData)
	} else {
		insertModelCaller(member, c.DB)
		templateData := createTemplateDataCaller(member, member, nil, "src/views/files/memberFile.html", errorMap)
		c.RenderHTML(w, templateData)
	}
}

func (c *Controller) deleteMember(w http.ResponseWriter, r *http.Request) {
	IdMember := getIdModelCaller(member, r)
	member.IdMember = IdMember
	deleteModelCaller(member, c.DB)
	c.renderMemberTable(w, r)
}

func (c *Controller) editMember(w http.ResponseWriter, r *http.Request) {
	member = parserCaller(memberParser, r)
	IdMember := getIdModelCaller(member, r)
	member.IdMember = IdMember
	// necesito poner esta linea ↑ para que se pueda editar 2 veces seguidas
	editModelCaller(member, c.DB)
	// hacer esto esta bien? estoy mostrando datos del newMember, no estan sacados de la DB
	templateData := createTemplateDataCaller(member, member, nil, "src/views/files/memberFile.html", nil)
	c.RenderHTML(w, templateData)

	// no puedo hacer esto ↓ porque estoy en POST, no puedo redireccionar
	// http.Redirect(w, r, "/index", http.StatusSeeOther) // con este status me anda, con otros de 300 no
}

func (c *Controller) renderMemberTable(w http.ResponseWriter, r *http.Request) {
	members := searchModelsCaller(member, r, c.DB)
	templateData := createTemplateDataCaller(member, member, members, "src/views/tables/memberTable.html", nil)
	c.RenderHTML(w, templateData)
}

func (c *Controller) renderMemberFile(w http.ResponseWriter, r *http.Request) {
	member := searchOneModelByIdCaller(member, r, c.DB)
	templateData := createTemplateDataCaller(member, member, nil, "src/views/files/memberFile.html", nil)
	c.RenderHTML(w, templateData)
}

func (c *Controller) renderCreateMemberForm(w http.ResponseWriter, req *http.Request) {
	templateData := createTemplateDataCaller(member, member, nil, "src/views/forms/createMemberForm.html", nil)
	c.RenderHTML(w, templateData)
}

func (c *Controller) renderMemberParents(w http.ResponseWriter, r *http.Request) {
	IdMember := getIdModelCaller(member, r)
	result, err := c.DB.Query(fmt.Sprintf("SELECT Name, Rel, IdParent, IdMember FROM ParentTable WHERE IdMember = '%d'", IdMember))
	if err != nil {
		fmt.Println("error searching parents from db")
		panic(err)
	}

	// hacer un metodo para scan
	var parent models.Parent
	var parents []models.Parent
	for result.Next() {
		err = result.Scan(&parent.Name, &parent.Rel, &parent.IdParent, &parent.IdMember)
		if err != nil {
			fmt.Println("error scanning parent")
			panic(err)
		}
		parents = append(parents, parent)
	}

	templateData := createTemplateDataCaller(parent, parent, parents, "src/views/tables/memberParentTable.html", nil)
	c.RenderHTML(w, templateData)

}
