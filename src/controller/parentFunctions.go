package controller

import (
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateParent(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Parent{}, c)
	p := parserCaller(i.ParentParser{}, c)
	if len(errorMap) > 0 {
		data := fiber.Map{"parent": p, "mode": "create", "errorMap": errorMap}
		return c.Render("parentFile", data)
	} else {
		p = insertModelCaller(p)
		data := fiber.Map{"parent": p, "mode": "edit"}
		return c.Render("parentFile", data)
	}
}

func DeleteParent(c *fiber.Ctx) error {
	p := searchOneModelByIdCaller(models.Parent{}, c)
	deleteModelCaller(p)
	return RenderMemberParents(c)
}

func EditParent(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Parent{}, c)
	p := parserCaller(i.ParentParser{}, c)
	IdParent := getIdModelCaller(p, c)
	p.IdParent = IdParent
	if len(errorMap) > 0 {
		data := fiber.Map{"parent": p, "mode": "edit", "errorMap": errorMap}
		return c.Render("parentFile", data)
	} else {
		editModelCaller(p)
		data := fiber.Map{"parent": p, "mode": "edit"}
		return c.Render("parentFile", data)
	}
}

func RenderParentFile(c *fiber.Ctx) error {
	p := searchOneModelByIdCaller(models.Parent{}, c)
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	data := fiber.Map{"parent": p, "role": role, "mode": "edit"}
	return c.Render("parentFile", data)
}

func RenderCreateParentForm(c *fiber.Ctx) error {
	IdMember := getIdModelCaller(models.Member{}, c)
	// creo un parent nuevo para que el form aparezca con campos vacios
	p := models.Parent{IdMember: IdMember}
	data := fiber.Map{"parent": p, "mode": "create"}
	return c.Render("parentFile", data)
}
