package api

import (
	"net/http"
)

func (c *Controller) RegisterRoutes(muxer *http.ServeMux) {
	muxer.HandleFunc("GET /index", c.renderIndex)

	muxer.HandleFunc("GET /renderMemberList", c.renderMemberList)

	muxer.HandleFunc("GET /forms/createMember", c.renderCreateMemberForm)
	muxer.HandleFunc("GET /forms/editMember/{IdMember}", c.renderEditMemberForm)

	muxer.HandleFunc("POST /member/create", c.createMember)
	muxer.HandleFunc("POST /member/{IdMember}/edit", c.editMember)
	muxer.HandleFunc("DELETE /member/{IdMember}/delete", c.deleteMember)
}
