package api

import (
	"html/template"
	"net/http"
)

func (c *Controller) renderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	tmpl.Execute(w, nil)
}

func (c *Controller) renderCreateMemberForm(w http.ResponseWriter, req *http.Request) {
	// creo el template para crear un afiliado y lo ejecuto
	path := "src/views/forms/createMemberForm.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, nil) // le paso un member vacio, no se puede pasar nil
}

func (c *Controller) renderCreateParentForm(w http.ResponseWriter, r *http.Request) {
	path := "src/views/forms/createParentForm.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, nil)
}

func (c *Controller) renderCreateEnterpriseForm(w http.ResponseWriter, r *http.Request) {
	path := "src/views/forms/createEnterpriseForm.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, nil)
}
