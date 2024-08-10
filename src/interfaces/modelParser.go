package interfaces

import (
	"fmt"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

type ModelParser[M models.TypeModel] interface {
	ParseModel(*fiber.Ctx) M
}

type MemberParser struct{}

func (m MemberParser) ParseModel(c *fiber.Ctx) models.Member {
	member := models.Member{}
	member.Name = c.FormValue("name")
	member.DNI = c.FormValue("dni")
	return member
}

type ParentParser struct{}

func (p ParentParser) ParseModel(c *fiber.Ctx) models.Parent {
	parent := models.Parent{}
	parent.Name = c.FormValue("name")
	parent.Rel = c.FormValue("rel")
	IdMemberStr := c.FormValue("idmember")
	IdMember, err := strconv.Atoi(IdMemberStr)
	if err != nil {
		fmt.Println("error converting IdMemberStr to int")
		panic(err)
	}
	parent.IdMember = IdMember

	return parent
}

type EnterpriseParser struct{}

func (p EnterpriseParser) ParseModel(c *fiber.Ctx) models.Enterprise {
	enterprise := models.Enterprise{}
	enterprise.Name = c.FormValue("name")
	enterprise.Address = c.FormValue("address")
	return enterprise
}
