package errors

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type CustomError struct {
	Msg string
}

func (e *CustomError) Error() string {
	return e.Msg
}

var QueryError = &CustomError{}
var ScanError = &CustomError{}
var FormatError = &CustomError{}
var ValidationError = &CustomError{}
var UnauthorizedError = &CustomError{}
var InsufficientPermisionsError = &CustomError{}
var InternalServerError = &CustomError{}
var StrConvError = &CustomError{}
var DatabaseConnectionError = &CustomError{}
var ParamsError = &CustomError{}

func CheckError(c *fiber.Ctx, err error) error {
	if errors.Is(err, QueryError) {
		// guardar el log en algun lugar
		// os.Guardar("errors.txt", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error interno de la base de datos"})
	}
	if errors.Is(err, ScanError) {
		// guardar el log en algun lugar
		// os.Guardar("errors.txt", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error funcional interno de la base de datos"})
	}
	return err
}
