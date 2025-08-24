package controller

import (
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/LucasBastino/app-sindicato/src/errors/errorHandler"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddParent(c *fiber.Ctx) error {
	if customErr := validateFieldsCaller(models.Parent{}, c); (customErr != customError.CustomError{}) {
		return errorHandler.HandleError(c, &customErr)
	}
	p, customErr := parserCaller(i.ParentParser{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}

	p, customErr = insertModelCaller(p)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	createdAt, updatedAt, customErr := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	data := fiber.Map{"parent": p, "mode": "edit", "createdAt": createdAt, "updatedAt": updatedAt}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteParent"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeParent"]
	return c.Status(fiber.StatusCreated).Render("parentFile", data)

}

func DeleteParent(c *fiber.Ctx) error {
	p, customErr := searchOneModelByIdCaller(models.Parent{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	deleteModelCaller(p)
	return RenderParentTable(c)
}

func EditParent(c *fiber.Ctx) error {
	if customErr := validateFieldsCaller(models.Parent{}, c); (customErr != customError.CustomError{}) {
		return errorHandler.HandleError(c, &customErr)
	}
	p, customErr := parserCaller(i.ParentParser{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	IdParent, customErr := getIdModelCaller(p, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	p.IdParent = IdParent
	p, customErr = updateModelCaller(p)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	createdAt, updatedAt, customErr := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	data := fiber.Map{"parent": p, "mode": "edit", "createdAt": createdAt, "updatedAt": updatedAt}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteParent"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeParent"]
	return c.Render("parentFile", data)

}

func RenderParentFile(c *fiber.Ctx) error {
	p, customErr := searchOneModelByIdCaller(models.Parent{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	createdAt, updatedAt, customErr := formatTimeStamps(p.CreatedAt, p.UpdatedAt)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	data := fiber.Map{"parent": p, "mode": "edit", "createdAt": createdAt, "updatedAt": updatedAt}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteParent"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeParent"]
	return c.Render("parentFile", data)
}

func RenderAddParentForm(c *fiber.Ctx) error {
	IdMember, customErr := getIdModelCaller(models.Member{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	// creo un parent nuevo para que el form aparezca con campos vacios
	p := models.Parent{IdMember: IdMember}
	data := fiber.Map{"parent": p, "mode": "add"}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteParent"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeParent"]
	return c.Render("parentFile", data)
}
