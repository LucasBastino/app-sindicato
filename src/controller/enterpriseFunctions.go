package controller

import (
	"fmt"
	"log"

	"github.com/LucasBastino/app-sindicato/src/database"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateEnterprise(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Enterprise{}, c)
	e := parserCaller(i.EnterpriseParser{}, c)
	if len(errorMap) > 0 {
		data := fiber.Map{"enterprise": e, "mode": "create", "errorMap": errorMap}
		return c.Render("enterpriseFile", data)
	} else {
		e = insertModelCaller(e)
		data := fiber.Map{"enterprise": e, "mode": "edit"}
		return c.Render("enterpriseFile", data)
	}
}

func DeleteEnterprise(c *fiber.Ctx) error {
	IdEnterprise := getIdModelCaller(models.Enterprise{}, c)
	e := models.Enterprise{IdEnterprise: IdEnterprise}
	deleteModelCaller(e)
	return RenderEnterpriseTable(c)
}

func EditEnterprise(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Enterprise{}, c)
	e := parserCaller(i.EnterpriseParser{}, c)
	IdEnterprise := getIdModelCaller(e, c)
	e.IdEnterprise = IdEnterprise
	if len(errorMap) > 0 {
		data := fiber.Map{"enterprise": e, "mode": "edit", "errorMap": errorMap}
		return c.Render("enterpriseFile", data)
	} else {
		editModelCaller(e)
		data := fiber.Map{"enterprise": e, "mode": "edit"}
		return c.Render("enterpriseFile", data)
	}
}

func RenderEnterpriseTable(c *fiber.Ctx) error {
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
	e := searchOneModelByIdCaller(models.Enterprise{}, c)
	data := fiber.Map{"enterprise": e, "mode": "edit"}
	return c.Render("enterpriseFile", data)
}

func RenderCreateEnterpriseForm(c *fiber.Ctx) error {
	// le paso un enterprise vacio para que los campos del form aparezcan vacios
	data := fiber.Map{"enterprise": models.Enterprise{}, "mode": "create"}
	return c.Render("enterpriseFile", data)
}

func RenderEnterpriseMembers(c *fiber.Ctx) error {
	// Busco todos los members por key de la empresa y renderizo la tabla de miembros

	// obtengo la currentPage del path
	currentPage := GetPageFromPath(c)

	// calculo la cantidad de resultados
	var totalRows int
	searchKey := c.FormValue("search-key")
	params := struct {
		IdEnterprise int `params:"IdEnterprise"`
	}{}

	c.ParamsParser(&params)
	IdEnterprise := params.IdEnterprise
	row := database.DB.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM MemberTable WHERE IdEnterprise = %d AND Name LIKE '%%%s%%'`, IdEnterprise, searchKey))
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("cantidad de resultados:", totalRows)

	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.Render("searchWithNoResults", fiber.Map{})
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		searchKey := c.FormValue("search-key")
		result, err := database.DB.Query(fmt.Sprintf(`SELECT * FROM MemberTable WHERE IdEnterprise = %d AND Name LIKE '%%%s%%' LIMIT 10 OFFSET %d`, IdEnterprise, searchKey, offset))
		if err != nil {
			fmt.Println("error searching member in DB")
			panic(err)
		}
		_, members := models.Member{}.ScanResult(result, false)
		// fmt.Println(members)

		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		// creo un map con todas las variables
		data := getFiberMapCaller(models.Member{}, members, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)

		// renderizo la tabla y le envio el map con las variables
		return c.Render("memberTable", data)
	}
}
