package controller

import (
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/LucasBastino/app-sindicato/src/errors/errorHandler"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddEnterprise(c *fiber.Ctx) error {
	if customErr := validateFieldsCaller(models.Enterprise{}, c); (customErr != customError.CustomError{}) {
		return errorHandler.HandleError(c, &customErr)
	}
	e, customErr := parserCaller(i.EnterpriseParser{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	e, customErr = insertModelCaller(e)
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	createdAt, updatedAt, customErr := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	years, customErr := getPaymentYears(e.IdEnterprise)
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	data := fiber.Map{"enterprise": e, "numberOfMembers": 0, "mode": "edit", "years": years, "createdAt": createdAt, "updatedAt": updatedAt}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeEnterprise"]
	return c.Status(fiber.StatusCreated).Render("enterpriseFile", data)

}

func DeleteEnterprise(c *fiber.Ctx) error {
	IdEnterprise, customErr := getIdModelCaller(models.Enterprise{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	if IdEnterprise == 1 {
		// logear "error": "cannot delete enterprise 1"
		return errorHandler.HandleError(c, &customErr)
	}
	e := models.Enterprise{IdEnterprise: IdEnterprise}
	members, customErr := getAllEnterpriseMembers(IdEnterprise)
	if (customErr != customError.CustomError{}) {
		// logear el error
		return errorHandler.HandleError(c, &customErr)
	}
	customErr = deleteModelCaller(e)
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	customErr = setIdEnterpriseToOne(members)
	if (customErr != customError.CustomError{}) {
		// logear el error
		return errorHandler.HandleError(c, &customErr)
	}
	switch c.Get("mode") {
	case "table":
		return RenderEnterpriseTable(c)
	case "edit":
		return c.Render("tablePage", fiber.Map{"withEnterpriseTable": true})
	default:
		// log error with deleting mode
		return errorHandler.HandleError(c, &customErr)
	}

}

func EditEnterprise(c *fiber.Ctx) error {
	if customErr := validateFieldsCaller(models.Enterprise{}, c); (customErr != customError.CustomError{}) {
		return errorHandler.HandleError(c, &customErr)
	}
	e, customErr := parserCaller(i.EnterpriseParser{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	IdEnterprise, customErr := getIdModelCaller(e, c)
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	e.IdEnterprise = IdEnterprise
	years, customErr := getPaymentYears(e.IdEnterprise)
	if (customErr != customError.CustomError{}) {
		// logearlo
		return errorHandler.HandleError(c, &customErr)
	}
	e, customErr = updateModelCaller(e)
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	numberOfMembers, customErr := getNumberOfMembers(e.IdEnterprise, "")
	if (customErr != customError.CustomError{}) {
		// logearlo
		return errorHandler.HandleError(c, &customErr)
	}
	createdAt, updatedAt, customErr := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
	if (customErr != customError.CustomError{}) {
		// logearlo en algun lado
		return errorHandler.HandleError(c, &customErr)
	}
	data := fiber.Map{"enterprise": e, "numberOfMembers": numberOfMembers, "mode": "edit", "years": years, "createdAt": createdAt, "updatedAt": updatedAt}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeEnterprise"]
	return c.Render("enterpriseFile", data)
}

func getPaymentYears(idEnterprise int) ([]string, customError.CustomError) {
	result, err := database.DB.Query("SELECT Year FROM PaymentTable WHERE IdEnterprise = ? GROUP BY Year ORDER BY YEAR DESC", idEnterprise)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return nil, customError.QueryError
	}

	var years []string
	var year string
	for result.Next() {
		err = result.Scan(&year)
		if err != nil {
			customError.ScanError.Msg = err.Error()
			return nil, customError.ScanError
		}
		years = append(years, year)
	}
	return years, customError.CustomError{}
}

func RenderEnterpriseTable(c *fiber.Ctx) error {
	// obtengo la currentPage del path
	currentPage := GetPageFromPath(c)
	// calculo la cantidad de resultados
	totalRows, customErr := getTotalRowsCaller(models.Enterprise{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.SendString(`<div class="no-result">No se encontraron empresas<a href="/enterprise/addForm"><button class="add-btn-table">Agregar</button></a></div> `)
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		enterprises, searchKey, customErr := searchModelsCaller(models.Enterprise{}, c, offset)
		if (customErr != customError.CustomError{}) {
			// guardar el err
			return errorHandler.HandleError(c, &customErr)
		}

		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		// creo un map con todas las variables
		data := getFiberMapCaller(models.Enterprise{}, enterprises, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)
		data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
		data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeEnterprise"]
		// renderizo la tabla y le envio el map con las variables
		return c.Render("enterpriseTable", data)
	}
}

func RenderEnterpriseFile(c *fiber.Ctx) error {
	e, customErr := searchOneModelByIdCaller(models.Enterprise{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	numberOfMembers, customErr := getNumberOfMembers(e.IdEnterprise, "")
	if (customErr != customError.CustomError{}) {
		// logear el err
		return errorHandler.HandleError(c, &customErr)
	}
	canDelete := c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
	canWrite := c.Locals("claims").(jwt.MapClaims)["writeEnterprise"]
	if e.IdEnterprise == 1 {
		data := fiber.Map{"enterprise": e, "canDelete": canDelete, "canWrite": canWrite, "numberOfMembers": numberOfMembers, "mode": "edit"}
		return c.Render("withoutEnterprise", data)
	} else {
		createdAt, updatedAt, customErr := formatTimeStamps(e.CreatedAt, e.UpdatedAt)
		if (customErr != customError.CustomError{}) {
			// logear el err
			return errorHandler.HandleError(c, &customErr)
		}
		data := fiber.Map{"enterprise": e, "canDelete": canDelete, "canWrite": canWrite, "numberOfMembers": numberOfMembers, "mode": "edit", "withPaymentTable": false, "createdAt": createdAt, "updatedAt": updatedAt}
		return c.Render("enterpriseFile", data)
	}
}

func getNumberOfMembers(IdEnterprise int, searchKey string) (int, customError.CustomError) {
	var totalRows int
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM MemberTable 
		WHERE IdEnterprise = ? AND (Name LIKE concat('%', ?, '%') OR LastName LIKE concat('%', ?, '%'))`,
		IdEnterprise, searchKey, searchKey)
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		customError.ScanError.Msg = err.Error()
		return 0, customError.ScanError
	}
	return totalRows, customError.CustomError{}
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

	IdEnterprise, customErr := func() (int, customError.CustomError) {
		// c.Get devuelve un valor del header
		if c.Get("mode") == "edit" {
			if IdEnterprise, customErr := getIdModelCaller(models.Enterprise{}, c); (customErr != customError.CustomError{}) {
				// guardar el err
				return 0, customErr
			} else {
				return IdEnterprise, customError.CustomError{}
			}
		} else if c.Get("mode") == "enterpriseMemberTable" {
			IdEnterpriseStr := c.Get("idEnterprise")
			IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
			if err != nil {
				customError.StrConvError.Msg = err.Error()
				return 0, customError.StrConvError
			}
			return IdEnterprise, customError.CustomError{}
		} else {
			return 0, customError.CustomError{}
		}
	}()
	if (customErr != customError.CustomError{}) {
		// logear el error
		return errorHandler.HandleError(c, &customErr)
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
	totalRows, customErr := getNumberOfMembers(IdEnterprise, searchKey)
	if (customErr != customError.CustomError{}) {
		// logear el err
		return errorHandler.HandleError(c, &customErr)
	}

	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.SendString(`<div class="no-result-file">No se encontraron afiliados</div>`)
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		members, customErr := getEnterpriseMembers(IdEnterprise, searchKey, offset)
		if (customErr != customError.CustomError{}) {
			// logear el err
			return errorHandler.HandleError(c, &customErr)
		}

		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		enterpriseId := c.Get("enterpriseId")
		// creo un map con todas las variables
		data := getFiberMapCaller(models.Member{}, members, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)
		data["mode"] = "enterpriseMemberTable"
		data["IdEnterprise"] = IdEnterprise
		data["fromEnterprise"] = true
		data["enterpriseId"] = enterpriseId
		data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
		data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeEnterprise"]
		// renderizo la tabla y le envio el map con las variables
		return c.Render("memberTable", data)
	}
}

func getEnterpriseMembers(IdEnterprise int, searchKey string, offset int) ([]models.Member, customError.CustomError) {
	result, err := database.DB.Query(`
			SELECT * FROM MemberTable 
			WHERE IdEnterprise = ?
			AND (Name LIKE concat('%', ?, '%') OR LastName LIKE concat('%', ?, '%')) LIMIT 10 OFFSET ?`, IdEnterprise, searchKey, searchKey, offset)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return nil, customError.QueryError
	}
	_, members, customErr := models.Member{}.ScanResult(result, false)
	if (customErr != customError.CustomError{}) {
		return nil, customErr
	}
	return members, customError.CustomError{}
}

func getAllEnterpriseMembers(IdEnterprise int) ([]models.Member, customError.CustomError) {
	result, err := database.DB.Query(`
			SELECT * FROM MemberTable WHERE IdEnterprise = ?`, IdEnterprise)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return nil, customError.QueryError
	}
	_, members, customErr := models.Member{}.ScanResult(result, false)
	if (customErr != customError.CustomError{}) {
		return nil, customError.CustomError{}
	}
	return members, customError.CustomError{}
}

func setIdEnterpriseToOne(members []models.Member) customError.CustomError {
	for _, m := range members {
		update, err := database.DB.Query(`
		UPDATE MemberTable SET IdEnterprise = '1' WHERE IdMember = ?`, m.IdMember)
		if err != nil {
			customError.QueryError.Msg = err.Error()
			return customError.QueryError
		}
		update.Close()
		// query.Close()
	}
	return customError.CustomError{}
}


func RenderEnterpriseTableSelect(c *fiber.Ctx) error {
	// calculo la cantidad de resultados
	totalRows, customErr := getTotalRowsCaller(models.Enterprise{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
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
		Name LIKE concat('%', ?, '%') 
		ORDER BY Name ASC`,
			searchKey)
		if err != nil {
			// logear el err
			customError.QueryError.Msg = err.Error()
			return errorHandler.HandleError(c, &customError.QueryError)
		}
		_, enterprises, customErr := models.Enterprise{}.ScanResult(result, false)
		if (customErr != customError.CustomError{}) {
			// guardar el err
			return errorHandler.HandleError(c, &customErr)
		}

		// creo un map con todas las variables
		canDelete := c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
		canWrite := c.Locals("claims").(jwt.MapClaims)["writeEnterprise"]
		// renderizo la tabla y le envio el map con las variables
		return c.Render("enterpriseTableSelect", fiber.Map{"enterprises": enterprises, "canDelete": canDelete, "canWrite": canWrite})
	}
}

func GetAllEnterprisesId(c *fiber.Ctx) error {
	enterprisesId, customErr := models.GetAllEnterprisesIdFromDB()
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	return c.JSON(enterprisesId)
}

func GetAllEnterprisesNumber(c *fiber.Ctx) error {
	enterprisesNumbers, customErr := models.GetAllEnterprisesNumbersFromDB()
	if (customErr != customError.CustomError{}) {
		// guardar el err
		return errorHandler.HandleError(c, &customErr)
	}
	return c.JSON(enterprisesNumbers)
}
