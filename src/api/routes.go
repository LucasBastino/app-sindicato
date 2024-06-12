package api

import (
	"net/http"
)

func (c *Controller) RegisterRoutes(muxer *http.ServeMux) {
	muxer.HandleFunc("GET /home", c.renderHome)

	muxer.HandleFunc("GET /members", c.renderMembers1)
	muxer.HandleFunc("GET /members2", c.renderMembers2)
	muxer.HandleFunc("GET /renderMemberList", c.renderMemberList)
	muxer.HandleFunc("POST /member/create", c.createMember)
	muxer.HandleFunc("PUT /member/{IdMember}/edit", c.renderEditMemberForm)
	muxer.HandleFunc("DELETE /member/{IdMember}/delete", c.deleteMember)
	muxer.HandleFunc("GET /forms/createMember", c.renderCreateMemberForm)
}
