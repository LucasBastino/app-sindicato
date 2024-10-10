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

	err1 := validateDateFormat(birthday)
	if err1 != "" {
		return err1
	}
	err2 := validateDateValue(birthday)
	if err2 != "" {
		return err2
	}
	return ""
}

func ValidateGender(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("gender")) == "" {
		return "Elegir un género"
	}
	return ""
}

func ValidateMaritalStatus(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("marital-status")) == "" {
		return "Elegir un estado civil"
	}
	return ""
}

func ValidatePhone(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("phone")) == "" {
		return "El campo Teléfono no puede estar vacío"
	}
	return ""
}

func ValidateEmail(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("email")) == "" {
		return "El campo E-mail no puede estar vacío"
	}
	return ""
}

func ValidateAddress(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("address")) == "" {
		return "El campo Dirección no puede estar vacío"
	}
	return ""
}

func ValidatePostalCode(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("postal-code")) == "" {
		return "El campo Codigo postal no puede estar vacío"
	}
	return ""
}

func ValidateDistrict(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("district")) == "" {
		return "El campo Localidad no puede estar vacío"
	}
	return ""
}

func ValidateMemberNumber(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("member-number")) == "" {
		return "El campo Numero de afiliado no puede estar vacío"
	}
	return ""
}

func ValidateCUIL(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("cuil")) == "" {
		return "El campo CUIL no puede estar vacío"
	}
	return ""
}

func ValidateIdEnterprise(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("id-enterprise")) == "" {
		return "Elegir una empresa"
	}
	return ""
}

func ValidateCategory(c *fiber.Ctx) string {
	if strings.TrimSpace(c.FormValue("category")) == "" {
		return "Elegir una categoría"
	}
	return ""
}

func ValidateEntryDate(c *fiber.Ctx) string {
	entryDate := c.FormValue("entry-date")
	if strings.TrimSpace(entryDate) == "" {
		return "El campo Fecha de ingreso no puede estar vacío"
	}
	err1 := validateDateFormat(entryDate)
	if err1 != "" {
		return err1
	}
	err2 := validateDateValue(entryDate)
	if err2 != "" {
		return err2
	}
	return ""
}

func getDay(date string) (int, error) {
	return strconv.Atoi(date[0:2])
}

func getMonth(date string) (int, error) {
	return strconv.Atoi(date[3:5])
}

func getYear(date string) (int, error) {
	return strconv.Atoi(date[6:])
}

func validateDateFormat(date string) string {
	if len(date) != 10 {
		return "Formato de fecha erróneo"
	}
	_, dayErr := getDay(date)
	_, monthErr := getMonth(date)
	_, yearErr := getYear(date)
	if dayErr != nil || monthErr != nil || yearErr != nil ||
		string(date[2]) != "/" || string(date[5]) != "/" {
		return "Formato de fecha erróneo"
	}
	return ""
}

func validateDateValue(date string) string {
	day, _ := getDay(date)
	month, _ := getDay(date)
	year, _ := getDay(date)

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
