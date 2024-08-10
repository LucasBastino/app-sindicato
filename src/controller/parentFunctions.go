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
	// errorMap := validateFieldsCaller(parent, c)
	parser := i.ParentParser{}
	parent = parserCaller(parser, c)
	// if len(errorMap) > 0 {
	// 	templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", errorMap)
	// 	return RenderHTML(c, templateData)
	// } else {
	insertModelCaller(parent)
	// templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", errorMap)
	return RenderHTML(c, fiber.Map{})
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
	parentEdited := parserCaller(parentParser, c)
	IdParent := getIdModelCaller(parent, c)
	parentEdited.IdParent = IdParent
	editModelCaller(parentEdited)
	// templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", nil)
	return RenderHTML(c, fiber.Map{})
}

func RenderParentTable(c *fiber.Ctx) error {
	// parents := searchModelsCaller(parent, c)
	// templateData := createTemplateDataCaller(parent, parent, parents, "src/views/tables/parentTable.html", nil)
	return RenderHTML(c, fiber.Map{})
}

func RenderParentFile(c *fiber.Ctx) error {
	// parent := searchOneModelByIdCaller(parent, c)
	// templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", nil)
	return RenderHTML(c, fiber.Map{})
}

func RenderCreateParentForm(c *fiber.Ctx) error {
	IdMember := getIdModelCaller(member, c)
	// creo un parent nuevo para que el form aparezca con campos vacios
	parent = models.Parent{}
	parent.IdMember = IdMember
	// templateData := createTemplateDataCaller(parent, parent, nil, "src/views/forms/createParentForm.html", nil)
	return RenderHTML(c, fiber.Map{})
}
