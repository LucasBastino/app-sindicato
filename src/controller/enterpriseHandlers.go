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

func AddEnterprise(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Enterprise{}, c)
	e := parserCaller(i.EnterpriseParser{}, c)
	if len(errorMap) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorMap)
	} else {
		e = insertModelCaller(e)
		data := fiber.Map{"enterprise": e, "mode": "edit"}
		return c.Render("enterpriseFile", data)
	}
}

func DeleteEnterprise(c *fiber.Ctx) error {
	IdEnterprise := getIdModelCaller(models.Enterprise{}, c)
	if IdEnterprise == 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot delete enterprise 1"})
	} else {
		e := models.Enterprise{IdEnterprise: IdEnterprise}
		members := getAllEnterpriseMembers(IdEnterprise)
		deleteModelCaller(e)
		setIdEnterpriseToOne(members)
		switch c.Get("mode") {
		case "table":
			return RenderEnterpriseTable(c)
		case "edit":
			return c.Render("index", fiber.Map{"withEnterpriseTable": true})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error with deleting mode"})
		}
	}
}

func EditEnterprise(c *fiber.Ctx) error {
	errorMap := validateFieldsCaller(models.Enterprise{}, c)
	e := parserCaller(i.EnterpriseParser{}, c)
	IdEnterprise := getIdModelCaller(e, c)
	e.IdEnterprise = IdEnterprise
	result, err := database.DB.Query(fmt.Sprintf("SELECT Year FROM PaymentTable WHERE IdEnterprise = '%d' GROUP BY Year ORDER BY YEAR DESC", e.IdEnterprise))
	if err != nil {
		fmt.Println("error searching different Years in PaymentTable")
		panic(err)
	}

	var years []string
	var year string
	for result.Next() {
		result.Scan(&year)
		years = append(years, year)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]

	if len(errorMap) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorMap)
	} else {
		e = updateModelCaller(e)
		fmt.Printf("%+v", e)
		numberOfMembers := getNumberOfMembers(e.IdEnterprise, "")
		createdAt, updatedAt := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
		data := fiber.Map{"enterprise": e, "numberOfMembers": numberOfMembers, "role": role, "mode": "edit", "years": years, "createdAt": createdAt, "updatedAt": updatedAt}
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
		return c.SendString(`<div class="no-result">No se encontraron empresas</div>`)
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
	createdAt, updatedAt := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
	numberOfMembers := getNumberOfMembers(e.IdEnterprise, "")
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	data := fiber.Map{"enterprise": e, "role": role, "numberOfMembers": numberOfMembers, "mode": "edit", "withPaymentTable": false, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("enterpriseFile", data)
}

func getNumberOfMembers(IdEnterprise int, searchKey string) int {
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

func RenderAddEnterpriseForm(c *fiber.Ctx) error {
	// le paso un enterprise vacio para que los campos del form aparezcan vacios
	data := fiber.Map{"enterprise": models.Enterprise{}, "mode": "add"}
	return c.Render("enterpriseFile", data)
}

func RenderEnterpriseMembers(c *fiber.Ctx) error {
	// Busco todos los members por key de la empresa y renderizo la tabla de miembros

	// obtengo la currentPage del path
	currentPage := GetPageFromPath(c)

	IdEnterprise := func() int {
		// c.Get devuelve un valor del header
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

	var searchKey string

	if c.Get("deleteMode") == "true" {
		// si estamos en deleteMode que el searchKey lo saque del header, ya que no se lo voy a mandar por el form
		// asi cuando elimino un miembro se quedan los miembros que busque antes menos el que elimine
		searchKey = c.Get("searchKey")
	} else {
		// sino se lo mando por el form normalmente
		searchKey = c.FormValue("search-key")
	}

	// calculo la cantidad de resultados
	totalRows := getNumberOfMembers(IdEnterprise, searchKey)

	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.SendString(`<div class="no-result">No se encontraron afiliados</div>`)
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		members := getEnterpriseMembers(IdEnterprise, searchKey, offset)

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

func getEnterpriseMembers(IdEnterprise int, searchKey string, offset int) []models.Member {
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

func getAllEnterpriseMembers(IdEnterprise int) []models.Member {
	result, err := database.DB.Query(fmt.Sprintf(`
			SELECT * FROM MemberTable WHERE IdEnterprise = %d`, IdEnterprise))
	if err != nil {
		fmt.Println("error searching member in DB")
		panic(err)
	}
	_, members := models.Member{}.ScanResult(result, false)
	return members
}

func setIdEnterpriseToOne(members []models.Member) {
	// query, err := database.DB.Query("SET GLOBAL FOREIGN_KEY_CHECKS = 0")
	// if err!=nil{
	// 	fmt.Println("error setting foreing keys to 1")
	// 	panic(err)
	// }
	for _, m := range members {
		update, err := database.DB.Query(fmt.Sprintf(`
		UPDATE MemberTable SET IdEnterprise = '1' WHERE IdMember = %d`, m.IdMember))
		if err != nil {
			fmt.Println("error setting IdEnterprise to '1'")
			panic(err)
		}
		update.Close()
		// query.Close()
	}
}

func RenderEnterprisePaymentsTable(c *fiber.Ctx) error {
	params := struct {
		Year int `params:"year"`
	}{}
	c.ParamsParser(&params)
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	IdEnterprise := getIdModelCaller(models.Enterprise{}, c)
	var lastYear int
	result, err := database.DB.Query(fmt.Sprintf("SELECT MAX(Year) FROM PaymentTable WHERE IdEnterprise = '%d'", IdEnterprise))
	if err != nil {
		fmt.Println("error selecting last year")
	}
	for result.Next() {
		err = result.Scan(&lastYear)
	}
	if err != nil {
		fmt.Println("error scaning lastyear")
	}
	result2, err := database.DB.Query(fmt.Sprintf("SELECT Year FROM PaymentTable WHERE IdEnterprise = '%d' GROUP BY Year ORDER BY YEAR DESC", IdEnterprise))
	if err != nil {
		fmt.Println("error searching different Years in PaymentTable")
		panic(err)
	}

	var years []string
	var year string
	for result2.Next() {
		result2.Scan(&year)
		years = append(years, year)
	}
	yearInt := params.Year
	if params.Year == 0 {
		payments, _ := searchModelsCaller(models.Payment{}, c, lastYear)
		data := fiber.Map{"payments": payments, "idEnterprise": IdEnterprise, "role": role, "mode": "edit", "years": years, "year": lastYear}
		return c.Render("paymentTable", data)
	} else {
		payments, _ := searchModelsCaller(models.Payment{}, c, params.Year)
		data := fiber.Map{"payments": payments, "idEnterprise": IdEnterprise, "role": role, "mode": "edit", "years": years, "year": yearInt}
		return c.Render("paymentTable", data)
	}
}

func RenderEnterpriseTableSelect(c *fiber.Ctx) error {
	// calculo la cantidad de resultados
	totalRows := getTotalRowsCaller(models.Enterprise{}, c)
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.SendString(`<div class="no-result">No se encontraron empresas</div>`)
	} else {
		// si hay resultados...

		searchKey := c.FormValue("search-key")
		result, err := database.DB.Query(fmt.Sprintf(`
		SELECT
		*
		FROM EnterpriseTable 
		WHERE 
		Name LIKE '%%%s%%' 
		ORDER BY Name ASC`,
			searchKey))
		if err != nil {
			fmt.Println("error searching Enterprise in DB")
			panic(err)
		}
		_, enterprises := models.Enterprise{}.ScanResult(result, false)

		// creo un map con todas las variables
		role := c.Locals("claims").(jwt.MapClaims)["role"]
		// renderizo la tabla y le envio el map con las variables
		return c.Render("enterpriseTableSelect", fiber.Map{"enterprises": enterprises, "role": role})
	}
}

func GetAllEnterprisesId(c *fiber.Ctx) error {
	enterprisesId := models.GetAllEnterprisesIdFromDB()
	return c.JSON(enterprisesId)
}
