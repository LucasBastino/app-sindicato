package controller

import (
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateParent(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	errorMap := validateFieldsCaller(models.Parent{}, c)
	p := parserCaller(i.ParentParser{}, c)
	if len(errorMap) > 0 {
		data := fiber.Map{"parent": p, "errorMap": errorMap}
		return c.Render("createParentForm", data)
	} else {
		p = insertModelCaller(p)
		data := fiber.Map{"parent": p}
		return c.Render("parentFile", data)
	}
}

func DeleteParent(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	p := searchOneModelByIdCaller(models.Parent{}, c)
	deleteModelCaller(p)
	// renderiza de nuevo la tabla
	// falla aca abajo, fijarse. Tambien probar agregar familiares
	return RenderMemberParents(c)
}

func EditParent(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// falta hacer la validacion
	p := parserCaller(i.ParentParser{}, c)
	IdParent := getIdModelCaller(p, c)
	p.IdParent = IdParent
	editModelCaller(p)
	data := fiber.Map{"parent": p}
	return c.Render("parentFile", data)
}

func RenderParentTable(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// obtengo la currentPage del path
	currentPage := GetPageFromPath(c)
	// calculo la cantidad de resultados
	totalRows := getTotalRowsCaller(models.Parent{}, c)
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.Render("searchWithNoResults", fiber.Map{})
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		parents, searchKey := searchModelsCaller(models.Parent{}, c, offset)

		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		// creo un map con todas las variables
		mapData := getFiberMapCaller(models.Parent{}, parents, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)

		// renderizo la tabla y le envio el map con las variables
		return c.Render("parentTable", mapData)
	}
}

func RenderParentFile(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	p := searchOneModelByIdCaller(models.Parent{}, c)
	data := fiber.Map{"parent": p}
	return c.Render("parentFile", data)
}

func RenderCreateParentForm(c *fiber.Ctx) error {
	err := validateToken(c.Cookies("Authorization"))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	IdMember := getIdModelCaller(models.Member{}, c)
	// creo un parent nuevo para que el form aparezca con campos vacios
	p := models.Parent{IdMember: IdMember}
	data := fiber.Map{"parent": p}
	return c.Render("createParentForm", data)
}
