package controller

import (
	"fmt"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddEnterprise(c *fiber.Ctx) error {
	errorMap, err := validateFieldsCaller(models.Enterprise{}, c)
	if err != nil {
		fmt.Println(errorMap)
		// borrar esto ↑ y logear en algun lado el error de validacion con el errorMap
		return er.CheckError(c, er.ValidationError)
	}
	e := parserCaller(i.EnterpriseParser{}, c)
	e = insertModelCaller(e)
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	createdAt, updatedAt, err := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
	if err != nil {
		// guardar el err o aca o alla
		return er.CheckError(c, er.FormatError)
	}
	years := getPaymentYears(e.IdEnterprise)
	data := fiber.Map{"enterprise": e, "numberOfMembers": 0, "role": role, "mode": "edit", "years": years, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("enterpriseFile", data)

}

func DeleteEnterprise(c *fiber.Ctx) error {
	IdEnterprise := getIdModelCaller(models.Enterprise{}, c)
	if IdEnterprise == 1 {
		// logear "error": "cannot delete enterprise 1"
		return er.CheckError(c, er.InsufficientPermisionsError)
	}
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
		// log error with deleting mode
		return er.CheckError(c, er.InternalServerError)
	}

}

func EditEnterprise(c *fiber.Ctx) error {
	errorMap, err := validateFieldsCaller(models.Enterprise{}, c)
	if err != nil {
		fmt.Println(errorMap)
		// borrar esto ↑ y logear en algun lado el error de validacion con el errorMap
		return er.CheckError(c, er.ValidationError)
	}
	e := parserCaller(i.EnterpriseParser{}, c)
	IdEnterprise := getIdModelCaller(e, c)
	e.IdEnterprise = IdEnterprise
	years, err := getPaymentYears(e.IdEnterprise)
	if err != nil {
		// logearlo
		return er.CheckError(c, er.QueryError)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	e = updateModelCaller(e)
	numberOfMembers := getNumberOfMembers(e.IdEnterprise, "")
	createdAt, updatedAt, err := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
	if err != nil {
		// logearlo en algun lado
		return er.CheckError(c, er.FormatError)
	}
	data := fiber.Map{"enterprise": e, "numberOfMembers": numberOfMembers, "role": role, "mode": "edit", "years": years, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("enterpriseFile", data)
}

func getPaymentYears(idEnterprise int) ([]string, error) {
	result, err := database.DB.Query(fmt.Sprintf("SELECT Year FROM PaymentTable WHERE IdEnterprise = '%d' GROUP BY Year ORDER BY YEAR DESC", idEnterprise))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, er.QueryError
	}

	var years []string
	var year string
	for result.Next() {
		err = result.Scan(&year)
		if err != nil {
			er.ScanError.Msg = err.Error()
			return nil, er.ScanError
		}
		years = append(years, year)
	}
	return years, nil
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
	numberOfMembers, err := getNumberOfMembers(e.IdEnterprise, "")
	if err != nil {
		// logear el err
		return er.CheckError(c, er.ScanError)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	if e.IdEnterprise == 1 {
		data := fiber.Map{"enterprise": e, "role": role, "numberOfMembers": numberOfMembers, "mode": "edit"}
		return c.Render("withoutEnterprise", data)
	} else {
		createdAt, updatedAt, err := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
		if err != nil {
			// logear el err
			return er.CheckError(c, er.FormatError)
		}
		data := fiber.Map{"enterprise": e, "role": role, "numberOfMembers": numberOfMembers, "mode": "edit", "withPaymentTable": false, "createdAt": createdAt, "updatedAt": updatedAt}
		return c.Render("enterpriseFile", data)
	}
}

func getNumberOfMembers(IdEnterprise int, searchKey string) (int, error) {
	var totalRows int
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM MemberTable 
		WHERE IdEnterprise = %d AND (Name LIKE '%%%s%%' OR LastName LIKE '%%%s%%')`,
		IdEnterprise, searchKey, searchKey))
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		er.ScanError.Msg = err.Error()
		return 0, er.ScanError
	}
	return totalRows, nil
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

	IdEnterprise, err := func() (int, error) {
		// c.Get devuelve un valor del header
		if c.Get("mode") == "edit" {
			return getIdModelCaller(models.Enterprise{}, c), nil
		} else if c.Get("mode") == "enterpriseMemberTable" {
			IdEnterpriseStr := c.Get("idEnterprise")
			IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
			if err != nil {
				er.StrConvError.Msg = err.Error()
				return 0, er.StrConvError
			}
			return IdEnterprise, nil
		} else {
			return 0, nil
		}
	}()
	if err != nil {
		// logear el error
		return er.CheckError(c, er.StrConvError)
	}

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
	totalRows, err := getNumberOfMembers(IdEnterprise, searchKey)
	if err != nil {
		// logear el err
		return er.CheckError(c, er.ScanError)
	}

	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.SendString(`<div class="no-result">No se encontraron afiliados</div>`)
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		members, err := getEnterpriseMembers(IdEnterprise, searchKey, offset)
		if err != nil {
			// logear el err
			return er.CheckError(c, er.QueryError)
		}

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

func getEnterpriseMembers(IdEnterprise int, searchKey string, offset int) ([]models.Member, error) {
	result, err := database.DB.Query(fmt.Sprintf(`
			SELECT * FROM MemberTable 
			WHERE IdEnterprise = %d
			AND (Name LIKE '%%%s%%' OR LastName LIKE '%%%s%%') LIMIT 10 OFFSET %d`, IdEnterprise, searchKey, searchKey, offset))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, er.QueryError
	}
	_, members := models.Member{}.ScanResult(result, false)
	return members, nil
}

func getAllEnterpriseMembers(IdEnterprise int) ([]models.Member, error) {
	result, err := database.DB.Query(fmt.Sprintf(`
			SELECT * FROM MemberTable WHERE IdEnterprise = %d`, IdEnterprise))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, er.QueryError
	}
	_, members := models.Member{}.ScanResult(result, false)
	return members, nil
}

func setIdEnterpriseToOne(members []models.Member) error {
	// query, err := database.DB.Query("SET GLOBAL FOREIGN_KEY_CHECKS = 0")
	// if err!=nil{
	// 	fmt.Println("error setting foreing keys to 1")
	// 	panic(err)
	// }
	for _, m := range members {
		update, err := database.DB.Query(fmt.Sprintf(`
		UPDATE MemberTable SET IdEnterprise = '1' WHERE IdMember = %d`, m.IdMember))
		if err != nil {
			er.QueryError.Msg = err.Error()
			return er.QueryError
		}
		update.Close()
		// query.Close()
	}
	return nil
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
		// logear el err
		return er.CheckError(c, er.QueryError)
	}
	for result.Next() {
		err = result.Scan(&lastYear)
	}
	if err != nil {
		// logear el err
		return er.CheckError(c, er.ScanError)
	}
	result2, err := database.DB.Query(fmt.Sprintf("SELECT Year FROM PaymentTable WHERE IdEnterprise = '%d' GROUP BY Year ORDER BY YEAR DESC", IdEnterprise))
	if err != nil {
		// logear el err
		return er.CheckError(c, er.QueryError)
	}

	var years []string
	var year string
	for result2.Next() {
		err = result2.Scan(&year)
		if err != nil {
			// logear el err
			return er.CheckError(c, er.ScanError)
		}
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
			// logear el err
			return er.CheckError(c, er.QueryError)
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
