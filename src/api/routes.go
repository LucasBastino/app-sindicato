package api

import (
	"net/http"
)

func (c *Controller) RegisterRoutes(muxer *http.ServeMux) {
	muxer.HandleFunc("GET /home", c.renderHome)

	muxer.HandleFunc("GET /members", c.renderMemberList)
	muxer.HandleFunc("GET /members2", c.renderMemberList2)
	muxer.HandleFunc("GET /getList", c.getList)
	muxer.HandleFunc("POST /member/create", c.createMember)
	muxer.HandleFunc("DELETE /member/{IdMember}/delete", c.deleteMember)
	muxer.HandleFunc("GET /forms/createMember", c.renderCreateMemberForm)
}
