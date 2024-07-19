package api

import (
	"database/sql"
	"html/template"
	"net/http"
	// "syscall/js"
)

type Controller struct {
	DB *sql.DB
}

// ------------------------------------

func (c *Controller) renderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	tmpl.Execute(w, nil)
}
