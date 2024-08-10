package controller

import (
	"github.com/gofiber/fiber/v2"
	// "syscall/js"
)

// ------------------------------------

func RenderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
	// tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	// return tmpl.Execute(c, nil)
}
