package models

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
)

func ValidateName(c *fiber.Ctx) (bool, string) {
	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return false, "El campo Nombre no puede estar vacío"
	}
	return isLetter(name, "")
}

func ValidateLastName(c *fiber.Ctx) (bool, string) {
	lastName := strings.TrimSpace(c.FormValue("last-name"))
	if lastName == "" {
		return false, "El campo Apellido no puede estar vacío"
	}
	return isLetter(lastName, "")
}

func ValidateEnterpriseName(c *fiber.Ctx) (bool, string) {
	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return false, "El campo Nombre no puede estar vacío"
	}
	return isAlphanumeric(name, ".")
}

func ValidateDNI(c *fiber.Ctx) (bool, string) {
	dni := strings.TrimSpace(c.FormValue("dni"))
	if dni == "" {
		return false, "El campo DNI no puede estar vacío"
	}
	if utf8.RuneCountInString(dni) > 8 {
		return false, "El DNI no puede tener mas de 8 caracteres"
	}
	return isNumber(dni, "")
}

func ValidateBirthday(c *fiber.Ctx) (bool, string) {
	birthday := strings.TrimSpace(c.FormValue("birthday"))
	if birthday == "" {
		return false, "El campo Fecha de nacimiento no puede estar vacío"
	}

	var valid bool
	var err string
	if valid, err = validateDateFormat(birthday); !valid {
		return false, err
	}
	if valid, err = validateDateValue(birthday); !valid {
		return false, err
	}
	return true, ""
}

func ValidateGender(c *fiber.Ctx) (bool, string) {
	gender := strings.TrimSpace(c.FormValue("gender"))
	if gender == "" {
		return false, "Elegir un género"
	}
	return true, ""
}

func ValidateRel(c *fiber.Ctx) (bool, string) {
	rel := strings.TrimSpace(c.FormValue("rel"))
	if rel == "" {
		return false, "El campo Parentesco no puede estar vacío"
	}
	return isLetter(rel, "")
}

func ValidateMaritalStatus(c *fiber.Ctx) (bool, string) {
	maritalStatus := strings.TrimSpace(c.FormValue("marital-status"))
	if maritalStatus == "" {
		return false, "Elegir un estado civil"
	}
	return true, ""
}

func ValidatePhone(c *fiber.Ctx) (bool, string) {
	phone := strings.TrimSpace(c.FormValue("phone"))
	if phone == "" {
		return false, "El campo Teléfono no puede estar vacío"
	}
	return isNumber(phone, "+")
}

func ValidateEmail(c *fiber.Ctx) (bool, string) {
	// email puede estar vacio
	email := strings.TrimSpace(c.FormValue("email"))
	if email == "" {
		return true, ""
	} else if !strings.Contains(email, "@") {
		return false, "No es un email valido"
	}
	return true, ""
}

func ValidateAddress(c *fiber.Ctx) (bool, string) {
	address := strings.TrimSpace(c.FormValue("address"))
	if address == "" {
		return false, "El campo Dirección no puede estar vacío"
	}
	return isAlphanumeric(address, ".")
}

func ValidatePostalCode(c *fiber.Ctx) (bool, string) {
	postalCode := strings.TrimSpace(c.FormValue("postal-code"))
	if postalCode == "" {
		return false, "El campo Codigo postal no puede estar vacío"
	}
	if utf8.RuneCountInString(postalCode) > 4 {
		return false, "El Codigo postal no puede tener mas de 4 numeros"
	}
	return isNumber(postalCode, "")
}

func ValidateDistrict(c *fiber.Ctx) (bool, string) {
	district := strings.TrimSpace(c.FormValue("district"))
	if district == "" {
		return false, "El campo Localidad no puede estar vacío"
	}
	return isAlphanumeric(district, ".")
}

func ValidateMemberNumber(c *fiber.Ctx) (bool, string) {
	memberNumber := strings.TrimSpace(c.FormValue("member-number"))
	if memberNumber == "" {
		return false, "El campo Numero de afiliado no puede estar vacío"
	}
	return isNumber(memberNumber, "")
}

func ValidateEnterpriseNumber(c *fiber.Ctx) (bool, string) {
	enterpriseNumber := strings.TrimSpace(c.FormValue("enterprise-number"))
	if enterpriseNumber == "" {
		return false, "El campo Numero de empresa no puede estar vacío"
	}
	return isNumber(enterpriseNumber, "")
}

func ValidateCUIL(c *fiber.Ctx) (bool, string) {
	cuil := strings.TrimSpace(c.FormValue("cuil"))
	cuil = strings.Trim(cuil, "-")
	if cuil == "" {
		return false, "El campo CUIL no puede estar vacío"
	}
	return isNumber(cuil, "-")
}

func ValidateCUIT(c *fiber.Ctx) (bool, string) {
	cuit := strings.TrimSpace(c.FormValue("cuit"))
	cuit = strings.Trim(cuit, "-")
	if cuit == "" {
		return false, "El campo CUIT no puede estar vacío"
	}
	return isNumber(cuit, "-")
}

func ValidateIdEnterprise(c *fiber.Ctx) (bool, string) {
	idEnterprise := strings.TrimSpace(c.FormValue("id-enterprise"))
	if idEnterprise == "0" {
		return false, "Elegir una empresa"
	}
	return true, ""
}

func ValidateCategory(c *fiber.Ctx) (bool, string) {
	category := strings.TrimSpace(c.FormValue("category"))
	if category == "" {
		return false, "Elegir una categoría"
	}
	return true, ""
}

func ValidateEntryDate(c *fiber.Ctx) (bool, string) {
	entryDate := strings.TrimSpace(c.FormValue("entry-date"))
	if strings.TrimSpace(entryDate) == "" {
		return false, "El campo Fecha de ingreso no puede estar vacío"
	}
	var valid bool
	var err string
	if valid, err = validateDateFormat(entryDate); !valid {
		return false, err
	}
	if valid, err = validateDateValue(entryDate); !valid {
		return false, err
	}
	return true, ""
}

func validateDateFormat(date string) (bool, string) {
	if len(date) != 10 {
		return false, "Formato de fecha erróneo"
	}
	// verifico que sean numeros
	_, dayErr := strconv.Atoi(date[0:2])
	_, monthErr := strconv.Atoi(date[3:5])
	_, yearErr := strconv.Atoi(date[6:])
	if dayErr != nil || monthErr != nil || yearErr != nil ||
		string(date[2]) != "/" || string(date[5]) != "/" {
		return false, "Formato de fecha erróneo"
	}
	return true, ""
}

func validateDateValue(date string) (bool, string) {
	// verifico que sean fechas validas
	day, _ := strconv.Atoi(date[0:2])
	month, _ := strconv.Atoi(date[3:5])
	year, _ := strconv.Atoi(date[6:])

	if month < 1 || month > 12 {
		return false, "Fecha errónea"
	}
	if day < 1 || day > 31 {
		return false, "Fecha errónea"
	}
	switch month {
	case 2:
		if day > 29 {
			return false, "Fecha errónea"
		}
	case 4, 6, 9, 11:
		if day > 30 {
			return false, "Fecha errónea"
		}
	}
	if year < 1900 || year > int(time.Now().Year()) {
		return false, "Fecha errónea"
	}
	return true, ""
}

func validateMonth(month string) bool {
	months := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	if !slices.Contains(months, month) {
		return false
	}
	return true
}

func validateYear(year string) bool {
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println(err)
	}

	if yearInt < 1990 || yearInt > time.Now().Year() {
		return false
	}
	return true
}

func validatePayment(c *fiber.Ctx) (bool, string) {
	month := c.FormValue("month")
	year := c.FormValue("year")
	if !validateMonth(month) {
		return false, "El mes no es correcto"
	}
	if !validateYear(year) {
		return false, "El año no es correcto"
	}
	payment := fmt.Sprintf("01/%s/%s", month, year)
	return validateDateValue(payment)
}

func validateStatus(c *fiber.Ctx) (bool, string) {
	status := c.FormValue("status")
	if status == "PAGO" || status == "IMPAGO" {
		return true, ""
	}
	return false, "Elegir estado del pago"
}

func validatePaymentAmount(c *fiber.Ctx) (bool, string) {
	amount := c.FormValue("amount")
	if amount == "" {
		return false, "El campo Monto no puede estar vacío"
	}
	return isNumber(amount, "")
}

func validatePaymentDate(c *fiber.Ctx) (bool, string) {

	// esto se puede escalar, esta repetido con las otras fechas
	paymentDate := strings.TrimSpace(c.FormValue("payment-date"))
	if paymentDate == "" {
		return false, "El campo Fecha de pago no puede estar vacío"
	}

	var valid bool
	var err string
	if valid, err = validateDateFormat(paymentDate); !valid {
		return false, err
	}
	if valid, err = validateDateValue(paymentDate); !valid {
		return false, err
	}
	return true, ""
}

func validateCommentary(c *fiber.Ctx) (bool, string) {
	if len(c.FormValue("commentary")) > 400 {
		return false, "El máximo de caracteres es 400"
	}
	return true, ""
}

func isLetter(value, allowed string) (bool, string) {
	// se incluye Ã por la codificacion
	letters := " abcdefghijklmnñopqrstuvwxyzáéíóúüÃ"
	letters += allowed
	value = strings.ToLower(value)
	for i := range value {
		if !strings.Contains(letters, string(value[i])) {
			return false, "El campo posee un caracter erróneo"
		}
	}
	return true, ""
}

func isNumber(value, allowed string) (bool, string) {
	numbers := " 0123456789"
	numbers += allowed
	for i := range value {
		if !strings.Contains(numbers, string(value[i])) {
			return false, "El campo posee un caracter erróneo"
		}
	}
	return true, ""
}

func isAlphanumeric(value, allowedCharacter string) (bool, string) {
	// se incluye Ã por la codificacion
	characters := " abcdefghijklmnñopqrstuvwxyzáéíóúüÃ0123456789"
	characters += allowedCharacter
	value = strings.ToLower(value)
	for i := range value {
		if !strings.Contains(characters, string(value[i])) {
			return false, "El campo posee un caracter erróneo"
		}
	}
	return true, ""
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
