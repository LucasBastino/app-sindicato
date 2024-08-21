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
	// obtengo la currentPage del path
	currentPage := GetPageFromPath(c)
	// calculo la cantidad de resultados
	totalRows := getTotalRowsCaller(enterprise, c)
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.Render("searchWithNoResults", fiber.Map{})
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		enterprises, searchKey := searchModelsCaller(enterprise, c, offset)

		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		// creo un map con todas las variables
		mapData := getFiberMapCaller(enterprise, enterprises, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)

		// renderizo la tabla y le envio el map con las variables
		return c.Render("enterpriseTable", mapData)
	}
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
