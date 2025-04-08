package controller

import (
	"fmt"

	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddMember(c *fiber.Ctx) error {
	// Creo un mapa con los errores de validacion
	errorMap := validateFieldsCaller(models.Member{}, c)
	m := parserCaller(i.MemberParser{}, c)
	// Verifico si el mapa tiene errores
	if len(errorMap) > 0 {
		fmt.Println(errorMap)
		return c.Status(fiber.StatusBadRequest).JSON(errorMap)
	} else {
		// Si no tiene errores inserto el member en la DB y renderizo el su archivo
		m = insertModelCaller(m)
		path := fmt.Sprintf("/member/%d/file", m.IdMember)
		return c.Status(fiber.StatusCreated).Render("redirect", fiber.Map{"path": path})
	}
}

func DeleteMember(c *fiber.Ctx) error {
	// Obtengo el ID desde el path y lo elimino
	IdMember := getIdModelCaller(models.Member{}, c)
	m := models.Member{IdMember: IdMember}
	deleteModelCaller(m)

	// check if the member was deleted

	checkDeleted := checkDeletedCaller(models.Member{}, IdMember)
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
	enterprises := getAllModelsCaller(models.Enterprise{})
	errorMap := validateFieldsCaller(models.Member{}, c)
	// Parseo los datos obtenidos del form
	m := parserCaller(i.MemberParser{}, c)
	enterpriseName, err := getEnterpriseName(m.IdEnterprise)
	if err != nil {
		// ver esto
		c.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": err})
	}
	IdMember := getIdModelCaller(m, c)
	m.IdMember = IdMember
	// necesito poner esta linea ↑ para que se pueda editar 2 veces seguidas
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	if len(errorMap) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errorMap)
	} else {
		m = updateModelCaller(m)
		createdAt, updatedAt := formatTimeStamps(m.CreatedAt, m.UpdatedAt)
		// hacer esto esta bien? estoy mostrando datos del nuevo member, no estan sacados de la database.DB
		data := fiber.Map{"member": m, "mode": "edit", "role": role, "enterprises": enterprises, "enterpriseName": enterpriseName, "createdAt": createdAt, "updatedAt": updatedAt}
		return c.Render("memberFile", data)

	}
}

func RenderMemberTable(c *fiber.Ctx) error {

	// Busco todos los members por key y renderizo la tabla de miembros

	// obtengo la currentPage del path
	currentPage := GetPageFromPath(c)

	// calculo la cantidad de resultados
	totalRows := getTotalRowsCaller(models.Member{}, c)
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.SendString(`<div class="no-result">No se encontraron afiliados</div>`)
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		members, searchKey := searchModelsCaller(models.Member{}, c, offset)

		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		enterprises := getAllModelsCaller(models.Enterprise{})

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
	enterprises := getAllModelsCaller(models.Enterprise{})
	m := searchOneModelByIdCaller(models.Member{}, c)
	enterpriseName, err := getEnterpriseName(m.IdEnterprise)
	if err != nil {
		// ver esto
		c.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": err})
	}
	createdAt, updatedAt := formatTimeStamps(m.CreatedAt, m.UpdatedAt)
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	data := fiber.Map{"member": m, "mode": "edit", "role": role, "enterprises": enterprises, "enterpriseName": enterpriseName, "createdAt": createdAt, "updatedAt": updatedAt}
	return c.Render("memberFile", data)
}

func RenderAddMemberForm(c *fiber.Ctx) error {
	// le paso un member vacio para que los campos del form aparezcan vacios
	enterprises := getAllModelsCaller(models.Enterprise{})
	data := fiber.Map{"member": models.Member{}, "mode": "add", "enterprises": enterprises}
	return c.Render("memberFile", data)
}

func RenderParentTable(c *fiber.Ctx) error {
	// calculo la cantidad de resultados
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	totalRows := getTotalRowsCaller(models.Parent{}, c)
	IdMember := getIdModelCaller(models.Member{}, c)
	if totalRows == 0 {
		// si no hay resultados renderizar esto
		return c.Render("noResultsParents", fiber.Map{"idMember": IdMember, "role": role})
	} else {
		// Busco los parents asociados a ese member
		parents, _ := searchModelsCaller(models.Parent{}, c, 0)
		data := fiber.Map{"idMember": IdMember, "role": role, "parents": parents, "mode": "edit"}
		return c.Render("parentTable", data)
	}
}
