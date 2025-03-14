package router

import (
	c "github.com/LucasBastino/app-sindicato/src/controller"
	creators "github.com/LucasBastino/app-sindicato/src/creators"
	l "github.com/LucasBastino/app-sindicato/src/login"
	"github.com/gofiber/fiber/v2"
)

// FALTA HACER LOS GROUPS
func RegisterRoutes(app *fiber.App) {
	app.Get("/", l.VerifyToken, c.RenderIndex)
	app.Get("/login", c.RenderLogin)
	app.Post("/login", l.LoginUser)
	app.Get("/logout", l.LogoutUser)
	app.Get("/expiredSession", l.RenderExpiredSession)
	app.Get("/insufficientPermissions", l.RenderInsufficientPermissions)

	m := app.Group("/member")
	p := app.Group("/parent")
	e := app.Group("/enterprise")
	py := app.Group("/payment")

	m.Get("/renderTable", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderMemberTable)
	m.Get("/renderTable/:Page", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderMemberTable)
	m.Get("/addForm", l.VerifyToken, l.VerifyAdmin, c.RenderAddMemberForm)
	m.Post("/add", l.VerifyToken, l.VerifyAdmin, c.AddMember)
	m.Get("/:IdMember/file", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderMemberFile)
	m.Get("/:IdMember/parentTable", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderParentTable)
	m.Put("/:IdMember/edit", l.VerifyToken, l.VerifyAdmin, c.EditMember)
	m.Delete("/:IdMember/delete", l.VerifyToken, l.VerifyAdmin, c.DeleteMember)

	p.Get("/:IdMember/addForm", l.VerifyToken, l.VerifyAdmin, c.RenderAddParentForm)
	p.Post("/:IdMember/add", l.VerifyToken, l.VerifyAdmin, c.AddParent)
	p.Get("/:IdMember/:IdParent/file", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderParentFile)
	p.Delete("/:IdMember/:IdParent/delete", l.VerifyToken, l.VerifyAdmin, c.DeleteParent)
	p.Put("/:IdMember/:IdParent/edit", l.VerifyToken, l.VerifyAdmin, c.EditParent)

	e.Get("/renderTable", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderEnterpriseTable)
	e.Get("/renderTable/:Page", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderEnterpriseTable)
	e.Get("/renderTableSelect", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderEnterpriseTableSelect)
	e.Get("/addForm", l.VerifyToken, l.VerifyAdmin, c.RenderAddEnterpriseForm)
	e.Post("/add", l.VerifyToken, l.VerifyAdmin, c.AddEnterprise)
	e.Get("/:IdEnterprise/file", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderEnterpriseFile)
	e.Get("/:IdEnterprise/memberTable", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderEnterpriseMembers)
	e.Delete("/:IdEnterprise/delete", l.VerifyToken, l.VerifyAdmin, c.DeleteEnterprise)
	e.Put("/:IdEnterprise/edit", l.VerifyToken, l.VerifyAdmin, c.EditEnterprise)
	e.Get("/:IdEnterprise/paymentTable/:Year", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderEnterprisePaymentsTable)
	e.Get("/getAllEnterprisesId", l.VerifyToken, l.VerifyAdminOrGuest, c.GetAllEnterprisesId)

	py.Get("/:IdEnterprise/addForm", l.VerifyToken, l.VerifyAdmin, c.RenderAddPaymentForm)
	py.Post("/:IdEnterprise/add", l.VerifyToken, l.VerifyAdmin, c.AddPayment)
	py.Get("/:IdEnterprise/:IdPayment/file", l.VerifyToken, l.VerifyAdmin, c.RenderPaymentFile)
	py.Put("/:IdEnterprise/:IdPayment/edit", l.VerifyToken, l.VerifyAdmin, c.EditPayment)
	py.Delete("/:IdEnterprise/:IdPayment/delete", l.VerifyToken, l.VerifyAdmin, c.DeletePayment)
	py.Get("/:IdEnterprise/paymentTable", l.VerifyToken, l.VerifyAdminOrGuest, c.RenderEnterprisePaymentsTable)

	app.Get("/test", c.TestNull)
	app.Get("/createMembers", creators.CreateMembers)
	app.Get("/createParents", creators.CreateParents)
	app.Get("/createEnterprises", creators.CreateEnterprises)
	app.Get("/createPayments", creators.CreatePayments)
	app.Get("/renderElectoralMemberList", l.VerifyToken, c.RenderElectoralMemberList)
	app.Get("/pruebaEmpresas", c.RenderPruebaEmpresas)
	app.Get("/backupDB", l.VerifyToken, l.VerifyAdmin, c.BackupDB)
}
