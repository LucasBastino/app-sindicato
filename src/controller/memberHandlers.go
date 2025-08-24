package controller

import (
	"fmt"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/LucasBastino/app-sindicato/src/errors/errorHandler"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddMember(c *fiber.Ctx) error {
	m, customErr := parserCaller(i.MemberParser{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	// Creo un mapa con los errores de validacion y verifico si tiene errores
	if customErr := validateFieldsCaller(models.Member{}, c); (customErr != customError.CustomError{}) {
		return errorHandler.HandleError(c, &customErr)
	}
	if !m.Affiliated{
		m.IdEnterprise = 1
	}
	if m.IdEnterprise == 1{
		m.Affiliated = false
	}
	// Si no tiene errores inserto el member en la DB y renderizo el su archivo
	m, customErr = insertModelCaller(m)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	path := fmt.Sprintf("/member/%d/file", m.IdMember)
	return c.Status(fiber.StatusCreated).Render("redirect", fiber.Map{"path": path})

}

func DeleteMember(c *fiber.Ctx) error {
	// Obtengo el ID desde el path y lo elimino
	IdMember, customErr := getIdModelCaller(models.Member{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	m := models.Member{IdMember: IdMember}
	customErr = deleteModelCaller(m)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}

	// check if the member was deleted

	checkDeleted, customErr := checkDeletedCaller(models.Member{}, IdMember)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	if !checkDeleted {
		return c.Render("deleteUnsuccesfull", fiber.Map{"error": "error eliminando afiliado"})
	} else {

		switch c.Get("mode") {
		case "table":
			return RenderMemberTable(c)
		case "edit":
			return c.Render("redirect", fiber.Map{"path": "/member/list"})
		case "enterpriseMemberTable":
			return RenderEnterpriseMembers(c)
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error with deleting mode"})
		}

	}
}

func EditMember(c *fiber.Ctx) error {
	enterprises, customErr := getAllModelsCaller(models.Enterprise{})
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	if customErr := validateFieldsCaller(models.Member{}, c); (customErr != customError.CustomError{}) {
		return errorHandler.HandleError(c, &customErr)
	}
	// Parseo los datos obtenidos del form
	m, customErr := parserCaller(i.MemberParser{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	IdMember, customErr := getIdModelCaller(m, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	m.IdMember = IdMember
	// necesito poner esta linea ↑ para que se pueda editar 2 veces seguidas
	
	if !m.Affiliated{
		m.IdEnterprise = 1
	}
	if m.IdEnterprise == 1{
		m.Affiliated = false
	}
	m, customErr = updateModelCaller(m)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	enterpriseName, customErr := getEnterpriseName(m.IdEnterprise)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	createdAt, updatedAt, customErr := formatTimeStamps(m.CreatedAt, m.UpdatedAt)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	// hacer esto esta bien? estoy mostrando datos del nuevo member, no estan sacados de la database.DB
	data := fiber.Map{"member": m, "mode": "edit", "enterprises": enterprises, "enterpriseName": enterpriseName, "createdAt": createdAt, "updatedAt": updatedAt}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteMember"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeMember"]
	return c.Render("memberFile", data)

}

func RenderMemberTable(c *fiber.Ctx) error {
	// Busco todos los members por key y renderizo la tabla de miembros

	// obtengo la currentPage del path
	currentPage := GetPageFromPath(c)

	// calculo la cantidad de resultados
	totalRows, customErr := getTotalRowsCaller(models.Member{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.SendString(`<div class="no-result">No se encontraron afiliados<a href="/member/addForm"><button class="add-btn-table">Agregar</button></a></div> `)
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		members, searchKey, customErr := searchModelsCaller(models.Member{}, c, offset)
		if (customErr != customError.CustomError{}) {
			// guardar el error
			return errorHandler.HandleError(c, &customErr)
		}
		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		enterprises, customErr := getAllModelsCaller(models.Enterprise{})
		if (customErr != customError.CustomError{}) {
			// guardar el error
			return errorHandler.HandleError(c, &customErr)
		}

		// creo un map con todas las variables
		data := getFiberMapCaller(models.Member{}, members, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)
		data["enterprises"] = enterprises
		data["mode"] = "table"
		data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteMember"]
		data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeMember"]
		// renderizo la tabla y le envio el map con las variables
		return c.Render("memberTable", data)
	}
}

func RenderMemberFile(c *fiber.Ctx) error {
	// Busco el miembro por ID y renderizo su archivo
	enterprises, customErr := getAllModelsCaller(models.Enterprise{})
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	m, customErr := searchOneModelByIdCaller(models.Member{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}

	enterpriseName, customErr := getEnterpriseName(m.IdEnterprise)
	if (customErr != customError.CustomError{}) {
		return errorHandler.HandleError(c, &customErr)
	}
	createdAt, updatedAt, customErr := formatTimeStamps(m.CreatedAt, m.UpdatedAt)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	data := fiber.Map{"member": m, "mode": "edit", "enterprises": enterprises, "enterpriseName": enterpriseName, "createdAt": createdAt, "updatedAt": updatedAt}
	data["canDelete"] = c.Locals("claims").(jwt.MapClaims)["deleteMember"]
	data["canWrite"] = c.Locals("claims").(jwt.MapClaims)["writeMember"]
	return c.Render("memberFile", data)
}

func RenderAddMemberForm(c *fiber.Ctx) error {
	// le paso un member vacio para que los campos del form aparezcan vacios
	enterprises, customErr := getAllModelsCaller(models.Enterprise{})
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}

	fromEnterprise := c.Get("fromEnterprise")
	enterpriseId := 0
	enterpriseName := ""
	if fromEnterprise == "true" {

		enterpriseIdStr := c.Get("enterpriseId")
		if enterpriseIdStr != "" {
			enterpriseId2, err := strconv.Atoi(enterpriseIdStr)
			if err != nil {
				customError.StrConvError.Msg = err.Error()
				return errorHandler.HandleError(c, &customError.StrConvError)
			}
			enterpriseId = enterpriseId2
		}
		enterpriseName, customErr = getEnterpriseName(enterpriseId)
		if (customErr != customError.CustomError{}) {
			// guardar el error
			return errorHandler.HandleError(c, &customErr)
		}
	}
	data := fiber.Map{"member": models.Member{IdEnterprise: enterpriseId}, "mode": "add", "enterprises": enterprises, "enterpriseName": enterpriseName}
	return c.Render("memberFile", data)
}

func RenderParentTable(c *fiber.Ctx) error {
	canDelete := c.Locals("claims").(jwt.MapClaims)["deleteParent"]
	canWrite := c.Locals("claims").(jwt.MapClaims)["writeParent"]
	// calculo la cantidad de resultados
	totalRows, customErr := getTotalRowsCaller(models.Parent{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	IdMember, customErr := getIdModelCaller(models.Member{}, c)
	if (customErr != customError.CustomError{}) {
		// guardar el error
		return errorHandler.HandleError(c, &customErr)
	}
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.Render("noResultsParents", fiber.Map{"idMember": IdMember, "canDelete": canDelete, "canWrite": canWrite})
	} else {
		// Busco los parents asociados a ese member
		parents, _, customErr := searchModelsCaller(models.Parent{}, c, 0)
		if (customErr != customError.CustomError{}) {
			// guardar el error
			return errorHandler.HandleError(c, &customErr)
		}
		data := fiber.Map{"idMember": IdMember, "canDelete": canDelete, "canWrite": canWrite, "parents": parents, "mode": "edit"}
		return c.Render("parentTable", data)
	}
}
