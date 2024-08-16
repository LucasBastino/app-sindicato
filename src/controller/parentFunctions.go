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
	parents := searchModelsCaller(parent, c)
	data := fiber.Map{"model": parent, "parents": parents}
	return c.Render("parentTable", data)
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
