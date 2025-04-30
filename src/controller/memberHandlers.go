package controller

import (
	"fmt"

	er "github.com/LucasBastino/app-sindicato/src/errors"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddMember(c *fiber.Ctx) error {
	m, err := parserCaller(i.MemberParser{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	// Creo un mapa con los errores de validacion y verifico si tiene errores
	if err := validateFieldsCaller(models.Member{}, c); err != nil {
		return er.CheckError(c, err)
	}
	// Si no tiene errores inserto el member en la DB y renderizo el su archivo
	m, err = insertModelCaller(m)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	path := fmt.Sprintf("/member/%d/file", m.IdMember)
	return c.Status(fiber.StatusCreated).Render("redirect", fiber.Map{"path": path})

}

func DeleteMember(c *fiber.Ctx) error {
	// Obtengo el ID desde el path y lo elimino
	IdMember, err := getIdModelCaller(models.Member{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	m := models.Member{IdMember: IdMember}
	err = deleteModelCaller(m)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}

	// check if the member was deleted

	checkDeleted, err := checkDeletedCaller(models.Member{}, IdMember)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	if !checkDeleted {
		return c.Render("deleteUnsuccesfull", fiber.Map{"error": "error eliminando afiliado"})
	} else {

		switch c.Get("mode") {
		case "table":
			return RenderMemberTable(c)
		case "edit":
			return c.Render("redirect", fiber.Map{"path": "/"})
		case "enterpriseMemberTable":
			return RenderEnterpriseMembers(c)
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error with deleting mode"})
		}

	}
}

func EditMember(c *fiber.Ctx) error {
	enterprises, err := getAllModelsCaller(models.Enterprise{})
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	if err := validateFieldsCaller(models.Member{}, c); err != nil {
		return er.CheckError(c, err)
	}
	// Parseo los datos obtenidos del form
	m, err := parserCaller(i.MemberParser{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	enterpriseName, err := getEnterpriseName(m.IdEnterprise)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	IdMember, err := getIdModelCaller(m, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	m.IdMember = IdMember
	// necesito poner esta linea ↑ para que se pueda editar 2 veces seguidas
	role := c.Locals("claims").(jwt.MapClaims)["role"]

	m, err = updateModelCaller(m)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	createdAt, updatedAt, err := formatTimeStamps(m.CreatedAt, m.UpdatedAt)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	// hacer esto esta bien? estoy mostrando datos del nuevo member, no estan sacados de la database.DB
	data := fiber.Map{"member": m, "mode": "edit", "role": role, "enterprises": enterprises, "enterpriseName": enterpriseName, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("memberFile", data)

}

func RenderMemberTable(c *fiber.Ctx) error {
	// Busco todos los members por key y renderizo la tabla de miembros

	// obtengo la currentPage del path
	currentPage := GetPageFromPath(c)

	// calculo la cantidad de resultados
	totalRows, err := getTotalRowsCaller(models.Member{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.SendString(`<div class="no-result">No se encontraron afiliados</div>`)
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		members, searchKey, err := searchModelsCaller(models.Member{}, c, offset)
		if err != nil {
			// guardar el error
			return er.CheckError(c, err)
		}
		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		enterprises, err := getAllModelsCaller(models.Enterprise{})
		if err != nil {
			// guardar el error
			return er.CheckError(c, err)
		}

		// creo un map con todas las variables
		data := getFiberMapCaller(models.Member{}, members, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)
		data["enterprises"] = enterprises
		data["mode"] = "table"
		role := c.Locals("claims").(jwt.MapClaims)["role"]
		data["role"] = role
		// renderizo la tabla y le envio el map con las variables
		return c.Render("memberTable", data)
	}
}

func RenderMemberFile(c *fiber.Ctx) error {
	// Busco el miembro por ID y renderizo su archivo
	enterprises, err := getAllModelsCaller(models.Enterprise{})
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	m, err := searchOneModelByIdCaller(models.Member{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}

	enterpriseName, err := getEnterpriseName(m.IdEnterprise)
	if err != nil {
		return er.CheckError(c, err)
	}
	createdAt, updatedAt, err := formatTimeStamps(m.CreatedAt, m.UpdatedAt)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	data := fiber.Map{"member": m, "mode": "edit", "role": role, "enterprises": enterprises, "enterpriseName": enterpriseName, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("memberFile", data)
}

func RenderAddMemberForm(c *fiber.Ctx) error {
	// le paso un member vacio para que los campos del form aparezcan vacios
	enterprises, err := getAllModelsCaller(models.Enterprise{})
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	data := fiber.Map{"member": models.Member{}, "mode": "add", "enterprises": enterprises}
	return c.Render("memberFile", data)
}

func RenderParentTable(c *fiber.Ctx) error {
	// calculo la cantidad de resultados
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	totalRows, err := getTotalRowsCaller(models.Parent{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	IdMember, err := getIdModelCaller(models.Member{}, c)
	if err != nil {
		// guardar el error
		return er.CheckError(c, err)
	}
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.Render("noResultsParents", fiber.Map{"idMember": IdMember, "role": role})
	} else {
		// Busco los parents asociados a ese member
		parents, _, err := searchModelsCaller(models.Parent{}, c, 0)
		if err != nil {
			// guardar el error
			return er.CheckError(c, err)
		}
		data := fiber.Map{"idMember": IdMember, "role": role, "parents": parents, "mode": "edit"}
		return c.Render("parentTable", data)
	}
}
