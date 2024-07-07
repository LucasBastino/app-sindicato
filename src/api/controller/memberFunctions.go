package api

import (
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) createMember(w http.ResponseWriter, r *http.Request) {
	memberParser := i.MemberParser{}
	newMember := parserCaller(memberParser, r)
	insertInDBCaller(newMember, c.DB)
	renderTemplateCaller(newMember, w, "src/views/files/memberFile.html")

	// http.Redirect(w, r, "/index", http.StatusSeeOther) // poner un status de redirect (30X), sino no funciona
	// c.renderMemberList(w, r) // esto tambien funciona
}

func (c *Controller) deleteMember(w http.ResponseWriter, r *http.Request) {
	IdMember := getIdModel("Member", r)
	deleteFromDBCaller(models.Member{IdMember: IdMember}, c.DB)
	c.renderMemberTable(w, r)
}

func (c *Controller) editMember(w http.ResponseWriter, r *http.Request) {
	memberParser := i.MemberParser{}
	memberEdited := parserCaller(memberParser, r)
	IdMember := getIdModel("Member", r)
	updateInDBCaller(memberEdited, IdMember, c.DB)

	// no puedo hacer esto ↓ porque estoy en POST, no puedo redireccionar
	http.Redirect(w, r, "/index", http.StatusSeeOther) // con este status me anda, con otros de 300 no
}

func (c *Controller) searchMember(w http.ResponseWriter, r *http.Request) {
	ms := i.MemberSearcher{}

	members := searcherCaller(ms, r, c.DB)

}
