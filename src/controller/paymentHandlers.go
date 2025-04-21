package controller

import (
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RenderAddPaymentForm(c *fiber.Ctx) error {
	// se manda un payment vacio para que esten todos los input en blanco
	IdEnterprise, err := getIdModelCaller(models.Enterprise{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	p := models.Payment{IdEnterprise: IdEnterprise}
	data := fiber.Map{"payment": p, "mode": "add"}
	return c.Render("paymentFile", data)
}

func RenderPaymentFile(c *fiber.Ctx) error {
	p, err := searchOneModelByIdCaller(models.Payment{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	IdEnterprise, err := getIdModelCaller(models.Enterprise{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}

	result2, err := database.DB.Query("SELECT Year FROM PaymentTable WHERE IdEnterprise = ? GROUP BY Year ORDER BY YEAR DESC", IdEnterprise)
	if err != nil {
		// guardar el error
		er.QueryError.Msg = err.Error()
		return er.CheckError(c, er.QueryError)
	}

	var years []string
	var year string
	for result2.Next() {
		result2.Scan(&year)
		years = append(years, year)
	}

	createdAt, updatedAt, err := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	data := fiber.Map{"payment": p, "role": role, "mode": "edit", "idEnterprise": IdEnterprise, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("paymentFile", data)
}

func AddPayment(c *fiber.Ctx) error {
	if err := validateFieldsCaller(models.Payment{}, c); err != nil {
		return er.CheckError(c, err)
	}
	p, err := parserCaller(i.PaymentParser{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}

	IdEnterprise, err := getIdModelCaller(models.Enterprise{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}

	p, err = insertModelCaller(p)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	years, err := getPaymentYearsFromDB(IdEnterprise)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	createdAt, updatedAt, err := getPaymentTimestampsFromDB(p.IdPayment)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}

	role := c.Locals("claims").(jwt.MapClaims)["role"]
	data := fiber.Map{"payment": p, "mode": "edit", "role": role, "idEnterprise": IdEnterprise, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("paymentFile", data)
}

func EditPayment(c *fiber.Ctx) error {
	if err := validateFieldsCaller(models.Payment{}, c); err != nil {
		return er.CheckError(c, err)
	}
	p, err := parserCaller(i.PaymentParser{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	IdEnterprise, err := getIdModelCaller(models.Enterprise{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	IdPayment, err := getIdModelCaller(models.Payment{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	p.IdPayment = IdPayment

	p, err = updateModelCaller(p)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	years, err := getPaymentYearsFromDB(IdEnterprise)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	createdAt, updatedAt, err := getPaymentTimestampsFromDB(p.IdPayment)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}

	data := fiber.Map{"payment": p, "mode": "edit", "idEnterprise": IdEnterprise, "role": role, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("paymentFile", data)

}

func DeletePayment(c *fiber.Ctx) error {
	IdPayment, err := getIdModelCaller(models.Payment{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	p := models.Payment{IdPayment: IdPayment}
	deleteModelCaller(p)
	return RenderEnterprisePaymentsTable(c)
}

func getPaymentYearsFromDB(idEnterprise int) ([]string, error) {
	result, err := database.DB.Query("SELECT Year FROM PaymentTable WHERE IdEnterprise = '?' GROUP BY Year ORDER BY YEAR DESC", idEnterprise)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, er.QueryError
	}

	var years []string
	var year string
	for result.Next() {
		err = result.Scan(&year)
		if err != nil {
			er.ScanError.Msg = err.Error()
			return nil, er.ScanError
		}
		years = append(years, year)
	}
	return years, nil
}

func getPaymentTimestampsFromDB(idPayment int) (string, string, error) {
	result, err := database.DB.Query("SELECT CreatedAt, UpdatedAt FROM PaymentTable WHERE IdPayment = ?", idPayment)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return "", "", er.QueryError
	}
	var createdAtUnformatted, updatedAtUnformatted time.Time
	for result.Next() {
		err = result.Scan(&createdAtUnformatted, &updatedAtUnformatted)
		if err != nil {
			er.ScanError.Msg = err.Error()
			return "", "", er.ScanError
		}
	}
	return formatTimeStamps(createdAtUnformatted, updatedAtUnformatted)
}
