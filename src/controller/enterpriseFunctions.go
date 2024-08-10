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
	// errorMap := validateFieldsCaller(enterprise, c)
	parser := i.EnterpriseParser{}
	enterprise = parserCaller(parser, c)
	// if len(errorMap) > 0 {
	// 	data := fiber.Map{"model": models.Enterprise{}, "template": "enterpriseFile", "errorMap": errorMap}
	// 	return RenderHTML(c, data)
	// } else {
	insertModelCaller(enterprise)
	// templateData := createTemplateDataCaller(enterprise, enterprise, nil, "src/views/files/enterpriseFile.html", errorMap)
	return RenderHTML(c, fiber.Map{})

}

func DeleteEnterprise(c *fiber.Ctx) error {
	IdEnterprise := getIdModelCaller(enterprise, c)
	enterprise.IdEnterprise = IdEnterprise
	deleteModelCaller(enterprise)
	return RenderEnterpriseTable(c)
}

func EditEnterprise(c *fiber.Ctx) error {
	enterpriseEdited := parserCaller(enterpriseParser, c)
	IdEnterprise := getIdModelCaller(enterprise, c)
	enterpriseEdited.IdEnterprise = IdEnterprise
	editModelCaller(enterpriseEdited)
	// templateData := createTemplateDataCaller(enterprise, enterpriseEdited, nil, "src/views/files/enterpriseFile.html", nil)
	return RenderHTML(c, fiber.Map{})
}

func RenderEnterpriseTable(c *fiber.Ctx) error {
	// enterprises := searchModelsCaller(enterprise, c)
	// templateData := createTemplateDataCaller(enterprise, enterprise, enterprises, "src/views/tables/enterpriseTable.html", nil)
	return RenderHTML(c, fiber.Map{})
}

func RenderEnterpriseFile(c *fiber.Ctx) error {
	// enterprise := searchOneModelByIdCaller(enterprise, c)
	// templateData := createTemplateDataCaller(enterprise, enterprise, nil, "src/views/files/enterpriseFile.html", nil)
	return RenderHTML(c, fiber.Map{})
}

func RenderCreateEnterpriseForm(c *fiber.Ctx) error {
	// le paso un enterprise vacio para que los campos del form aparezcan vacios
	// templateData := createTemplateDataCaller(enterprise, models.Enterprise{}, nil, "src/views/forms/createEnterpriseForm.html", nil)
	return RenderHTML(c, fiber.Map{})
}
