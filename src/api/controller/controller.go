package api

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
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

func (c *Controller) RenderHTML(w http.ResponseWriter, templateData models.TemplateData) {
	tmpl, err := template.ParseFiles(templateData.Path)
	if err != nil {
		fmt.Println("error parsing file", templateData.Path)
		panic(err)
	}
	tmpl.Execute(w, templateData)
}
