package controller

import (
	"github.com/gofiber/fiber/v2"
	// "syscall/js"
)

// ------------------------------------

func RenderIndex(c *fiber.Ctx) error {
	return c.Render("prueba", fiber.Map{})
	// tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	// return tmpl.Execute(c, nil)
}

func RenderHTML(c *fiber.Ctx, data fiber.Map) error {
	return c.Render("memberTable", data)

	// tmpl, err := template.ParseFiles(templateData.Path)
	// if err != nil {
	// 	fmt.Println("error parsing file", templateData.Path)
	// 	panic(err)
	// }
	// return tmpl.Execute(c, templateData)
}
