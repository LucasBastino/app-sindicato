package session

import (
	"github.com/LucasBastino/app-sindicato/src/config/logger"
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Session = session.New()

func SetError(c *fiber.Ctx, err *customError.CustomError) {
	sess := GetSession(c)
	sess.Set("Err", err)
	sess.Save()
}

func GetError(c *fiber.Ctx) *customError.CustomError{
	sess := GetSession(c)
	val := sess.Get("Err")
	sess.Delete("Err")
	sess.Save()

	if val == nil{
		return nil
	}

	// convierto el err tipo interface en tipo *CustomError
	err, ok := val.(*customError.CustomError)
	if !ok{
		// hacer algo
		return nil
	}
	return err
}

func GetSession(c *fiber.Ctx) *session.Session{
	sess, err := Session.Get(c)
	if err!=nil{
		logger.Log.Error("session error")
	}
	return sess
}