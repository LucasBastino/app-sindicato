package api

import (
	"fmt"
	"html/template"
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

var (
	member       models.Member
	memberParser i.MemberParser
)

func (c *Controller) createMember(w http.ResponseWriter, r *http.Request) {
	errorsMap := validateFieldsCaller(member, r)
	if len(errorsMap) > 0 {
		templateData := createTemplateDataCaller(member, member, nil, "src/views/forms/createMemberForm.html", errorsMap)
		renderCreateModelFormCaller(member, w, templateData)
	}

	// if validateFieldsCaller(member, r) {
	// 	memberParser := i.MemberParser{}
	// 	newMember := parserCaller(memberParser, r)
	// 	insertModelCaller(newMember, c.DB)
	// 	// hacer esto esta bien? estoy mostrando datos del newMember, no estan sacados de la DB
	// renderFileTemplateCaller(newMember, w, "src/views/files/memberFile.html")
	// } else {
	// 	renderFileTemplateCaller(member, w, "src/views/files/memberFileError.html")
	// }

	// memberParser := i.MemberParser{}
	// newMember := parserCaller(memberParser, r)
	// insertModelCaller(newMember, c.DB)
	// renderFileTemplateCaller(newMember, w, "src/views/files/memberFile.html")

	// http.Redirect(w, r, "/index", http.StatusSeeOther) // poner un status de redirect (30X), sino no funciona
	// c.renderMemberList(w, r) // esto tambien funciona
}

func (c *Controller) deleteMember(w http.ResponseWriter, r *http.Request) {
	IdMember := getIdModelCaller(member, r)
	member.IdMember = IdMember
	deleteModelCaller(member, c.DB)
	allMembers := searchModelsCaller(member, r, c.DB)
	templateData := createTemplateDataCaller(member, member, allMembers, "src/views/tables/memberTable.html", nil)
	renderTableTemplateCaller(member, w, templateData)
}

func (c *Controller) editMember(w http.ResponseWriter, r *http.Request) {
	member = parserCaller(memberParser, r)
	IdMember := getIdModelCaller(member, r)
	member.IdMember = IdMember
	// necesito poner esta linea ↑ para que se pueda editar 2 veces seguidas
	editModelCaller(member, IdMember, c.DB)
	// hacer esto esta bien? estoy mostrando datos del newMember, no estan sacados de la DB
	templateData := createTemplateDataCaller(member, member, nil, "src/views/files/memberFile.html", nil)
	renderFileTemplateCaller(member, w, templateData)

	// no puedo hacer esto ↓ porque estoy en POST, no puedo redireccionar
	// http.Redirect(w, r, "/index", http.StatusSeeOther) // con este status me anda, con otros de 300 no
}

func (c *Controller) renderMemberTable(w http.ResponseWriter, r *http.Request) {
	members := searchModelsCaller(member, r, c.DB)
	templateData := createTemplateDataCaller(member, member, members, "src/views/tables/memberTable.html", nil)
	renderTableTemplateCaller(member, w, templateData)
}

func (c *Controller) renderMemberFile(w http.ResponseWriter, r *http.Request) {
	member := searchOneModelByIdCaller(member, r, c.DB)
	templateData := createTemplateDataCaller(member, member, nil, "src/views/files/memberFile.html", nil)
	renderFileTemplateCaller(member, w, templateData)
}

func (c *Controller) renderCreateMemberForm(w http.ResponseWriter, req *http.Request) {
	templateData := createTemplateDataCaller(member, member, nil, "src/views/forms/createMemberForm.html", nil)
	renderCreateModelFormCaller(member, w, templateData)
}

func (c *Controller) renderMemberParents(w http.ResponseWriter, r *http.Request) {
	IdMember := getIdModelCaller(member, r)
	result, err := c.DB.Query(fmt.Sprintf("SELECT Name, Rel, IdParent FROM ParentTable WHERE IdMember = '%d'", IdMember))
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

	path := "src/views/tables/memberParentTable.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, parents)

}
