package models

import (
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
)

func ValidateName(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("name")) == "" {
		return "El campo Nombre no puede estar vacío"
	}
	return ""
}

func ValidateLastName(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("last-name")) == "" {
		return "El campo Apellido no puede estar vacío"
	}
	return ""
}

func ValidateDNI(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("dni")) == "" {
		return "El campo DNI no puede estar vacío"
	}
	if utf8.RuneCountInString(c.FormValue("dni")) > 8 {
		return "El DNI no puede tener mas de 8 caracteres"
	}
	// chequear que sean numeros
	return ""
}

func ValidateBirthday(c *fiber.Ctx) string {
	birthday := c.FormValue("birthday")
	if strings.TrimSpace(birthday) == "" {
		return "El campo Fecha de nacimiento no puede estar vacío"
	}
	if len(birthday) != 10 {
		return "Formato de fecha erróneo"
	}
	day, dayErr := strconv.Atoi(birthday[0:2])
	month, monthErr := strconv.Atoi(birthday[3:5])
	year, yearErr := strconv.Atoi(birthday[6:])
	if dayErr != nil || monthErr != nil || yearErr != nil ||
		string(birthday[2]) != "/" || string(birthday[5]) != "/" {
		return "Formato de fecha erróneo"
	}
	if month < 1 || month > 12 {
		return "Fecha errónea"
	}
	if day < 1 || day > 31 {
		return "Fecha errónea"
	}
	switch month {
	case 2:
		if day > 29 {
			return "Fecha errónea"
		}
	case 4, 6, 9, 11:
		if day > 30 {
			return "Fecha errónea"
		}
	}
	if year < 1900 || year > int(time.Now().Year()) {
		return "Fecha errónea"
	}
	return ""
}
