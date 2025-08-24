package controller

import (
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/LucasBastino/app-sindicato/src/errors/errorHandler"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RenderAddPaymentForm(c *fiber.Ctx) error {
	// se manda un payment vacio para que esten todos los input en blanco
	IdEnterprise, customErr := getIdModelCaller(models.Enterprise{}, c)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	enterpriseName, customErr := getEnterpriseName(IdEnterprise)
	if (customErr != customError.CustomError{}){
		return errorHandler.HandleError(c, &customErr)
	}
	p := models.Payment{IdEnterprise: IdEnterprise}
	data := fiber.Map{"payment": p, "mode": "add", "enterpriseName": enterpriseName}
	return c.Render("paymentFile", data)
}

func RenderPaymentFile(c *fiber.Ctx) error {
	p, customErr := searchOneModelByIdCaller(models.Payment{}, c)

	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	IdEnterprise, customErr := getIdModelCaller(models.Enterprise{}, c)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}

	result2, err := database.DB.Query("SELECT Year FROM PaymentTable WHERE IdEnterprise = ? GROUP BY Year ORDER BY YEAR DESC", IdEnterprise)
	if err != nil{
		// guardar el error
		customError.QueryError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.QueryError)
	}

	var years []string
	var year string
	for result2.Next() {
		result2.Scan(&year)
		years = append(years, year)
	}

	createdAt, updatedAt, customErr := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	enterpriseName, customErr := getEnterpriseName(IdEnterprise)
	if (customErr != customError.CustomError{}){
		return errorHandler.HandleError(c, &customErr)
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
	IdEnterprise, customErr := getIdModelCaller(models.Enterprise{}, c)
	if (customErr != customError.CustomError{}){
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	enterpriseName, customErr := getEnterpriseName(IdEnterprise)
	if (customErr != customError.CustomError{}){
		return errorHandler.HandleError(c, &customErr)
	}

	var totalRows int
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM PaymentTable 
		WHERE IdEnterprise = ?`, IdEnterprise)
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		customError.ScanError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.ScanError)
	}

	if totalRows == 0 {
		data := fiber.Map{"idEnterprise": IdEnterprise, "canDelete": canDelete, "canWrite": canWrite, "mode": "edit", "empty": true, "enterpriseName": enterpriseName}
		return c.Render("paymentTable", data)
	}
	var lastYear int
	result, err := database.DB.Query("SELECT MAX(Year) FROM PaymentTable WHERE IdEnterprise = ?", IdEnterprise)
	if err != nil {
		// logear el err
		customError.QueryError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.QueryError)
	}
	for result.Next() {
		err = result.Scan(&lastYear)
	}
	if err != nil {
		// logear el err
		customError.ScanError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.ScanError)
	}

	result2, err := database.DB.Query("SELECT Year FROM PaymentTable WHERE IdEnterprise = ? GROUP BY Year ORDER BY YEAR DESC", IdEnterprise)
	if err != nil {
		// logear el err
		customError.QueryError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.QueryError)
	}

	var years []string
	var year string
	for result2.Next() {
		err = result2.Scan(&year)
		if err != nil {
			// logear el err
			customError.ScanError.Msg = err.Error()
			return errorHandler.HandleError(c, &customError.ScanError)
		}
		years = append(years, year)
	}
	yearInt := params.Year
	if params.Year == 0 {
		payments, _, customErr := searchModelsCaller(models.Payment{}, c, lastYear)
		if (customErr != customError.CustomError{}){
			// guardar el err
			return errorHandler.HandleError(c, &customErr)
		}
		data := fiber.Map{"payments": payments, "idEnterprise": IdEnterprise, "canDelete": canDelete, "canWrite": canWrite, "mode": "edit", "years": years, "year": lastYear, "enterpriseName": enterpriseName}
		return c.Render("paymentTable", data)
	} else {
		payments, _, customErr := searchModelsCaller(models.Payment{}, c, params.Year)
		if (customErr != customError.CustomError{}){
			// guardar el err
			return errorHandler.HandleError(c, &customErr)
		}
	
	data := fiber.Map{"payments": payments, "idEnterprise": IdEnterprise, "canDelete": canDelete, "canWrite": canWrite, "mode": "edit", "years": years, "year": yearInt, "enterpriseName": enterpriseName}
	return c.Render("paymentTable", data)
	}
}

func AddPayment(c *fiber.Ctx) error {
	if customErr := validateFieldsCaller(models.Payment{}, c); (customErr != customError.CustomError{}) {
		return errorHandler.HandleError(c, &customErr)
	}
	p, customErr := parserCaller(i.PaymentParser{}, c)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}

	IdEnterprise, customErr := getIdModelCaller(models.Enterprise{}, c)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}

	p, customErr = insertModelCaller(p)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	years, customErr := getPaymentYearsFromDB(IdEnterprise)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	createdAt, updatedAt, customErr := getPaymentTimestampsFromDB(p.IdPayment)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	enterpriseName, customErr := getEnterpriseName(IdEnterprise)
	if (customErr != customError.CustomError{}){
		return errorHandler.HandleError(c, &customErr)
	}
	data := fiber.Map{"payment": p, "mode": "edit", "idEnterprise": IdEnterprise, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt, "enterpriseName": enterpriseName}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deletePayment"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writePayment"]
	return c.Status(fiber.StatusCreated).Render("paymentFile", data)
}

func EditPayment(c *fiber.Ctx) error {
	if customErr := validateFieldsCaller(models.Payment{}, c); (customErr != customError.CustomError{}){
		return errorHandler.HandleError(c, &customErr)
	}
	p, customErr := parserCaller(i.PaymentParser{}, c)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	IdEnterprise, customErr := getIdModelCaller(models.Enterprise{}, c)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	IdPayment, customErr := getIdModelCaller(models.Payment{}, c)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	p.IdPayment = IdPayment

	p, customErr = updateModelCaller(p)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	years, customErr := getPaymentYearsFromDB(IdEnterprise)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	createdAt, updatedAt, customErr := getPaymentTimestampsFromDB(p.IdPayment)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	enterpriseName, customErr := getEnterpriseName(IdEnterprise)
	if (customErr != customError.CustomError{}){
		return errorHandler.HandleError(c, &customErr)
	}

	data := fiber.Map{"payment": p, "mode": "edit", "idEnterprise": IdEnterprise, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt, "enterpriseName": enterpriseName}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deletePayment"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writePayment"]
	return c.Render("paymentFile", data)

}

func DeletePayment(c *fiber.Ctx) error {
	IdPayment, customErr := getIdModelCaller(models.Payment{}, c)
	if (customErr != customError.CustomError{}){
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	p := models.Payment{IdPayment: IdPayment}
	deleteModelCaller(p)
	return RenderPaymentTable(c)
}

func getPaymentYearsFromDB(idEnterprise int) ([]string, customError.CustomError) {
	result, err := database.DB.Query("SELECT Year FROM PaymentTable WHERE IdEnterprise = ? GROUP BY Year ORDER BY YEAR DESC", idEnterprise)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return nil, customError.QueryError
	}

	var years []string
	var year string
	for result.Next() {
		err = result.Scan(&year)
		if err != nil {
			customError.ScanError.Msg = err.Error()
			return nil, customError.ScanError
		}
		years = append(years, year)
	}
	return years, customError.CustomError{}
}

func getPaymentTimestampsFromDB(idPayment int) (string, string, customError.CustomError) {
	result, err := database.DB.Query("SELECT CreatedAt, UpdatedAt FROM PaymentTable WHERE IdPayment = ?", idPayment)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return "", "", customError.QueryError
	}
	var createdAtUnformatted, updatedAtUnformatted time.Time
	for result.Next() {
		err = result.Scan(&createdAtUnformatted, &updatedAtUnformatted)
		if err != nil {
			customError.ScanError.Msg = err.Error()
			return "", "", customError.ScanError
		}
	}
	return formatTimeStamps(createdAtUnformatted, updatedAtUnformatted)
}
