package api

import (
	"net/http"
)

func (c *Controller) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", c.renderHome)
	r.HandleFunc("GET /users", c.getUsers)
	r.HandleFunc("GET /users/", c.getUsers)
	r.HandleFunc("POST /users/insert/{nombre}/{edad}", c.insertUser)
	r.HandleFunc("PUT /users/update/{id}", c.updateUser)
	r.HandleFunc("PUT /users/update", c.updateUser)
	// r.HandleFunc("POST /createTable", c.createTable)
}
