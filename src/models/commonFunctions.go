package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
)

func ValidateName(c *fiber.Ctx) string {
	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return "El campo Nombre no puede estar vacío"
	}
	return isLetter(name, "")
}

func ValidateLastName(c *fiber.Ctx) string {
	lastName := strings.TrimSpace(c.FormValue("last-name"))
	if lastName == "" {
		return "El campo Apellido no puede estar vacío"
	}
	return isLetter(lastName, "")
}

func ValidateEnterpriseName(c *fiber.Ctx) string {
	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return "El campo Nombre no puede estar vacío"
	}
	return isAlphanumeric(name, ".")
}

func ValidateDNI(c *fiber.Ctx) string {
	dni := strings.TrimSpace(c.FormValue("dni"))
	if dni == "" {
		return "El campo DNI no puede estar vacío"
	}
	if utf8.RuneCountInString(dni) > 8 {
		return "El DNI no puede tener mas de 8 caracteres"
	}
	return isNumber(dni, "")
}

func ValidateBirthday(c *fiber.Ctx) string {
	birthday := strings.TrimSpace(c.FormValue("birthday"))
	if birthday == "" {
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
	gender := strings.TrimSpace(c.FormValue("gender"))
	if gender == "" {
		return "Elegir un género"
	}
	return ""
}

func ValidateRel(c *fiber.Ctx) string {
	rel := strings.TrimSpace(c.FormValue("rel"))
	if rel == "" {
		return "El campo Parentesco no puede estar vacío"
	}
	return isLetter(rel, "")
}

func ValidateMaritalStatus(c *fiber.Ctx) string {
	maritalStatus := strings.TrimSpace(c.FormValue("marital-status"))
	if maritalStatus == "" {
		return "Elegir un estado civil"
	}
	return ""
}

func ValidatePhone(c *fiber.Ctx) string {
	phone := strings.TrimSpace(c.FormValue("phone"))
	if phone == "" {
		return "El campo Teléfono no puede estar vacío"
	}
	return isNumber(phone, "+")
}

func ValidateEmail(c *fiber.Ctx) string {
	// email puede estar vacio
	email := strings.TrimSpace(c.FormValue("email"))
	if email == "" {
		return ""
	} else if !strings.Contains(email, "@") {
		return "No es un email valido"
	}
	return ""
}

func ValidateAddress(c *fiber.Ctx) string {
	address := strings.TrimSpace(c.FormValue("address"))
	if address == "" {
		return "El campo Dirección no puede estar vacío"
	}
	fmt.Println(address)
	return isAlphanumeric(address, ".")

}

func ValidatePostalCode(c *fiber.Ctx) string {
	postalCode := strings.TrimSpace(c.FormValue("postal-code"))
	if postalCode == "" {
		return "El campo Codigo postal no puede estar vacío"
	}
	if utf8.RuneCountInString(postalCode) > 4 {
		return "El Codigo postal no puede tener mas de 4 numeros"
	}
	return isNumber(postalCode, "")
}

func ValidateDistrict(c *fiber.Ctx) string {
	district := strings.TrimSpace(c.FormValue("district"))
	if district == "" {
		return "El campo Localidad no puede estar vacío"
	}
	return isAlphanumeric(district, ".")
}

func ValidateMemberNumber(c *fiber.Ctx) string {
	memberNumber := strings.TrimSpace(c.FormValue("member-number"))
	if memberNumber == "" {
		return "El campo Numero de afiliado no puede estar vacío"
	}
	return isNumber(memberNumber, "")
}

func ValidateEnterpriseNumber(c *fiber.Ctx) string {
	enterpriseNumber := strings.TrimSpace(c.FormValue("enterprise-number"))
	if enterpriseNumber == "" {
		return "El campo Numero de empresa no puede estar vacío"
	}
	return isNumber(enterpriseNumber, "")
}

func ValidateCUIL(c *fiber.Ctx) string {
	cuil := strings.TrimSpace(c.FormValue("cuil"))
	cuil = strings.Trim(cuil, "-")
	if cuil == "" {
		return "El campo CUIL no puede estar vacío"
	}
	return isNumber(cuil, "-")
}

func ValidateCUIT(c *fiber.Ctx) string {
	cuit := strings.TrimSpace(c.FormValue("cuit"))
	cuit = strings.Trim(cuit, "-")
	if cuit == "" {
		return "El campo CUIT no puede estar vacío"
	}
	return isNumber(cuit, "-")
}

func ValidateIdEnterprise(c *fiber.Ctx) string {
	idEnterprise := strings.TrimSpace(c.FormValue("id-enterprise"))
	if idEnterprise == "" {
		return "Elegir una empresa"
	}
	return ""
}

func ValidateCategory(c *fiber.Ctx) string {
	category := strings.TrimSpace(c.FormValue("category"))
	if category == "" {
		return "Elegir una categoría"
	}
	return ""
}

func ValidateEntryDate(c *fiber.Ctx) string {
	entryDate := strings.TrimSpace(c.FormValue("entry-date"))
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

func validateDateFormat(date string) string {
	if len(date) != 10 {
		return "Formato de fecha erróneo"
	}
	// verifico que sean numeros
	_, dayErr := strconv.Atoi(date[0:2])
	_, monthErr := strconv.Atoi(date[3:5])
	_, yearErr := strconv.Atoi(date[6:])
	if dayErr != nil || monthErr != nil || yearErr != nil ||
		string(date[2]) != "/" || string(date[5]) != "/" {
		return "Formato de fecha erróneo"
	}
	return ""
}

func validateDateValue(date string) string {
	// verifico que sean fechas validas
	day, _ := strconv.Atoi(date[0:2])
	month, _ := strconv.Atoi(date[3:5])
	year, _ := strconv.Atoi(date[6:])

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

func isLetter(value, allowed string) string {
	// se incluye Ã por la codificacion
	letters := " abcdefghijklmnñopqrstuvwxyzáéíóúüÃ"
	letters += allowed
	value = strings.ToLower(value)
	for i := range value {
		if !strings.Contains(letters, string(value[i])) {
			return "El campo posee un caracter erróneo"
		}
	}
	return ""
}

func isNumber(value, allowed string) string {
	numbers := " 0123456789"
	numbers += allowed
	for i := range value {
		if !strings.Contains(numbers, string(value[i])) {
			return "El campo posee un caracter erróneo"
		}
	}
	return ""
}

func isAlphanumeric(value, allowed string) string {
	// se incluye Ã por la codificacion
	characters := " abcdefghijklmnñopqrstuvwxyzáéíóúüÃ0123456789"
	characters += allowed
	value = strings.ToLower(value)
	for i := range value {
		if !strings.Contains(characters, string(value[i])) {
			return "El campo posee un caracter erróneo2"
		}
	}
	return ""
}

func FormatToYYYYMMDD(date string) string {
	day := date[0:2]
	month := date[3:5]
	year := date[6:]
	date = year + "/" + month + "/" + day
	return date
}

func FormatToDDMMYYYY(date string) string {
	year := date[0:4]
	month := date[5:7]
	day := date[8:]
	date = day + "/" + month + "/" + year
	return date
}
