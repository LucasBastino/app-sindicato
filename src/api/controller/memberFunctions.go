package api

import (
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) createMember(w http.ResponseWriter, r *http.Request) {
	memberParser := i.MemberParser{}
	newMember := parserCaller(memberParser, r)
	insertModelCaller(newMember, c.DB)
	// hacer esto esta bien? estoy mostrando datos del newMember, no estan sacados de la DB
	renderFileTemplateCaller(newMember, w, "src/views/files/memberFile.html")

	// http.Redirect(w, r, "/index", http.StatusSeeOther) // poner un status de redirect (30X), sino no funciona
	// c.renderMemberList(w, r) // esto tambien funciona
}

func (c *Controller) deleteMember(w http.ResponseWriter, r *http.Request) {
	IdMember := getIdModelCaller(models.Member{}, r)
	deleteModelCaller(models.Member{IdMember: IdMember}, c.DB)
	allMembers := searchAllModelsCaller(models.Member{}, c.DB)
	renderTableTemplateCaller(models.Member{}, w, "src/views/tables/memberTable.html", allMembers)
}

func (c *Controller) editMember(w http.ResponseWriter, r *http.Request) {
	memberParser := i.MemberParser{}
	memberEdited := parserCaller(memberParser, r)
	IdMember := getIdModelCaller(models.Member{}, r)
	memberEdited.IdMember = IdMember
	// necesito hacer esto ↑ para que se pueda editar 2 veces seguidas
	editModelCaller(memberEdited, IdMember, c.DB)
	// hacer esto esta bien? estoy mostrando datos del newMember, no estan sacados de la DB
	renderFileTemplateCaller(memberEdited, w, "src/views/files/memberFile.html")

	// no puedo hacer esto ↓ porque estoy en POST, no puedo redireccionar
	// http.Redirect(w, r, "/index", http.StatusSeeOther) // con este status me anda, con otros de 300 no
}

func (c *Controller) renderMemberTable(w http.ResponseWriter, r *http.Request) {
	allMembers := searchAllModelsCaller(models.Member{}, c.DB)
	renderTableTemplateCaller(models.Member{}, w, "src/views/tables/memberTable.html", allMembers)
}

func (c *Controller) renderMemberFile(w http.ResponseWriter, r *http.Request) {
	member := searchOneModelByIdCaller(models.Member{}, r, c.DB)

}

func (c *Controller) searchMember(w http.ResponseWriter, r *http.Request) {
	members := searchInDBCaller(models.Member{}, r, c.DB)
	renderTableTemplateCaller(models.Member{}, w, "src/views/tables/memberTable.html", members)
}
