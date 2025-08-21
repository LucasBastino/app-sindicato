package errors

import (
	"errors"
	"fmt"

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


// hacer esto con una variable global?

func RenderError(c *fiber.Ctx) error {

	switch c.Cookies("ErrType") {
	case "Query":
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"error": "internal database error"})
	case "Scan":
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"error": "internal error"})
	case "Format":
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"error": "internal error"})
	case "Validation":
		return c.Status(fiber.StatusBadRequest).Render("error", fiber.Map{"error": "validation error"})
	case "Unauthorized":
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"error": "error interno de la base de datos"})
	case "InsufficientPermisions":
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"error": "error funcional interno de la base de datos"})
	case "InternalServer":
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"error": "error interno de la base de datos"})
	case "StrConv":
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"error": "error funcional interno de la base de datos"})
	case "DatabaseConnection":
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"error": "error interno de la base de datos"})
	case "Params":
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"error": "error funcional interno de la base de datos"})
	default:
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"error": "default internal error"})
	}

}

func CheckError(c *fiber.Ctx, err error) error {
	// faltar guardar el log del error

	// esto es temporal
	fmt.Println(err)
	cookie := fiber.Cookie{
		Name:        "ErrType",
		Value:       "",
		Path:        "/",
		Secure:      true,
		HTTPOnly:    true,
		SameSite:    "Strict",
		SessionOnly: true,
	}
	if errors.Is(err, QueryError) {
		cookie.Value = "Query"
	}
	if errors.Is(err, ScanError) {
		cookie.Value = "Scan"
	}
	if errors.Is(err, FormatError) {
		cookie.Value = "Format"
	}
	if errors.Is(err, ValidationError) {
		cookie.Value = "Validation"
	}
	if errors.Is(err, UnauthorizedError) {
		cookie.Value = "Unauthorized"
	}
	if errors.Is(err, InsufficientPermisionsError) {
		cookie.Value = "InsufficientPermisions"
	}
	if errors.Is(err, InternalServerError) {
		cookie.Value = "InternalServer"
	}
	if errors.Is(err, StrConvError) {
		cookie.Value = "StrConv"
	}
	if errors.Is(err, DatabaseConnectionError) {
		cookie.Value = "DatabaseConnection"
	}
	if errors.Is(err, ParamsError) {
		cookie.Value = "Params"
	}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusAccepted).Render("redirect", fiber.Map{"path": "/error"})
}
