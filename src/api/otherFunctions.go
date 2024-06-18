package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

func createTemplate(path string) *template.Template {
	tmpl, err := template.New("newTemplate").ParseFiles(path)
	if err != nil {
		fmt.Println("error creating template from", path)
		log.Panic(err.Error())
	}
	return tmpl
}

func execTemplate(w http.ResponseWriter, data any, tmpl *template.Template, file string) {
	// w.Header().Set("Content-Type", "application/json")
	err := tmpl.ExecuteTemplate(w, file, data)
	if err != nil {
		fmt.Println("error executing template")
		log.Panic(err)
	}
}

func parseMember(r *http.Request) models.Member {
	var member models.Member
	member.Name = r.FormValue("name")
	member.DNI = r.FormValue("dni")
	return member
}

func returnHtmlTemplate(path string) *template.Template {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("error parsing", path)
		panic(err)
	}
	return tmpl
}
