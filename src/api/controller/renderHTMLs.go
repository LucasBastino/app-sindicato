package api

import (
	"html/template"
	"net/http"
)

func (c *Controller) renderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	tmpl.Execute(w, nil)
}
