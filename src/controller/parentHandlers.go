package controller

import (
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddParent(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Parent{}, c)
	p := parserCaller(i.ParentParser{}, c)
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	if len(errorMap) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorMap)
	} else {
		p = insertModelCaller(p)
		createdAt, updatedAt := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
		data := fiber.Map{"parent": p, "mode": "edit", "role": role, "createdAt": createdAt, "updatedAt": updatedAt}
		return c.Render("parentFile", data)
	}
}

func DeleteParent(c *fiber.Ctx) error {
	p := searchOneModelByIdCaller(models.Parent{}, c)
	deleteModelCaller(p)
	return RenderParentTable(c)
}

func EditParent(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Parent{}, c)
	p := parserCaller(i.ParentParser{}, c)
	IdParent := getIdModelCaller(p, c)
	p.IdParent = IdParent
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	if len(errorMap) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorMap)
	} else {
		p = updateModelCaller(p)
		createdAt, updatedAt := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
		data := fiber.Map{"parent": p, "mode": "edit", "role": role, "createdAt": createdAt, "updatedAt": updatedAt}
		return c.Render("parentFile", data)
	}
}

func RenderParentFile(c *fiber.Ctx) error {
	p := searchOneModelByIdCaller(models.Parent{}, c)
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	createdAt, updatedAt := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
	data := fiber.Map{"parent": p, "mode": "edit", "role": role, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("parentFile", data)
}

func RenderAddParentForm(c *fiber.Ctx) error {
	IdMember := getIdModelCaller(models.Member{}, c)
	// creo un parent nuevo para que el form aparezca con campos vacios
	p := models.Parent{IdMember: IdMember}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	data := fiber.Map{"parent": p, "mode": "add", "role": role}
	return c.Render("parentFile", data)
}
