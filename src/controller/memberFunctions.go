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
	// errorMap := validateFieldsCaller(member, c)
	parser := i.MemberParser{}
	member = parserCaller(parser, c)
	// if len(errorMap) > 0 {
	// 	templateData := createTemplateDataCaller(member, member, nil, "src/views/forms/createMemberForm.html", errorMap)
	// 	return RenderHTML(c, templateData)
	// } else {
	insertModelCaller(member)
	// templateData := createTemplateDataCaller(member, member, nil, "src/views/files/memberFile.html", errorMap)
	return RenderHTML(c, fiber.Map{})

}

func DeleteMember(c *fiber.Ctx) error {
	IdMember := getIdModelCaller(member, c)
	member.IdMember = IdMember
	deleteModelCaller(member)
	return RenderMemberTable(c)
}

func EditMember(c *fiber.Ctx) error {
	member = parserCaller(memberParser, c)
	IdMember := getIdModelCaller(member, c)
	member.IdMember = IdMember
	// necesito poner esta linea ↑ para que se pueda editar 2 veces seguidas
	editModelCaller(member)
	// hacer esto esta bien? estoy mostrando datos del nuevo member, no estan sacados de la database.DB
	// templateData := createTemplateDataCaller(member, member, nil, "src/views/files/memberFile.html", nil)
	return RenderHTML(c, fiber.Map{})

	// no puedo hacer esto ↓ porque estoy en POST, no puedo redireccionar
	// http.Redirect(c , c, "/index", http.StatusSeeOther) // con este status me anda, con otros de 300 no
}

func RenderMemberTable(c *fiber.Ctx) error {
	members := searchModelsCaller(member, c)
	data := fiber.Map{"model": models.Member{}, "members": members, "template": "memberTable"}
	return RenderHTML(c, data)
}

func RenderMemberFile(c *fiber.Ctx) error {
	// member := searchOneModelByIdCaller(member, c)
	// templateData := createTemplateDataCaller(member, member, nil, "src/views/files/memberFile.html", nil)
	return RenderHTML(c, fiber.Map{})
}

func RenderCreateMemberForm(c *fiber.Ctx) error {
	// le paso un member vacio para que los campos del form aparezcan vacios
	// templateData := createTemplateDataCaller(member, models.Member{}, nil, "src/views/forms/createMemberForm.html", nil)
	return RenderHTML(c, fiber.Map{})
}

func RenderMemberParents(c *fiber.Ctx) error {
	IdMember := getIdModelCaller(member, c)
	result, err := database.DB.Query(fmt.Sprintf("SELECT Name, cel, IdParent, IdMember FROM ParentTable WHERE IdMember = '%d'", IdMember))
	if err != nil {
		fmt.Println("error searching parents from database.db")
		panic(err)
	}

	// hacer un metodo para scan
	var parent models.Parent
	var parents []models.Parent
	for result.Next() {
		err = result.Scan(&parent.Name, &parent.Rel, &parent.IdParent, &parent.IdMember)
		if err != nil {
			fmt.Println("error scanning parent")
			panic(err)
		}
		parents = append(parents, parent)
	}

	// templateData := createTemplateDataCaller(parent, parent, parents, "src/views/tables/memberParentTable.html", nil)
	return RenderHTML(c, fiber.Map{})

}
