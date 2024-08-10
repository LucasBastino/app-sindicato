package controller

import (
	"fmt"

	"github.com/LucasBastino/app-sindicato/src/database"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

var (
	member       models.Member
	memberParser i.MemberParser
)

func CreateMember(c *fiber.Ctx) error {
	// Creo un mapa con los errores de validacion
	errorMap := validateFieldsCaller(member, c)
	parser := i.MemberParser{}
	member = parserCaller(parser, c)

	// Verifico si el mapa tiene errores
	if len(errorMap) > 0 {
		// Si tiene errores renderizo nuevamente el form
		data := fiber.Map{"model": member, "member": member, "errorMap": errorMap}
		return c.Render("createMemberForm", data)

	} else {
		// Si no tiene errores inserto el member en la DB y renderizo el su archivo
		insertModelCaller(member)
		data := fiber.Map{"model": member, "member": member}
		return c.Render("memberFile", data)
	}
}

func DeleteMember(c *fiber.Ctx) error {
	// Obtengo el ID desde el path y lo elimino
	IdMember := getIdModelCaller(member, c)
	member.IdMember = IdMember
	deleteModelCaller(member)
	return RenderMemberTable(c)
}

func EditMember(c *fiber.Ctx) error {
	// falta hacer la validacion
	// Parseo los datos obtenidos del form
	member = parserCaller(memberParser, c)
	IdMember := getIdModelCaller(member, c)
	member.IdMember = IdMember
	// necesito poner esta linea ↑ para que se pueda editar 2 veces seguidas
	editModelCaller(member)
	// hacer esto esta bien? estoy mostrando datos del nuevo member, no estan sacados de la database.DB

	data := fiber.Map{"model": member, "member": member}
	return c.Render("memberFile", data)
}

func RenderMemberTable(c *fiber.Ctx) error {
	// Busco todos los members y renderizo la tabla de miembros
	members := searchModelsCaller(member, c)
	data := fiber.Map{"model": member, "members": members}
	return c.Render("memberTable", data)
}

func RenderMemberFile(c *fiber.Ctx) error {
	// Busco el miembro por ID y renderizo su archivo
	member := searchOneModelByIdCaller(member, c)
	data := fiber.Map{"model": member, "member": member}
	return c.Render("memberFile", data)
}

func RenderCreateMemberForm(c *fiber.Ctx) error {
	// Renderizo el form para crear miembro
	// le paso un member vacio para que los campos del form aparezcan vacios
	data := fiber.Map{"model": member, "member": models.Member{}}
	return c.Render("createMemberForm", data)
}

func RenderMemberParents(c *fiber.Ctx) error {
	// Obtengo el ID del member mediante el path
	IdMember := getIdModelCaller(member, c)

	// Busco los parents asociados a ese member
	result, err := database.DB.Query(fmt.Sprintf("SELECT Name, cel, IdParent, IdMember FROM ParentTable WHERE IdMember = '%d'", IdMember))
	if err != nil {
		fmt.Println("error searching parents from database.db")
		panic(err)
	}

	var parent models.Parent
	var parents []models.Parent
	for result.Next() {
		// Scanneo los datos y los agrego a un slice
		err = result.Scan(&parent.Name, &parent.Rel, &parent.IdParent, &parent.IdMember)
		if err != nil {
			fmt.Println("error scanning parent")
			panic(err)
		}
		parents = append(parents, parent)
	}

	data := fiber.Map{"model": parent, "parents": parents}
	return c.Render("memberParentTable", data)

}
