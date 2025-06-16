package models

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
	"github.com/gofiber/fiber/v2"
)

func ValidateName(c *fiber.Ctx) error {
	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		er.ValidationError.Msg = "field 'name' can't be empty"
		return er.ValidationError
	}
	return isLetter(name, "name", "")
}

func ValidateLastName(c *fiber.Ctx) error {
	lastName := strings.TrimSpace(c.FormValue("last-name"))
	if lastName == "" {
		er.ValidationError.Msg = "field 'last-name' can't be empty"
		return er.ValidationError
	}
	return isLetter(lastName, "last-name", "")
}

func ValidateEnterpriseName(c *fiber.Ctx) error {
	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		er.ValidationError.Msg = "field 'name' can't be empty"
		return er.ValidationError
	}
	return isAlphanumeric(name, "name", ".")
}

func ValidateDNI(c *fiber.Ctx) error {
	dni := strings.TrimSpace(c.FormValue("dni"))
	if dni == "" {
		er.ValidationError.Msg = "the field 'DNI' has a invalid character"
		return er.ValidationError
	}
	if utf8.RuneCountInString(dni) > 8 {
		er.ValidationError.Msg = "field 'DNI' can't have more than 8 characters"
		return er.ValidationError
	}
	return isNumber(dni, "DNI", "")
}

func ValidateBirthday(c *fiber.Ctx) error {
	birthday := strings.TrimSpace(c.FormValue("birthday"))
	if birthday == "" {
		er.ValidationError.Msg = "field 'birthday' can't be empty"
		return er.ValidationError
	}

	return validateDate(birthday, "birthday")
}

func ValidateGender(c *fiber.Ctx) error {
	gender := strings.TrimSpace(c.FormValue("gender"))
	genders := []string{"Masculino", "Femenino", "Otro"}
	if !slices.Contains(genders, gender) {
		er.ValidationError.Msg = "field 'gender' is invalid"
		return er.ValidationError
	}
	return nil
}

func ValidateRel(c *fiber.Ctx) error {
	rel := strings.TrimSpace(c.FormValue("rel"))
	if rel == "" {
		er.ValidationError.Msg = "field 'rel' can't be empty"
		return er.ValidationError
	}
	return isLetter(rel, "rel", "")
}

func ValidateMaritalStatus(c *fiber.Ctx) error {
	maritalStatus := strings.TrimSpace(c.FormValue("marital-status"))
	mmSS := []string{"Soltero", "Casado", "Separado", "Divorciado", "Viudo"}
	if !slices.Contains(mmSS, maritalStatus) {
		er.ValidationError.Msg = "field 'marital-status' is invalid"
		return er.ValidationError
	}
	return nil
}

func ValidatePhone(c *fiber.Ctx) error {
	phone := strings.TrimSpace(c.FormValue("phone"))
	if phone == "" {
		er.ValidationError.Msg = "field 'phone' can't be empty"
		return er.ValidationError
	}
	return isNumber(phone, "phone", "+")
}

func ValidateEmail(c *fiber.Ctx) error {
	// email puede estar vacio
	email := strings.TrimSpace(c.FormValue("email"))
	if email == "" {
		return nil
	} else if !strings.Contains(email, "@") {
		er.ValidationError.Msg = "the field 'email' is invalid"
		return er.ValidationError
	}
	return nil
}

func ValidateContact(c *fiber.Ctx) error {
	if len(c.FormValue("contact")) > 200 {
		er.ValidationError.Msg = "field 'contact' can't have more than 200 characters"
		return er.ValidationError
	}
	return nil
}

func ValidateAddress(c *fiber.Ctx) error {
	address := strings.TrimSpace(c.FormValue("address"))
	if address == "" {
		er.ValidationError.Msg = "field 'address' can't be empty"
		return er.ValidationError
	}
	return isAlphanumeric(address, "address", ".")
}

func ValidatePostalCode(c *fiber.Ctx) error {
	postalCode := strings.TrimSpace(c.FormValue("postal-code"))
	if postalCode == "" {
		er.ValidationError.Msg = "field 'postal-code' can't be empty"
		return er.ValidationError
	}
	if utf8.RuneCountInString(postalCode) > 4 {
		er.ValidationError.Msg = "field 'postal-code' can't have more than 4 characters"
		return er.ValidationError
	}
	return isNumber(postalCode, "postal-code", "")
}

func ValidateDistrict(c *fiber.Ctx) error {
	district := strings.TrimSpace(c.FormValue("district"))
	if district == "" {
		er.ValidationError.Msg = "field 'district' can't be empty"
		return er.ValidationError
	}
	return isAlphanumeric(district, "district", ".")
}

func ValidateMemberNumber(c *fiber.Ctx) error {
	memberNumber := strings.TrimSpace(c.FormValue("member-number"))
	if memberNumber == "" {
		er.ValidationError.Msg = "field 'member-number' can't be empty"
		return er.ValidationError
	}
	return isNumber(memberNumber, "member-number", "")
}

func ValidateEnterpriseNumber(c *fiber.Ctx) error {
	enterpriseNumber := strings.TrimSpace(c.FormValue("enterprise-number"))
	oldEnterpriseNumber := strings.TrimSpace(c.FormValue("old-enterprise-number"))
	if enterpriseNumber == "" {
		er.ValidationError.Msg = "field 'enterprise-number' can't be empty"
		return er.ValidationError
	}
	if oldEnterpriseNumber == enterpriseNumber {
		return nil
	}
	enterprisesNumbers, err := GetAllEnterprisesNumbersFromDB()
	if err != nil {
		return err
	}
	if slices.Contains(enterprisesNumbers, enterpriseNumber) {
		er.ValidationError.Msg = "enterprise number already exists"
		return er.ValidationError
	}

	return isNumber(enterpriseNumber, "enterprise-number", "")
}

func ValidateAffiliated(c *fiber.Ctx) error {
	affiliated, err := strconv.ParseBool(c.FormValue("affiliated"))
	if err != nil {
		er.InternalServerError.Msg = err.Error()
		return er.CheckError(c, er.InternalServerError)
	}
	if affiliated || !affiliated {
		return nil
	} else {
		er.ValidationError.Msg = "field affiliated not contains true or false"
		return er.ValidationError
	}
}

func ValidateCUIL(c *fiber.Ctx) error {
	cuil := strings.TrimSpace(c.FormValue("cuil"))
	cuil = strings.Trim(cuil, "-")
	if cuil == "" {
		er.ValidationError.Msg = "field 'CUIL' can't be empty"
		return er.ValidationError
	}
	return isNumber(cuil, "cuil", "-")
}

func ValidateCUIT(c *fiber.Ctx) error {
	cuit := strings.TrimSpace(c.FormValue("cuit"))
	cuit = strings.Trim(cuit, "-")
	if cuit == "" {
		er.ValidationError.Msg = "field 'cuit' can't be empty"
		return er.ValidationError
	}
	return isNumber(cuit, "cuit", "-")
}

func ValidateIdEnterprise(c *fiber.Ctx) error {
	idEnterpriseStr := c.FormValue("id-enterprise")
	idEnterprise, err := strconv.Atoi(idEnterpriseStr)
	if err != nil {
		er.ValidationError.Msg = "enterprise's ID is not valid"
		return er.ValidationError
	}
	enterprisesId, err := GetAllEnterprisesIdFromDB()
	if err != nil {
		er.ValidationError.Msg = err.Error()
		return er.ValidationError
	}
	for _, id := range enterprisesId {
		if id == idEnterprise {
			return nil
		}
	}
	er.ValidationError.Msg = "enterprise's ID is not valid"
	return er.ValidationError
}

func GetAllEnterprisesIdFromDB() ([]int, error) {
	var enterprisesId []int
	var idEnterprise int
	result, err := database.DB.Query("SELECT IdEnterprise FROM EnterpriseTable")
	if err != nil {
		er.QueryError.Msg = "internal error"
		return nil, er.QueryError
	}
	for result.Next() {
		err = result.Scan(&idEnterprise)
		if err != nil {
			er.ScanError.Msg = "internal error"
			return nil, er.ScanError
		}
		enterprisesId = append(enterprisesId, idEnterprise)
	}
	return enterprisesId, nil
}

func GetAllEnterprisesNumbersFromDB() ([]string, error) {
	var ee []string
	var e string
	result, err := database.DB.Query("SELECT EnterpriseNumber FROM EnterpriseTable")
	if err != nil {
		er.QueryError.Msg = "internal error"
		return nil, er.QueryError
	}
	for result.Next() {
		err = result.Scan(&e)
		if err != nil {
			er.ScanError.Msg = "internal error"
			return nil, er.ScanError
		}
		ee = append(ee, e)
	}
	return ee, nil
}

func ValidateCategory(c *fiber.Ctx) error {
	category := c.FormValue("category")
	categories := []string{"Nivel 1: Oficial Múltiple", "Nivel 2: Oficial Especializado", "Nivel 3: Oficial General", "Nivel 4: Medio Oficial", "Nivel 5: Ayudante", "Nivel 6: Operario Act. Industrial"}
	if !slices.Contains(categories, category) {
		er.ValidationError.Msg = "field 'category' is invalid"
		return er.ValidationError
	}
	return nil
}

func ValidateEntryDate(c *fiber.Ctx) error {
	entryDate := strings.TrimSpace(c.FormValue("entry-date"))
	if strings.TrimSpace(entryDate) == "" {
		er.ValidationError.Msg = "field 'entry-date' can't be empty"
		return er.ValidationError
	}
	return validateDate(entryDate, "entry-date")
}

func validateDate(date, field string) error {
	if date == "" {
		er.ValidationError.Msg = fmt.Sprintf("field '%s' can't be empty", field)
		return er.ValidationError
	}
	if len(date) != 10 {
		er.ValidationError.Msg = fmt.Sprintf("field '%s' has an invalid format date", field)
		return er.ValidationError
	}
	// verifico que sean numeros
	day, dayErr := strconv.Atoi(date[0:2])
	month, monthErr := strconv.Atoi(date[3:5])
	year, yearErr := strconv.Atoi(date[6:])

	// si alguno no es numero o si no estan bien colocadas las "/"
	if dayErr != nil || monthErr != nil || yearErr != nil || string(date[2]) != "/" || string(date[5]) != "/" {
		er.ValidationError.Msg = fmt.Sprintf("field '%s' has an invalid format date", field)
		return er.ValidationError
	}

	// verifico que sea una fecha valida
	if month < 1 || month > 12 {
		er.ValidationError.Msg = fmt.Sprintf("field '%s' has an invalid format date", field)
		return er.ValidationError
	}
	if day < 1 || day > 31 {
		er.ValidationError.Msg = fmt.Sprintf("field '%s' has an invalid format date", field)
		return er.ValidationError
	}
	switch month {
	case 2:
		if day > 29 {
			er.ValidationError.Msg = fmt.Sprintf("field '%s' has an invalid format date", field)
			return er.ValidationError
		}
	case 4, 6, 9, 11:
		if day > 30 {
			er.ValidationError.Msg = fmt.Sprintf("field '%s' has an invalid format date", field)
			return er.ValidationError
		}
	}
	if year < 1900 || year > int(time.Now().Year()) {
		er.ValidationError.Msg = fmt.Sprintf("field '%s' has an invalid format date", field)
		return er.ValidationError
	}
	return nil
}

func ValidatePayment(c *fiber.Ctx) error {
	month := c.FormValue("month")
	year := c.FormValue("year")

	payment := fmt.Sprintf("01/%s/%s", month, year)
	return validateDate(payment, "payment")
}

func ValidateStatus(c *fiber.Ctx) error {
	status, err := strconv.ParseBool(c.FormValue("status"))
	if err != nil {
		er.InternalServerError.Msg = err.Error()
		return er.CheckError(c, er.InternalServerError)
	}
	if status == true || status == false {
		return nil
	} else {
		er.ValidationError.Msg = "field status not contains any boolean value"
		return er.ValidationError
	}
}

func ValidatePaymentAmount(c *fiber.Ctx) error {
	amount := c.FormValue("amount")
	if amount == "" {
		er.ValidationError.Msg = "field 'amount' can't be empty"
		return er.ValidationError
	}
	return isNumber(amount, "amount", "")
}

func ValidatePaymentDate(c *fiber.Ctx) error {
	paymentDate := strings.TrimSpace(c.FormValue("payment-date"))
	// la fecha de pago puede estar vacia
	if paymentDate == "" {
		return nil
	}
	return validateDate(paymentDate, "payment-date")

}

func ValidateObservations(c *fiber.Ctx) error {
	if len(c.FormValue("observations")) > 1000 {
		er.ValidationError.Msg = "field 'observations' can't have more than 1000 characters"
		return er.ValidationError
	}
	return nil
}

func isLetter(value, field, allowed string) error {
	// se incluye Ã por la codificacion
	letters := " abcdefghijklmnñopqrstuvwxyzáéíóúüÃ"
	letters += allowed
	value = strings.ToLower(value)
	for i := range value {
		if !strings.Contains(letters, string(value[i])) {
			er.ValidationError.Msg = fmt.Sprintf("the field '%s' has a invalid character", field)
			return er.ValidationError
		}
	}
	return nil
}

func isNumber(value, field, allowed string) error {
	numbers := " 0123456789"
	numbers += allowed
	for i := range value {
		if !strings.Contains(numbers, string(value[i])) {
			er.ValidationError.Msg = fmt.Sprintf("the field '%s' has a invalid character", field)
			return er.ValidationError
		}
	}
	return nil
}

func isAlphanumeric(value, field, allowedCharacter string) error {
	// se incluye Ã por la codificacion
	characters := " abcdefghijklmnñopqrstuvwxyzáéíóúüÃ0123456789"
	characters += allowedCharacter
	value = strings.ToLower(value)
	for i := range value {
		if !strings.Contains(characters, string(value[i])) {
			er.ValidationError.Msg = fmt.Sprintf("the field '%s' has a invalid character", field)
			return er.ValidationError
		}
	}
	return nil
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
