package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/database"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateEnterprise(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Enterprise{}, c)
	e := parserCaller(i.EnterpriseParser{}, c)
	if len(errorMap) > 0 {
		// Si tiene errores renderizo nuevamente el form
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
		numberOfMembers := GetNumberOfMembers(e.IdEnterprise, "")
		data := fiber.Map{"enterprise": e, "numberOfMembers": numberOfMembers, "mode": "edit"}
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
		role := c.Locals("claims").(jwt.MapClaims)["role"]
		data["role"] = role
		// renderizo la tabla y le envio el map con las variables
		return c.Render("enterpriseTable", data)
	}
}

func RenderEnterpriseFile(c *fiber.Ctx) error {
	e := searchOneModelByIdCaller(models.Enterprise{}, c)
	numberOfMembers := GetNumberOfMembers(e.IdEnterprise, "")
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	data := fiber.Map{"enterprise": e, "role": role, "numberOfMembers": numberOfMembers, "mode": "edit"}
	return c.Render("enterpriseFile", data)
}

func GetNumberOfMembers(IdEnterprise int, searchKey string) int {
	var totalRows int
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM MemberTable 
		WHERE IdEnterprise = %d AND (Name LIKE '%%%s%%' OR LastName LIKE '%%%s%%')`,
		IdEnterprise, searchKey, searchKey))
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}
	return totalRows
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

	searchKey := c.FormValue("search-key")
	IdEnterprise := func() int {
		if c.Get("mode") == "edit" {
			return getIdModelCaller(models.Enterprise{}, c)
		} else if c.Get("mode") == "enterpriseMemberTable" {
			IdEnterpriseStr := c.Get("idEnterprise")
			IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
			if err != nil {
				fmt.Println("error converting IdEnterpriseStr to INT")
				panic(err)
			}
			return IdEnterprise
		} else {
			return 0
		}
	}()

	// calculo la cantidad de resultados
	totalRows := GetNumberOfMembers(IdEnterprise, searchKey)

	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.Render("searchWithNoResults", fiber.Map{})
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		searchKey := c.FormValue("search-key")
		members := GetEnterpriseMembers(IdEnterprise, searchKey, offset)

		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		// creo un map con todas las variables
		data := getFiberMapCaller(models.Member{}, members, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)
		data["mode"] = "enterpriseMemberTable"
		data["IdEnterprise"] = IdEnterprise
		role := c.Locals("claims").(jwt.MapClaims)["role"]
		data["role"] = role
		// renderizo la tabla y le envio el map con las variables
		return c.Render("memberTable", data)
	}
}

func GetEnterpriseMembers(IdEnterprise int, searchKey string, offset int) []models.Member {
	result, err := database.DB.Query(fmt.Sprintf(`
			SELECT * FROM MemberTable 
			WHERE IdEnterprise = %d
			AND (Name LIKE '%%%s%%' OR LastName LIKE '%%%s%%') LIMIT 10 OFFSET %d`, IdEnterprise, searchKey, searchKey, offset))
	if err != nil {
		fmt.Println("error searching member in DB")
		panic(err)
	}
	_, members := models.Member{}.ScanResult(result, false)
	return members
}
