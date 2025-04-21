package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CustomError struct {
	Msg  string
	Code int
}

func (e *CustomError) Error() string {
	return e.Msg
}

var QueryError = &CustomError{Code: 1}
var ScanError = &CustomError{Code: 2}
var FormatError = &CustomError{Code: 3}
var ValidationError = &CustomError{Code: 4}
var UnauthorizedError = &CustomError{Code: 5}
var InsufficientPermisionsError = &CustomError{Code: 6}
var InternalServerError = &CustomError{Code: 7}
var StrConvError = &CustomError{Code: 8}
var DatabaseConnectionError = &CustomError{Code: 9}
var ParamsError = &CustomError{Code: 10}

func CheckError(c *fiber.Ctx, err *CustomError) error {
	// guardar el error
	c.Locals("err", err.Code)
	fmt.Println("hola")
	fmt.Println(c.Locals("err"))
	fmt.Println("hola")
	return c.Status(fiber.StatusAccepted).Render("redirect", fiber.Map{"path": "/error", "error": "internal error"})

}

// func CheckError2(c *fiber.Ctx) error {
// 	err := c.Locals("err")

// 	if errors.Is(err, QueryError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal database error"})
// 	}
// 	if errors.Is(err, ScanError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		// SI PONES STATUS INTERNAL SERVER ERROR NO TE DEJA DIRECCIONAR, asi que el status se lo pongo directamente cuando renderizo el error.html
// 		return c.Status(fiber.StatusInternalServerError).Render("redirect", fiber.Map{"path": "/error", "error": "internal error"})
// 	}
// 	if errors.Is(err, FormatError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal error"})
// 	}
// 	if errors.Is(err, ValidationError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation error"})
// 	}
// 	if errors.Is(err, UnauthorizedError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error interno de la base de datos"})
// 	}
// 	if errors.Is(err, InsufficientPermisionsError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error funcional interno de la base de datos"})
// 	}
// 	if errors.Is(err, InternalServerError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error interno de la base de datos"})
// 	}
// 	if errors.Is(err, StrConvError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error funcional interno de la base de datos"})
// 	}
// 	if errors.Is(err, DatabaseConnectionError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error interno de la base de datos"})
// 	}
// 	if errors.Is(err, ParamsError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error funcional interno de la base de datos"})
// 	}
// 	return err
// }

// func CheckError(c *fiber.Ctx, err error) error {
// 	if errors.Is(err, QueryError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal database error"})
// 	}
// 	if errors.Is(err, ScanError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		// SI PONES STATUS INTERNAL SERVER ERROR NO TE DEJA DIRECCIONAR, asi que el status se lo pongo directamente cuando renderizo el error.html
// 		return c.Status(fiber.StatusInternalServerError).Render("redirect", fiber.Map{"path": "/error", "error": "internal error"})
// 	}
// 	if errors.Is(err, FormatError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal error"})
// 	}
// 	if errors.Is(err, ValidationError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation error"})
// 	}
// 	if errors.Is(err, UnauthorizedError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error interno de la base de datos"})
// 	}
// 	if errors.Is(err, InsufficientPermisionsError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error funcional interno de la base de datos"})
// 	}
// 	if errors.Is(err, InternalServerError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error interno de la base de datos"})
// 	}
// 	if errors.Is(err, StrConvError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error funcional interno de la base de datos"})
// 	}
// 	if errors.Is(err, DatabaseConnectionError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error interno de la base de datos"})
// 	}
// 	if errors.Is(err, ParamsError) {
// 		// guardar el log en algun lugar
// 		// os.Guardar("errors.txt", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error funcional interno de la base de datos"})
// 	}
// 	return err
// }
