package interfaces

import (
	"strconv"

	er "github.com/LucasBastino/app-sindicato/src/errors"

	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

type ModelParser[M models.TypeModel] interface {
	ParseModel(*fiber.Ctx) (M, error)
}

type MemberParser struct{}

func (parser MemberParser) ParseModel(c *fiber.Ctx) (models.Member, error) {
	m := models.Member{}
	m.Name = c.FormValue("name")
	m.LastName = c.FormValue("last-name")
	m.DNI = c.FormValue("dni")
	m.Birthday = c.FormValue("birthday")
	m.Gender = c.FormValue("gender")
	m.MaritalStatus = c.FormValue("marital-status")
	m.Phone = c.FormValue("phone")
	m.Email = c.FormValue("email")
	m.Address = c.FormValue("address")
	m.PostalCode = c.FormValue("postal-code")
	m.District = c.FormValue("district")
	m.MemberNumber = c.FormValue("member-number")
	m.CUIL = c.FormValue("cuil")
	IdEnterpriseStr := c.FormValue("id-enterprise")
	if IdEnterpriseStr == "" {
		m.IdEnterprise = 0
		// este valor igualmente no se usa
		// es solamente para que no aparezca un error
	} else {
		IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
		if err != nil {
			er.StrConvError.Msg = err.Error()
			return models.Member{}, er.StrConvError
		}
		m.IdEnterprise = IdEnterprise
	}
	m.Category = c.FormValue("category")
	m.EntryDate = c.FormValue("entry-date")
	return m, nil
}

type ParentParser struct{}

func (parser ParentParser) ParseModel(c *fiber.Ctx) (models.Parent, error) {
	p := models.Parent{}
	p.Name = c.FormValue("name")
	p.LastName = c.FormValue("last-name")
	p.Rel = c.FormValue("rel")
	p.Gender = c.FormValue("gender")
	p.Birthday = c.FormValue("birthday")
	p.CUIL = c.FormValue("cuil")
	IdMemberStr := c.FormValue("id-member")
	IdMember, err := strconv.Atoi(IdMemberStr)
	if err != nil {
		er.StrConvError.Msg = err.Error()
		return models.Parent{}, er.StrConvError
	}
	p.IdMember = IdMember

	return p, nil
}

type EnterpriseParser struct{}

func (parser EnterpriseParser) ParseModel(c *fiber.Ctx) (models.Enterprise, error) {
	e := models.Enterprise{}
	e.Name = c.FormValue("name")
	e.EnterpriseNumber = c.FormValue("enterprise-number")
	e.Address = c.FormValue("address")
	e.CUIT = c.FormValue("cuit")
	e.District = c.FormValue("district")
	e.PostalCode = c.FormValue("postal-code")
	e.Phone = c.FormValue("phone")
	return e, nil
}

type PaymentParser struct{}

func (parser PaymentParser) ParseModel(c *fiber.Ctx) (models.Payment, error) {
	p := models.Payment{}
	p.Month = c.FormValue("month")
	p.Year = c.FormValue("year")
	p.Status = c.FormValue("status")
	AmountStr := c.FormValue("amount")
	if AmountStr != "" {
		Amount, err := strconv.Atoi(AmountStr)
		if err != nil {
			er.StrConvError.Msg = err.Error()
			return models.Payment{}, er.StrConvError
		}
		p.Amount = Amount
	}
	p.PaymentDate = c.FormValue("payment-date")
	p.Commentary = c.FormValue("commentary")
	IdEnterpriseStr := c.FormValue("id-enterprise")
	IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
	if err != nil {
		er.StrConvError.Msg = err.Error()
		return models.Payment{}, err
	}
	p.IdEnterprise = IdEnterprise

	return p, nil
}
