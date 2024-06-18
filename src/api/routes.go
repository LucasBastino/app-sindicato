package api

import (
	"net/http"
)

func (c *Controller) RegisterRoutes(muxer *http.ServeMux) {
	muxer.HandleFunc("GET /index", c.renderIndex)

	muxer.HandleFunc("GET /memberTable", c.renderMemberTable)
	muxer.HandleFunc("GET /form/createMember", c.renderCreateMemberForm)
	muxer.HandleFunc("POST /member/create", c.createMember)
	muxer.HandleFunc("GET /form/editMember/{IdMember}", c.renderEditMemberForm)
	muxer.HandleFunc("POST /member/{IdMember}/edit", c.editMember)
	muxer.HandleFunc("DELETE /member/{IdMember}/delete", c.deleteMember)

	// muxer.HandleFunc("GET /enterPriseTable", c.renderEnterpriseTable)
	// muxer.HandleFunc("GET /parentTable", c.renderParentTable)

	// muxer.HandleFunc("GET /form/createParent", c.renderCreateParentForm)
}
