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
	enterpriseName, err := getEnterpriseName(IdEnterprise)
	if err!=nil{
		return er.CheckError(c, err)
	}
	p := models.Payment{IdEnterprise: IdEnterprise}
	data := fiber.Map{"payment": p, "mode": "add", "enterpriseName": enterpriseName}
	return c.Render("paymentFile", data)
}

func RenderPaymentFile(c *fiber.Ctx) error {
	p, err := searchOneModelByIdCaller(models.Payment{}, c)

	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
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
	enterpriseName, err := getEnterpriseName(IdEnterprise)
	if err!=nil{
		return er.CheckError(c, err)
	}
	data := fiber.Map{"payment": p, "mode": "edit", "idEnterprise": IdEnterprise, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt, "enterpriseName": enterpriseName}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deletePayment"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writePayment"]
	return c.Render("paymentFile", data)
}

func RenderPaymentTable(c *fiber.Ctx) error {
	params := struct {
		Year int `params:"year"`
	}{}
	c.ParamsParser(&params)
	canDelete := c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
	canWrite := c.Locals("claims").(jwt.MapClaims)["writeEnterprise"]
	IdEnterprise, err := getIdModelCaller(models.Enterprise{}, c)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	enterpriseName, err := getEnterpriseName(IdEnterprise)
	if err!=nil{
		return er.CheckError(c, err)
	}

	var totalRows int
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM PaymentTable 
		WHERE IdEnterprise = ?`, IdEnterprise)
	// row.Scan copia el numero de fila en la variable count
	err = row.Scan(&totalRows)
	if err != nil {
		er.ScanError.Msg = err.Error()
		return er.CheckError(c, er.ScanError)
	}

	if totalRows == 0 {
		data := fiber.Map{"idEnterprise": IdEnterprise, "canDelete": canDelete, "canWrite": canWrite, "mode": "edit", "empty": true, "enterpriseName": enterpriseName}
		return c.Render("paymentTable", data)
	}
	var lastYear int
	result, err := database.DB.Query("SELECT MAX(Year) FROM PaymentTable WHERE IdEnterprise = ?", IdEnterprise)
	if err != nil {
		// logear el err
		er.QueryError.Msg = err.Error()
		return er.CheckError(c, er.QueryError)
	}
	for result.Next() {
		err = result.Scan(&lastYear)
	}
	if err != nil {
		// logear el err
		er.ScanError.Msg = err.Error()
		return er.CheckError(c, er.ScanError)
	}

	result2, err := database.DB.Query("SELECT Year FROM PaymentTable WHERE IdEnterprise = ? GROUP BY Year ORDER BY YEAR DESC", IdEnterprise)
	if err != nil {
		// logear el err
		er.QueryError.Msg = err.Error()
		return er.CheckError(c, er.QueryError)
	}

	var years []string
	var year string
	for result2.Next() {
		err = result2.Scan(&year)
		if err != nil {
			// logear el err
			er.ScanError.Msg = err.Error()
			return er.CheckError(c, er.ScanError)
		}
		years = append(years, year)
	}
	yearInt := params.Year
	if params.Year == 0 {
		payments, _, err := searchModelsCaller(models.Payment{}, c, lastYear)
		if err != nil {
			// guardar el err
			return er.CheckError(c, err)
		}
		data := fiber.Map{"payments": payments, "idEnterprise": IdEnterprise, "canDelete": canDelete, "canWrite": canWrite, "mode": "edit", "years": years, "year": lastYear, "enterpriseName": enterpriseName}
		return c.Render("paymentTable", data)
	} else {
		payments, _, err := searchModelsCaller(models.Payment{}, c, params.Year)
		if err != nil {
			// guardar el err
			return er.CheckError(c, err)
		}
	
	data := fiber.Map{"payments": payments, "idEnterprise": IdEnterprise, "canDelete": canDelete, "canWrite": canWrite, "mode": "edit", "years": years, "year": yearInt, "enterpriseName": enterpriseName}
	return c.Render("paymentTable", data)
	}
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
	enterpriseName, err := getEnterpriseName(IdEnterprise)
	if err!=nil{
		return er.CheckError(c, err)
	}
	data := fiber.Map{"payment": p, "mode": "edit", "idEnterprise": IdEnterprise, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt, "enterpriseName": enterpriseName}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deletePayment"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writePayment"]
	return c.Status(fiber.StatusCreated).Render("paymentFile", data)
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
	enterpriseName, err := getEnterpriseName(IdEnterprise)
	if err!=nil{
		return er.CheckError(c, err)
	}

	data := fiber.Map{"payment": p, "mode": "edit", "idEnterprise": IdEnterprise, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt, "enterpriseName": enterpriseName}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deletePayment"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writePayment"]
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
	return RenderPaymentTable(c)
}

func getPaymentYearsFromDB(idEnterprise int) ([]string, error) {
	result, err := database.DB.Query("SELECT Year FROM PaymentTable WHERE IdEnterprise = ? GROUP BY Year ORDER BY YEAR DESC", idEnterprise)
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
