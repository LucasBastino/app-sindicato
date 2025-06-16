package router

import (
	c "github.com/LucasBastino/app-sindicato/src/controller"
	creators "github.com/LucasBastino/app-sindicato/src/creators"
	er "github.com/LucasBastino/app-sindicato/src/errors"
	l "github.com/LucasBastino/app-sindicato/src/login"
	pe "github.com/LucasBastino/app-sindicato/src/permissions"
	"github.com/gofiber/fiber/v2"
)

// FALTA HACER LOS GROUPS
func RegisterRoutes(app *fiber.App) {
	app.Get("/", l.VerifyToken, c.RenderIndex)
	// borrar esto despues ↓ --------------------
	app.Get("/register", c.RenderRegisterUserForm)
	// app.Get("/register", l.VerifyToken, pe.VerifyAdmin, c.RenderRegisterUserForm)
	// app.Post("/register", l.VerifyToken, pe.VerifyAdmin, l.RegisterUser)
	app.Post("/register", l.RegisterUser)
	// --------------------------
	app.Get("/login", c.RenderLogin)
	app.Post("/login", l.LoginUser)
	app.Get("/logout", l.LogoutUser)
	app.Get("/expiredSession", l.RenderExpiredSession)
	app.Get("/insufficientPermissions", l.RenderInsufficientPermissions)

	m := app.Group("/member")
	p := app.Group("/parent")
	e := app.Group("/enterprise")
	py := app.Group("/payment")

	m.Get("/list", l.VerifyToken, c.RenderMemberList)
	m.Get("/renderTable", l.VerifyToken, c.RenderMemberTable)
	m.Get("/renderTable/:Page", l.VerifyToken, c.RenderMemberTable)
	m.Get("/addForm", l.VerifyToken, pe.VerifyWriteMember, c.RenderAddMemberForm)
	m.Post("/add", l.VerifyToken, pe.VerifyWriteMember, c.AddMember)
	m.Get("/:IdMember/file", l.VerifyToken, c.RenderMemberFile)
	m.Get("/:IdMember/parentTable", l.VerifyToken, c.RenderParentTable)
	m.Put("/:IdMember/edit", l.VerifyToken, pe.VerifyWriteMember, c.EditMember)
	m.Delete("/:IdMember/delete", l.VerifyToken, pe.VerifyDeleteMember, c.DeleteMember)

	p.Get("/:IdMember/addForm", l.VerifyToken, pe.VerifyWriteParent, c.RenderAddParentForm)
	p.Post("/:IdMember/add", l.VerifyToken, pe.VerifyWriteParent, c.AddParent)
	p.Get("/:IdMember/:IdParent/file", l.VerifyToken, c.RenderParentFile)
	p.Put("/:IdMember/:IdParent/edit", l.VerifyToken, pe.VerifyWriteParent, c.EditParent)
	p.Delete("/:IdMember/:IdParent/delete", l.VerifyToken, pe.VerifyDeleteParent, c.DeleteParent)

	e.Get("/list", l.VerifyToken, c.RenderEnterpriseList)
	e.Get("/renderTable", l.VerifyToken, c.RenderEnterpriseTable)
	e.Get("/renderTable/:Page", l.VerifyToken, c.RenderEnterpriseTable)
	e.Get("/renderTableSelect", l.VerifyToken, c.RenderEnterpriseTableSelect)
	e.Get("/addForm", l.VerifyToken, pe.VerifyWriteEnterprise, c.RenderAddEnterpriseForm)
	e.Post("/add", l.VerifyToken, pe.VerifyWriteEnterprise, c.AddEnterprise)
	e.Get("/:IdEnterprise/file", l.VerifyToken, c.RenderEnterpriseFile)
	e.Get("/:IdEnterprise/memberTable", l.VerifyToken, c.RenderEnterpriseMembers)
	e.Delete("/:IdEnterprise/delete", l.VerifyToken, pe.VerifyDeleteEnterprise, c.DeleteEnterprise)
	e.Put("/:IdEnterprise/edit", l.VerifyToken, pe.VerifyWriteEnterprise, c.EditEnterprise)
	e.Get("/:IdEnterprise/paymentTable/:Year", l.VerifyToken, c.RenderEnterprisePaymentsTable)
	e.Get("/getAllEnterprisesId", l.VerifyToken, c.GetAllEnterprisesId)
	e.Get("/getAllEnterprisesNumber", l.VerifyToken, c.GetAllEnterprisesNumber)

	py.Get("/:IdEnterprise/addForm", l.VerifyToken, pe.VerifyWritePayment, c.RenderAddPaymentForm)
	py.Post("/:IdEnterprise/add", l.VerifyToken, pe.VerifyWritePayment, c.AddPayment)
	py.Get("/:IdEnterprise/:IdPayment/file", l.VerifyToken, c.RenderPaymentFile)
	py.Put("/:IdEnterprise/:IdPayment/edit", l.VerifyToken, pe.VerifyWritePayment, c.EditPayment)
	py.Delete("/:IdEnterprise/:IdPayment/delete", l.VerifyToken, pe.VerifyDeletePayment, c.DeletePayment)
	py.Get("/:IdEnterprise/paymentTable", l.VerifyToken, c.RenderEnterprisePaymentsTable)

	app.Get("/error", er.RenderError)

	app.Get("/test", c.TestNull)
	app.Get("/createMembers", creators.CreateMembers)
	app.Get("/createParents", creators.CreateParents)
	app.Get("/createEnterprises", creators.CreateEnterprises)
	app.Get("/createPayments", creators.CreatePayments)
	app.Get("/renderElectoralMemberList", l.VerifyToken, c.RenderElectoralMemberList)
	app.Get("/pruebaEmpresas", c.RenderPruebaEmpresas)
	app.Get("/backupDB", l.VerifyToken, c.BackupDB)
}
