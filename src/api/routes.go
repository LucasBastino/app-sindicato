package api

import (
	"net/http"
)

func (c *Controller) RegisterRoutes(muxer *http.ServeMux) {
	muxer.HandleFunc("GET /index", c.renderIndex)
	muxer.HandleFunc("GET /renderMemberList", c.renderMemberList)
	muxer.HandleFunc("POST /member/create", c.createMember)
	muxer.HandleFunc("PUT /member/{IdMember}/edit", c.renderEditMemberForm)
	muxer.HandleFunc("DELETE /member/{IdMember}/delete", c.deleteMember)
	muxer.HandleFunc("GET /forms/createMember", c.renderCreateMemberForm)
	muxer.HandleFunc("GET /forms/editMember/{IdMember}", c.renderEditMemberForm)
}
