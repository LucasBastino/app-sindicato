package controller

import (
	"fmt"
	"html/template"

	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	// "syscall/js"
)

// ------------------------------------

func RenderIndex(c *fiber.Ctx) {
	tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	tmpl.Execute(w, nil)
}

func RenderHTML(c *fiber.Ctx, templateData models.TemplateData) {
	tmpl, err := template.ParseFiles(templateData.Path)
	if err != nil {
		fmt.Println("error parsing file", templateData.Path)
		panic(err)
	}
	tmpl.Execute(w, templateData)
}
