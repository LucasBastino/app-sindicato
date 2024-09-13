package controller

import (
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateEnterprise(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	errorMap := validateFieldsCaller(models.Enterprise{}, c)
	e := parserCaller(i.EnterpriseParser{}, c)
	if len(errorMap) > 0 {
		data := fiber.Map{"enterprise": e, "errorMap": errorMap}
		return c.Render("createEnterpriseForm", data)
	} else {
		e = insertModelCaller(e)
		data := fiber.Map{"enterprise": e}
		return c.Render("enterpriseFile", data)
	}
}

func DeleteEnterprise(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	IdEnterprise := getIdModelCaller(models.Enterprise{}, c)
	e := models.Enterprise{IdEnterprise: IdEnterprise}
	deleteModelCaller(e)
	return RenderEnterpriseTable(c)
}

func EditEnterprise(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// falta agregar la validacion
	e := parserCaller(i.EnterpriseParser{}, c)
	IdEnterprise := getIdModelCaller(e, c)
	e.IdEnterprise = IdEnterprise
	editModelCaller(e)
	data := fiber.Map{"enterprise": e}
	return c.Render("enterpriseFile", data)
}

func RenderEnterpriseTable(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// obtengo la currentPage del path
	currentPage := GetPageFromPath(c)
	// calculo la cantidad de resultados
	totalRows := getTotalRowsCaller(models.Enterprise{}, c)
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.Render("searchWithNoResults", fiber.Map{})
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		enterprises, searchKey := searchModelsCaller(models.Enterprise{}, c, offset)

		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		// creo un map con todas las variables
		data := getFiberMapCaller(models.Enterprise{}, enterprises, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)

		// renderizo la tabla y le envio el map con las variables
		return c.Render("enterpriseTable", data)
	}
}

func RenderEnterpriseFile(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	e := searchOneModelByIdCaller(models.Enterprise{}, c)
	data := fiber.Map{"enterprise": e}
	return c.Render("enterpriseFile", data)
}

func RenderCreateEnterpriseForm(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// le paso un enterprise vacio para que los campos del form aparezcan vacios
	data := fiber.Map{"enterprise": models.Enterprise{}}
	return c.Render("createEnterpriseForm", data)
}
