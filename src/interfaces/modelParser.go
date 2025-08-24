package interfaces

import (
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

type ModelParser[M models.TypeModel] interface {
	ParseModel(*fiber.Ctx) (M, customError.CustomError)
}

type MemberParser struct{}

func (parser MemberParser) ParseModel(c *fiber.Ctx) (models.Member, customError.CustomError) {
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
	affiliated, err := strconv.ParseBool(c.FormValue("affiliated"))
	if err != nil {
		customError.StrConvError.Msg = err.Error()
		return models.Member{}, customError.StrConvError
	}
	m.Affiliated = affiliated
	m.CUIL = c.FormValue("cuil")
	IdEnterpriseStr := c.FormValue("id-enterprise")
	if IdEnterpriseStr == "" {
		m.IdEnterprise = 0
		// este valor igualmente no se usa
		// es solamente para que no aparezca un error
	} else {
		IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
		if err != nil {
			customError.StrConvError.Msg = err.Error()
			return models.Member{},  customError.StrConvError
		}
		m.IdEnterprise = IdEnterprise
	}
	m.Category = c.FormValue("category")
	m.EntryDate = c.FormValue("entry-date")
	m.Observations = c.FormValue("observations")
	return m, customError.CustomError{}
}

type ParentParser struct{}

func (parser ParentParser) ParseModel(c *fiber.Ctx) (models.Parent, customError.CustomError) {
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
		customError.StrConvError.Msg = err.Error()
		return models.Parent{}, customError.StrConvError
	}
	p.IdMember = IdMember

	return p, customError.CustomError{}
}

type EnterpriseParser struct{}

func (parser EnterpriseParser) ParseModel(c *fiber.Ctx) (models.Enterprise, customError.CustomError) {
	e := models.Enterprise{}
	e.Name = c.FormValue("name")
	e.EnterpriseNumber = c.FormValue("enterprise-number")
	e.Address = c.FormValue("address")
	e.CUIT = c.FormValue("cuit")
	e.District = c.FormValue("district")
	e.PostalCode = c.FormValue("postal-code")
	e.Phone = c.FormValue("phone")
	e.Contact = c.FormValue("contact")
	e.Observations = c.FormValue("observations")
	return e, customError.CustomError{}
}

type PaymentParser struct{}

func (parser PaymentParser) ParseModel(c *fiber.Ctx) (models.Payment, customError.CustomError) {
	p := models.Payment{}
	p.Month = c.FormValue("month")
	p.Year = c.FormValue("year")
	status, err := strconv.ParseBool(c.FormValue("status"))
	if err != nil {
		customError.InternalServerError.Msg = err.Error()
		return models.Payment{}, customError.InternalServerError
	}
	p.Status = status
	AmountStr := c.FormValue("amount")
	if AmountStr != "" {
		Amount, err := strconv.Atoi(AmountStr)
		if err != nil {
			customError.StrConvError.Msg = err.Error()
			return models.Payment{}, customError.StrConvError
		}
		p.Amount = Amount
	}
	p.PaymentDate = c.FormValue("payment-date")
	p.Observations = c.FormValue("observations")
	IdEnterpriseStr := c.FormValue("id-enterprise")
	IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
	if err != nil {
		customError.StrConvError.Msg = err.Error()
		return models.Payment{}, customError.StrConvError
	}
	p.IdEnterprise = IdEnterprise

	return p, customError.CustomError{}
}
