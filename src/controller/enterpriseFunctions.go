package controller

import (
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

var (
	enterprise       models.Enterprise
	enterpriseParser i.EnterpriseParser
)

func CreateEnterprise(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(enterprise, c)
	parser := i.EnterpriseParser{}
	enterprise = parserCaller(parser, c)
	if len(errorMap) > 0 {
		data := fiber.Map{"model": enterprise, "enterprise": enterprise, "errorMap": errorMap}
		return c.Render("createEnterpriseForm", data)
	} else {
		insertModelCaller(enterprise)
		data := fiber.Map{"model": enterprise, "enterprise": enterprise}
		return c.Render("enterpriseFile", data)
	}
}

func DeleteEnterprise(c *fiber.Ctx) error {
	IdEnterprise := getIdModelCaller(enterprise, c)
	enterprise.IdEnterprise = IdEnterprise
	deleteModelCaller(enterprise)
	return RenderEnterpriseTable(c)
}

func EditEnterprise(c *fiber.Ctx) error {
	// falta agregar la validacion
	enterpriseEdited := parserCaller(enterpriseParser, c)
	IdEnterprise := getIdModelCaller(enterprise, c)
	enterpriseEdited.IdEnterprise = IdEnterprise
	editModelCaller(enterpriseEdited)
	data := fiber.Map{"model": enterprise, "enterprise": enterprise}
	return c.Render("enterpriseFile", data)
}

func RenderEnterpriseTable(c *fiber.Ctx) error {
	enterprises := searchModelsCaller(enterprise, c, 5)
	data := fiber.Map{"model": enterprise, "enterprises": enterprises}
	return c.Render("enterpriseTable", data)
}

func RenderEnterpriseFile(c *fiber.Ctx) error {
	enterprise := searchOneModelByIdCaller(enterprise, c)
	data := fiber.Map{"model": enterprise, "enterprise": enterprise}
	return c.Render("enterpriseFile", data)
}

func RenderCreateEnterpriseForm(c *fiber.Ctx) error {
	// le paso un enterprise vacio para que los campos del form aparezcan vacios
	data := fiber.Map{"model": enterprise, "enterprise": models.Enterprise{}}
	return c.Render("createEnterpriseForm", data)
}
