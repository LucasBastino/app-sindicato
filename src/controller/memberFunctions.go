package controller

import (
	"fmt"

	"github.com/LucasBastino/app-sindicato/src/database"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateMember(c *fiber.Ctx) error {
	// Creo un mapa con los errores de validacion
	enterprises := getAllModelsCaller(models.Enterprise{})
	errorMap := validateFieldsCaller(models.Member{}, c)
	m := parserCaller(i.MemberParser{}, c)

	// Verifico si el mapa tiene errores
	if len(errorMap) > 0 {
		// Si tiene errores renderizo nuevamente el form
		data := fiber.Map{"member": m, "errorMap": errorMap, "enterprises": enterprises}
		return c.Render("createMemberForm", data)

	} else {
		// Si no tiene errores inserto el member en la DB y renderizo el su archivo
		m = insertModelCaller(m)
		data := fiber.Map{"member": m, "enterprises": enterprises}
		return c.Render("memberFile", data)
	}
}

func DeleteMember(c *fiber.Ctx) error {
	// Obtengo el ID desde el path y lo elimino
	IdMember := getIdModelCaller(models.Member{}, c)
	m := models.Member{IdMember: IdMember}
	deleteModelCaller(m)
	return RenderMemberTable(c)
}

func EditMember(c *fiber.Ctx) error {
	enterprises := getAllModelsCaller(models.Enterprise{})
	errorMap := validateFieldsCaller(models.Member{}, c)
	m := parserCaller(i.MemberParser{}, c)
	IdMember := getIdModelCaller(m, c)
	// Parseo los datos obtenidos del form
	// necesito poner esta linea ↑ para que se pueda editar 2 veces seguidas
	m.IdMember = IdMember
	if len(errorMap) > 0 {
		data := fiber.Map{"member": m, "enterprises": enterprises, "errorMap": errorMap}
		return c.Render("memberFile", data)
	} else {
		editModelCaller(m)
		// hacer esto esta bien? estoy mostrando datos del nuevo member, no estan sacados de la database.DB
		data := fiber.Map{"member": m, "enterprises": enterprises}
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
		return c.Render("searchWithNoResults", fiber.Map{})
	} else {
		// si hay resultados...

		// calcular totalPages
		totalPages, offset, someBefore, someAfter := GetPaginationData(currentPage, totalRows)

		// busco los miembros y devuelvo el searchKey para usarlo nuevamente en la paginacion
		members, searchKey := searchModelsCaller(models.Member{}, c, offset)

		// hago un array para poder recorrerlo y crear botones cuando hay menos de 10 paginas en el template
		totalPagesArray := GetTotalPagesArray(totalPages)

		// creo un map con todas las variables
		data := getFiberMapCaller(models.Member{}, members, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)

		// renderizo la tabla y le envio el map con las variables
		return c.Render("memberTable", data)
	}
}

func RenderMemberFile(c *fiber.Ctx) error {
	// Busco el miembro por ID y renderizo su archivo
	enterprises := getAllModelsCaller(models.Enterprise{})
	m := searchOneModelByIdCaller(models.Member{}, c)
	data := fiber.Map{"member": m, "enterprises": enterprises}
	return c.Render("memberFile", data)
}

func RenderCreateMemberForm(c *fiber.Ctx) error {
	// le paso un member vacio para que los campos del form aparezcan vacios
	enterprises := getAllModelsCaller(models.Enterprise{})
	data := fiber.Map{"member": models.Member{}, "enterprises": enterprises}
	return c.Render("createMemberForm", data)
}

func RenderMemberParents(c *fiber.Ctx) error {
	// Obtengo el ID del member mediante el path
	IdMember := getIdModelCaller(models.Member{}, c)

	// Busco los parents asociados a ese member
	result, err := database.DB.Query(fmt.Sprintf("SELECT Name, Rel, IdParent, IdMember FROM ParentTable WHERE IdMember = '%d'", IdMember))
	if err != nil {
		fmt.Println("error searching parents from database.db")
		panic(err)
	}

	var p models.Parent
	var parents []models.Parent
	for result.Next() {
		// Scanneo los datos y los agrego a un slice
		err = result.Scan(&p.Name, &p.Rel, &p.IdParent, &p.IdMember)
		if err != nil {
			fmt.Println("error scanning parent")
			panic(err)
		}
		parents = append(parents, p)
	}

	data := fiber.Map{"parents": parents}
	return c.Render("memberParentTable", data)

}

func Prueba(c *fiber.Ctx) error {
	return c.Render("prueba", fiber.Map{})
}
