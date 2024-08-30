package controller

import (
	"github.com/gofiber/fiber/v2"
)

func TestNull(c *fiber.Ctx) error {
	return c.Render("test", fiber.Map{"variable": 1})
}
