package controller

import (
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddEnterprise(c *fiber.Ctx) error {
	if err := validateFieldsCaller(models.Enterprise{}, c); err != nil {
		return er.CheckError(c, err)
	}
	e, err := parserCaller(i.EnterpriseParser{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	e, err = insertModelCaller(e)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	createdAt, updatedAt, err := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	years, err := getPaymentYears(e.IdEnterprise)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	data := fiber.Map{"enterprise": e, "numberOfMembers": 0, "role": role, "mode": "edit", "years": years, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("enterpriseFile", data)

}

func DeleteEnterprise(c *fiber.Ctx) error {
	IdEnterprise, err := getIdModelCaller(models.Enterprise{}, c)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	if IdEnterprise == 1 {
		// logear "error": "cannot delete enterprise 1"
		return er.CheckError(c, er.InsufficientPermisionsError)
	}
	e := models.Enterprise{IdEnterprise: IdEnterprise}
	members, err := getAllEnterpriseMembers(IdEnterprise)
	if err != nil {
		// logear el error
		return er.CheckError(c, err)
	}
	err = deleteModelCaller(e)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	err = setIdEnterpriseToOne(members)
	if err != nil {
		// logear el error
		return er.CheckError(c, err)
	}
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
	if err := validateFieldsCaller(models.Enterprise{}, c); err != nil {
		return er.CheckError(c, err)
	}
	e, err := parserCaller(i.EnterpriseParser{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	IdEnterprise, err := getIdModelCaller(e, c)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	e.IdEnterprise = IdEnterprise
	years, err := getPaymentYears(e.IdEnterprise)
	if err != nil {
		// logearlo
		return er.CheckError(c, err)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	e, err = updateModelCaller(e)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	numberOfMembers, err := getNumberOfMembers(e.IdEnterprise, "")
	if err != nil {
		// logearlo
		return er.CheckError(c, err)
	}
	createdAt, updatedAt, err := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
	if err != nil {
		// logearlo en algun lado
		return er.CheckError(c, err)
	}
	data := fiber.Map{"enterprise": e, "numberOfMembers": numberOfMembers, "role": role, "mode": "edit", "years": years, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("enterpriseFile", data)
}

func getPaymentYears(idEnterprise int) ([]string, error) {
	result, err := database.DB.Query("SELECT Year FROM PaymentTable WHERE IdEnterprise = ? GROUP BY Year ORDER BY YEAR DESC", idEnterprise)
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
	totalRows, err := getTotalRowsCaller(models.Enterprise{}, c)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.SendString(`<div class="no-result">No se encontraron empresas</div>`)
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		enterprises, searchKey, err := searchModelsCaller(models.Enterprise{}, c, offset)
		if err != nil {
			// guardar el err
			return er.CheckError(c, err)
		}

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
	e, err := searchOneModelByIdCaller(models.Enterprise{}, c)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	numberOfMembers, err := getNumberOfMembers(e.IdEnterprise, "")
	if err != nil {
		// logear el err
		return er.CheckError(c, err)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	if e.IdEnterprise == 1 {
		data := fiber.Map{"enterprise": e, "role": role, "numberOfMembers": numberOfMembers, "mode": "edit"}
		return c.Render("withoutEnterprise", data)
	} else {
		createdAt, updatedAt, err := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
		if err != nil {
			// logear el err
			return er.CheckError(c, err)
		}
		data := fiber.Map{"enterprise": e, "role": role, "numberOfMembers": numberOfMembers, "mode": "edit", "withPaymentTable": false, "createdAt": createdAt, "updatedAt": updatedAt}
		return c.Render("enterpriseFile", data)
	}
}

func getNumberOfMembers(IdEnterprise int, searchKey string) (int, error) {
	var totalRows int
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM MemberTable 
		WHERE IdEnterprise = ? AND (Name LIKE '%?%' OR LastName LIKE '%?%')`,
		IdEnterprise, searchKey, searchKey)
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
			if IdEnterprise, err := getIdModelCaller(models.Enterprise{}, c); err != nil {
				// guardar el err
				return 0, err
			} else {
				return IdEnterprise, nil
			}
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
		return er.CheckError(c, err)
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
	result, err := database.DB.Query(`
			SELECT * FROM MemberTable 
			WHERE IdEnterprise = ?
			AND (Name LIKE '%?%' OR LastName LIKE '%?%') LIMIT 10 OFFSET %d`, IdEnterprise, searchKey, searchKey, offset)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, er.QueryError
	}
	_, members, err := models.Member{}.ScanResult(result, false)
	if err != nil {
		// guardar el err
		return nil, err
	}
	return members, nil
}

func getAllEnterpriseMembers(IdEnterprise int) ([]models.Member, error) {
	result, err := database.DB.Query(`
			SELECT * FROM MemberTable WHERE IdEnterprise = ?`, IdEnterprise)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, er.QueryError
	}
	_, members, err := models.Member{}.ScanResult(result, false)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func setIdEnterpriseToOne(members []models.Member) error {
	// query, err := database.DB.Query("SET GLOBAL FOREIGN_KEY_CHECKS = 0")
	// if err!=nil{
	// 	fmt.Println("error setting foreing keys to 1")
	// 	panic(err)
	// }
	for _, m := range members {
		update, err := database.DB.Query(`
		UPDATE MemberTable SET IdEnterprise = '1' WHERE IdMember = ?`, m.IdMember)
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
	IdEnterprise, err := getIdModelCaller(models.Enterprise{}, c)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	var lastYear int
	result, err := database.DB.Query("SELECT MAX(Year) FROM PaymentTable WHERE IdEnterprise = ?", IdEnterprise)
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
	result2, err := database.DB.Query("SELECT Year FROM PaymentTable WHERE IdEnterprise = ? GROUP BY Year ORDER BY YEAR DESC", IdEnterprise)
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
		payments, _, err := searchModelsCaller(models.Payment{}, c, lastYear)
		if err != nil {
			// guardar el err
			return er.CheckError(c, err)
		}
		data := fiber.Map{"payments": payments, "idEnterprise": IdEnterprise, "role": role, "mode": "edit", "years": years, "year": lastYear}
		return c.Render("paymentTable", data)
	} else {
		payments, _, err := searchModelsCaller(models.Payment{}, c, params.Year)
		if err != nil {
			// guardar el err
			return er.CheckError(c, err)
		}
		data := fiber.Map{"payments": payments, "idEnterprise": IdEnterprise, "role": role, "mode": "edit", "years": years, "year": yearInt}
		return c.Render("paymentTable", data)
	}
}

func RenderEnterpriseTableSelect(c *fiber.Ctx) error {
	// calculo la cantidad de resultados
	totalRows, err := getTotalRowsCaller(models.Enterprise{}, c)
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.SendString(`<div class="no-result">No se encontraron empresas</div>`)
	} else {
		// si hay resultados...

		searchKey := c.FormValue("search-key")
		result, err := database.DB.Query(`
		SELECT
		*
		FROM EnterpriseTable 
		WHERE 
		Name LIKE '%?%' 
		ORDER BY Name ASC`,
			searchKey)
		if err != nil {
			// logear el err
			return er.CheckError(c, er.QueryError)
		}
		_, enterprises, err := models.Enterprise{}.ScanResult(result, false)
		if err != nil {
			// guardar el err
			return er.CheckError(c, err)
		}

		// creo un map con todas las variables
		role := c.Locals("claims").(jwt.MapClaims)["role"]
		// renderizo la tabla y le envio el map con las variables
		return c.Render("enterpriseTableSelect", fiber.Map{"enterprises": enterprises, "role": role})
	}
}

func GetAllEnterprisesId(c *fiber.Ctx) error {
	enterprisesId, err := models.GetAllEnterprisesIdFromDB()
	if err != nil {
		// guardar el err
		return er.CheckError(c, err)
	}
	return c.JSON(enterprisesId)
}
