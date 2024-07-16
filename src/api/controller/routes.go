package api

import (
	"net/http"
)

func (c *Controller) RegisterRoutes(muxer *http.ServeMux) {
	muxer.HandleFunc("GET /index", c.renderIndex)

	muxer.HandleFunc("GET /memberTable", c.renderMemberTable)
	muxer.HandleFunc("GET /form/createMember", c.renderCreateMemberForm)
	muxer.HandleFunc("POST /member/create", c.createMember)
	muxer.HandleFunc("GET /member/{IdMember}/file", c.renderMemberFile)
	muxer.HandleFunc("GET /member/{IdMember}/parentTable", c.renderParentTable)
	muxer.HandleFunc("PUT /member/{IdMember}/edit", c.editMember)
	muxer.HandleFunc("DELETE /member/{IdMember}/delete", c.deleteMember)
	muxer.HandleFunc("POST /searchMember", c.searchMember)

	muxer.HandleFunc("GET /allParentsTable", c.renderAllParentsTable)
	muxer.HandleFunc("POST /parent/create", c.createParent)
	muxer.HandleFunc("GET /parent/{IdParent}/file", c.renderParentFile)
	muxer.HandleFunc("DELETE /parent/{IdParent}", c.deleteParent)
	muxer.HandleFunc("PUT /parent/{IdParent}/edit", c.editParent)
	muxer.HandleFunc("POST /searchParent", c.searchParent)

	muxer.HandleFunc("GET /form/createEnterprise", c.renderCreateEnterpriseForm)
	muxer.HandleFunc("POST /enterprise/create", c.createEnterprise)
	muxer.HandleFunc("GET /enterprise/{IdEnterprise}/file", c.renderEnterpriseFile)
	muxer.HandleFunc("DELETE /enterprise/{IdEnterprise}/delete", c.deleteEnterprise)
	// cambiar el de abajo a PUT
	muxer.HandleFunc("PUT /enterprise/{IdEnterprise}/edit", c.editEnterprise)
	muxer.HandleFunc("GET /enterpriseTable", c.renderEnterpriseTable)
	muxer.HandleFunc("POST /searchEnterprise", c.searchEnterprise)

	// muxer.HandleFunc("GET /enterPriseTable", c.renderEnterpriseTable)
	// muxer.HandleFunc("GET /parentTable", c.renderParentTable)

	// muxer.HandleFunc("GET /form/createParent", c.renderCreateParentForm)
}
