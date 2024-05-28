package api

import (
	"net/http"
)

func (c *Controller) RegisterRoutes(muxer *http.ServeMux) {
	muxer.HandleFunc("GET /home", c.renderHome)
	muxer.HandleFunc("GET /users", c.getUsers)
	muxer.HandleFunc("GET /users/", c.getUsers)
	muxer.HandleFunc("POST /users/insert/{name}/{age}", c.insertUser)
	muxer.HandleFunc("PUT /users/edit/{id}", c.updateUser)
	muxer.HandleFunc("PUT /users/edit", c.updateUser)

	muxer.HandleFunc("POST /test", c.test)
	muxer.HandleFunc("POST /createTable", c.createTable)
}
