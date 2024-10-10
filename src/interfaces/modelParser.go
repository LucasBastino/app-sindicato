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
	member.LastName = c.FormValue("last-name")
	member.DNI = c.FormValue("dni")
	member.Birthday = c.FormValue("birthday")
	member.Gender = c.FormValue("gender")
	member.MaritalStatus = c.FormValue("marital-status")
	member.Phone = c.FormValue("phone")
	member.Email = c.FormValue("email")
	member.Address = c.FormValue("address")
	member.PostalCode = c.FormValue("postal-code")
	member.District = c.FormValue("district")
	member.MemberNumber = c.FormValue("member-number")
	member.CUIL = c.FormValue("cuil")
	IdEnterpriseStr := c.FormValue("id-enterprise")
	if IdEnterpriseStr == "" {
		member.IdEnterprise = 0
		// este valor igualmente no se usa
		// es solamente para que no aparezca un error
	} else {
		IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
		if err != nil {
			fmt.Println(err)
		}
		member.IdEnterprise = IdEnterprise
	}
	member.Category = c.FormValue("category")
	member.EntryDate = c.FormValue("entry-date")
	return member
}

type ParentParser struct{}

func (p ParentParser) ParseModel(c *fiber.Ctx) models.Parent {
	parent := models.Parent{}
	parent.Name = c.FormValue("name")
	parent.LastName = c.FormValue("last-name")
	parent.Rel = c.FormValue("rel")
	parent.Gender = c.FormValue("gender")
	parent.Birthday = c.FormValue("birthday")
	parent.CUIL = c.FormValue("cuil")
	IdMemberStr := c.FormValue("id-member")
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
	enterprise.EnterpriseNumber = c.FormValue("enterprise-number")
	enterprise.Address = c.FormValue("address")
	enterprise.CUIT = c.FormValue("cuit")
	enterprise.District = c.FormValue("district")
	enterprise.PostalCode = c.FormValue("postal-code")
	enterprise.Phone = c.FormValue("phone")
	return enterprise
}
