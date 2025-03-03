package controller

import (
	"fmt"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RenderAddPaymentForm(c *fiber.Ctx) error {
	// se manda un payment vacio para que esten todos los input en blanco
	IdEnterprise := getIdModelCaller(models.Enterprise{}, c)
	p := models.Payment{IdEnterprise: IdEnterprise}
	data := fiber.Map{"payment": p, "mode": "add"}
	return c.Render("paymentFile", data)
}

func RenderPaymentFile(c *fiber.Ctx) error {
	// IdPayment := getIdModelCaller(models.Payment{},c)
	p := searchOneModelByIdCaller(models.Payment{}, c)
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	IdEnterprise := getIdModelCaller(models.Enterprise{}, c)

	result2, err := database.DB.Query(fmt.Sprintf("SELECT Year FROM PaymentTable WHERE IdEnterprise = '%d' GROUP BY Year ORDER BY YEAR DESC", IdEnterprise))
	if err != nil {
		fmt.Println("error searching different Years in PaymentTable")
		panic(err)
	}

	var years []string
	var year string
	for result2.Next() {
		result2.Scan(&year)
		years = append(years, year)
	}

	createdAt, updatedAt := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
	data := fiber.Map{"payment": p, "role": role, "mode": "edit", "idEnterprise": IdEnterprise, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("paymentFile", data)
}

func AddPayment(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Payment{}, c)
	p := parserCaller(i.PaymentParser{}, c)

	if len(errorMap) > 0 {
		fmt.Println(errorMap)
		data := fiber.Map{"payment": p, "errorMap": errorMap, "mode": "add"}
		return c.Render("paymentFile", data)
	}

	IdEnterprise := getIdModelCaller(models.Enterprise{}, c)

	p = insertModelCaller(p)
	years := getPaymentYearsFromDB(IdEnterprise)
	createdAt, updatedAt := getPaymentTimestampsFromDB(p.IdPayment)

	role := c.Locals("claims").(jwt.MapClaims)["role"]
	data := fiber.Map{"payment": p, "mode": "edit", "role": role, "idEnterprise": IdEnterprise, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("paymentFile", data)
}

func EditPayment(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Payment{}, c)
	p := parserCaller(i.PaymentParser{}, c)
	IdEnterprise := getIdModelCaller(models.Enterprise{}, c)
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	IdPayment := getIdModelCaller(models.Payment{}, c)
	p.IdPayment = IdPayment
	if len(errorMap) > 0 {
		data := fiber.Map{"payment": p, "mode": "edit", "role": role, "errorMap": errorMap}
		c.Render("paymentFile", data)
	}
	p = updateModelCaller(p)
	years := getPaymentYearsFromDB(IdEnterprise)
	createdAt, updatedAt := getPaymentTimestampsFromDB(p.IdPayment)

	data := fiber.Map{"payment": p, "mode": "edit", "idEnterprise": IdEnterprise, "role": role, "years": years, "year": p.Year, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("paymentFile", data)

}

func DeletePayment(c *fiber.Ctx) error {
	IdPayment := getIdModelCaller(models.Payment{}, c)
	p := models.Payment{IdPayment: IdPayment}
	deleteModelCaller(p)
	return RenderEnterprisePaymentsTable(c)
}

func getPaymentYearsFromDB(idEnterprise int) []string {
	result, err := database.DB.Query(fmt.Sprintf("SELECT Year FROM PaymentTable WHERE IdEnterprise = '%d' GROUP BY Year ORDER BY YEAR DESC", idEnterprise))
	if err != nil {
		fmt.Println("error searching different Years in PaymentTable")
		panic(err)
	}

	var years []string
	var year string
	for result.Next() {
		err = result.Scan(&year)
		if err != nil {
			fmt.Println("error scanning year")
			panic(err)
		}
		years = append(years, year)
	}
	return years
}

func getPaymentTimestampsFromDB(idPayment int) (string, string) {
	result, err := database.DB.Query(fmt.Sprintf("SELECT CreatedAt, UpdatedAt FROM PaymentTable WHERE IdPayment = %d", idPayment))
	if err != nil {
		fmt.Println("error searching createdAt and updatedAt from DB")
		panic(err)
	}
	var createdAtUnformatted, updatedAtUnformatted time.Time
	for result.Next() {
		err = result.Scan(&createdAtUnformatted, &updatedAtUnformatted)
		if err != nil {
			fmt.Println("error scanning createdAt and updatedAt")
			panic(err)
		}
	}
	return formatTimeStamps(createdAtUnformatted, updatedAtUnformatted)
}
