package controller

import (
	er "github.com/LucasBastino/app-sindicato/src/errors"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddParent(c *fiber.Ctx) error {
	if err := validateFieldsCaller(models.Parent{}, c); err != nil {
		return er.CheckError(c, err)
	}
	p, err := parserCaller(i.ParentParser{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]

	p, err = insertModelCaller(p)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	createdAt, updatedAt, err := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	data := fiber.Map{"parent": p, "mode": "edit", "role": role, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("parentFile", data)

}

func DeleteParent(c *fiber.Ctx) error {
	p, err := searchOneModelByIdCaller(models.Parent{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	deleteModelCaller(p)
	return RenderParentTable(c)
}

func EditParent(c *fiber.Ctx) error {
	if err := validateFieldsCaller(models.Parent{}, c); err != nil {
		return er.CheckError(c, err)
	}
	p, err := parserCaller(i.ParentParser{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	IdParent, err := getIdModelCaller(p, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	p.IdParent = IdParent
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	p, err = updateModelCaller(p)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	createdAt, updatedAt, err := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	data := fiber.Map{"parent": p, "mode": "edit", "role": role, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("parentFile", data)

}

func RenderParentFile(c *fiber.Ctx) error {
	p, err := searchOneModelByIdCaller(models.Parent{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	createdAt, updatedAt, err := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	data := fiber.Map{"parent": p, "mode": "edit", "role": role, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("parentFile", data)
}

func RenderAddParentForm(c *fiber.Ctx) error {
	IdMember, err := getIdModelCaller(models.Member{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	// creo un parent nuevo para que el form aparezca con campos vacios
	p := models.Parent{IdMember: IdMember}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	data := fiber.Map{"parent": p, "mode": "add", "role": role}
	return c.Render("parentFile", data)
}
