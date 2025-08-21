package session

import (
	"github.com/LucasBastino/app-sindicato/src/config/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Session = session.New()

func SetError(c *fiber.Ctx) {
	sess := GetSession(c)
	sess.Set("ErrType", "DatabaseConnection")
	sess.Save()
}

func GetError(c *fiber.Ctx) interface{}{
	sess := GetSession(c)
	errType := sess.Get("ErrType")
	sess.Delete("ErrType")
	sess.Save()
	return errType
}

func GetSession(c *fiber.Ctx) *session.Session{
	sess, err := Session.Get(c)
	if err!=nil{
		logger.Log.Error("session error")
	}
	return sess
}