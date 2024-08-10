package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
	// "syscall/js"
)

// ------------------------------------

func RenderIndex(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	tmpl.Execute(w, nil)
}

func RenderHTML(w http.ResponseWriter, templateData models.TemplateData) {
	tmpl, err := template.ParseFiles(templateData.Path)
	if err != nil {
		fmt.Println("error parsing file", templateData.Path)
		panic(err)
	}
	tmpl.Execute(w, templateData)
}
