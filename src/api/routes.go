package api

import (
	"database/sql"
	"net/http"
)

func (c *Controller) RegisterRoutes(r *http.ServeMux, db *sql.DB) {
	r.HandleFunc("GET /", c.renderHome)
	r.HandleFunc("GET /users", c.getUsers)
	r.HandleFunc("GET /users/", c.getUsers)
	r.HandleFunc("POST /users/insert/{:nombre}/{:edad}", c.insertUser)
}
