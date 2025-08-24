package errorHandler

import (
	"github.com/LucasBastino/app-sindicato/src/config/logger"
	"github.com/LucasBastino/app-sindicato/src/config/session"
	"github.com/LucasBastino/app-sindicato/src/errors/customError"

	"github.com/gofiber/fiber/v2"
)


func RenderError(c *fiber.Ctx) error{
	err := session.GetError(c)
	logger.Log.Error(err.Msg)
	return c.Status(err.Code).Render("error", err.ClientMsg)
}

func HandleError(c *fiber.Ctx, err *customError.CustomError) error{
	session.SetError(c, err)
	return c.Status(fiber.StatusFound).Render("redirect", fiber.Map{"path": "/error"})
}
