package controller

import (
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

var (
	parent       models.Parent
	parentParser i.ParentParser
)

func CreateParent(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(parent, c)
	parser := i.ParentParser{}
	parent = parserCaller(parser, c)
	if len(errorMap) > 0 {
		data := fiber.Map{"model": parent, "parent": parent, "errorMap": errorMap}
		return c.Render("createParentForm", data)
	} else {
		insertModelCaller(parent)
		data := fiber.Map{"model": parent, "parent": parent}
		return c.Render("parentFile", data)
	}
}

func DeleteParent(c *fiber.Ctx) error {
	parent = searchOneModelByIdCaller(parent, c)
	deleteModelCaller(parent)
	// renderiza de nuevo la tabla
	// falla aca abajo, fijarse. Tambien probar agregar familiares
	member = searchOneModelByIdCaller(member, c)
	return RenderMemberParents(c)
}

func EditParent(c *fiber.Ctx) error {
	// falta hacer la validacion
	parentEdited := parserCaller(parentParser, c)
	IdParent := getIdModelCaller(parent, c)
	parentEdited.IdParent = IdParent
	editModelCaller(parentEdited)
	data := fiber.Map{"model": parent, "parent": parent}
	return c.Render("parentFile", data)
}

func RenderParentTable(c *fiber.Ctx) error {
	// obtengo la currentPage del path
	currentPage := GetPageFromPath(c)
	// calculo la cantidad de resultados
	totalRows := getTotalRowsCaller(parent, c)
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.Render("searchWithNoResults", fiber.Map{})
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		parents, searchKey := searchModelsCaller(parent, c, offset)

		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		// creo un map con todas las variables
		mapData := getFiberMapCaller(parent, parents, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)

		// renderizo la tabla y le envio el map con las variables
		return c.Render("parentTable", mapData)
	}
}

func RenderParentFile(c *fiber.Ctx) error {
	parent := searchOneModelByIdCaller(parent, c)
	data := fiber.Map{"model": parent, "parent": parent}
	return c.Render("parentFile", data)
}

func RenderCreateParentForm(c *fiber.Ctx) error {
	IdMember := getIdModelCaller(member, c)
	// creo un parent nuevo para que el form aparezca con campos vacios
	parent = models.Parent{}
	parent.IdMember = IdMember
	data := fiber.Map{"model": parent, "parent": parent}
	return c.Render("createParentForm", data)
}
