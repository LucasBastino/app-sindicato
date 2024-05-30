package api

import (
	"net/http"
)

func (c *Controller) RegisterRoutes(muxer *http.ServeMux) {
	muxer.HandleFunc("GET /home", c.renderHome)

	muxer.HandleFunc("GET /members", c.renderMemberList)
	muxer.HandleFunc("POST /member/create", c.createMember)
	muxer.HandleFunc("GET /forms/createMember", c.renderCreateMemberForm)
}
