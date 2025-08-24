package router

import (
	c "github.com/LucasBastino/app-sindicato/src/controller"
	creators "github.com/LucasBastino/app-sindicato/src/creators"
	er "github.com/LucasBastino/app-sindicato/src/errors/errorHandler"
	l "github.com/LucasBastino/app-sindicato/src/login"
	pe "github.com/LucasBastino/app-sindicato/src/permissions"
	"github.com/gofiber/fiber/v2"
)

// FALTA HACER LOS GROUPS
func RegisterRoutes(app *fiber.App) {
	app.Get("/", pe.VerifyAuth, l.VerifyToken, c.RenderIndex)
	// borrar esto despues ↓ --------------------
	app.Get("/register", pe.VerifyAuth,c.RenderRegisterUserForm)
	// app.Get("/register", l.VerifyToken, pe.VerifyAdmin, c.RenderRegisterUserForm)
	// app.Post("/register", l.VerifyToken, pe.VerifyAdmin, l.RegisterUser)
	app.Post("/register", pe.VerifyAuth,l.RegisterUser)
	// --------------------------
	app.Get("/login", pe.VerifyAuth,c.RenderLogin)
	app.Post("/login", pe.VerifyAuth,l.LoginUser)
	app.Get("/logout", pe.VerifyAuth,l.LogoutUser)
	app.Get("/expiredSession", pe.VerifyAuth,l.RenderExpiredSession)
	app.Get("/insufficientPermissions", pe.VerifyAuth,l.RenderInsufficientPermissions)

	m := app.Group("/member")
	p := app.Group("/parent")
	e := app.Group("/enterprise")
	py := app.Group("/payment")

	m.Get("/list", pe.VerifyAuth, l.VerifyToken, c.RenderMemberList)
	m.Get("/renderTable", pe.VerifyAuth,l.VerifyToken, c.RenderMemberTable)
	m.Get("/renderTable/:Page", pe.VerifyAuth,l.VerifyToken, c.RenderMemberTable)
	m.Get("/addForm", pe.VerifyAuth,l.VerifyToken, pe.VerifyWriteMember, c.RenderAddMemberForm)
	m.Post("/add", pe.VerifyAuth,l.VerifyToken, pe.VerifyWriteMember, c.AddMember)
	m.Get("/:IdMember/file",pe.VerifyAuth,l.VerifyToken, c.RenderMemberFile)
	m.Get("/:IdMember/parentTable",pe.VerifyAuth, l.VerifyToken, c.RenderParentTable)
	m.Put("/:IdMember/edit",pe.VerifyAuth, l.VerifyToken, pe.VerifyWriteMember, c.EditMember)
	m.Delete("/:IdMember/delete",pe.VerifyAuth, l.VerifyToken, pe.VerifyDeleteMember, c.DeleteMember)

	p.Get("/:IdMember/addForm",pe.VerifyAuth, l.VerifyToken, pe.VerifyWriteParent, c.RenderAddParentForm)
	p.Post("/:IdMember/add",pe.VerifyAuth, l.VerifyToken, pe.VerifyWriteParent, c.AddParent)
	p.Get("/:IdMember/:IdParent/file",pe.VerifyAuth, l.VerifyToken, c.RenderParentFile)
	p.Put("/:IdMember/:IdParent/edit",pe.VerifyAuth, l.VerifyToken, pe.VerifyWriteParent, c.EditParent)
	p.Delete("/:IdMember/:IdParent/delete",pe.VerifyAuth, l.VerifyToken, pe.VerifyDeleteParent, c.DeleteParent)

	e.Get("/list",pe.VerifyAuth, l.VerifyToken, c.RenderEnterpriseList)
	e.Get("/renderTable",pe.VerifyAuth, l.VerifyToken, c.RenderEnterpriseTable)
	e.Get("/renderTable/:Page",pe.VerifyAuth, l.VerifyToken, c.RenderEnterpriseTable)
	e.Get("/renderTableSelect",pe.VerifyAuth, l.VerifyToken, c.RenderEnterpriseTableSelect)
	e.Get("/addForm",pe.VerifyAuth, l.VerifyToken, pe.VerifyWriteEnterprise, c.RenderAddEnterpriseForm)
	e.Post("/add",pe.VerifyAuth, l.VerifyToken, pe.VerifyWriteEnterprise, c.AddEnterprise)
	e.Get("/:IdEnterprise/file",pe.VerifyAuth, l.VerifyToken, c.RenderEnterpriseFile)
	e.Get("/:IdEnterprise/memberTable",pe.VerifyAuth, l.VerifyToken, c.RenderEnterpriseMembers)
	e.Delete("/:IdEnterprise/delete",pe.VerifyAuth, l.VerifyToken, pe.VerifyDeleteEnterprise, c.DeleteEnterprise)
	e.Put("/:IdEnterprise/edit",pe.VerifyAuth, l.VerifyToken, pe.VerifyWriteEnterprise, c.EditEnterprise)
	e.Get("/:IdEnterprise/paymentTable/:Year",pe.VerifyAuth, l.VerifyToken, c.RenderPaymentTable)
	e.Get("/getAllEnterprisesId",pe.VerifyAuth, l.VerifyToken, c.GetAllEnterprisesId)
	e.Get("/getAllEnterprisesNumber",pe.VerifyAuth, l.VerifyToken, c.GetAllEnterprisesNumber)

	py.Get("/:IdEnterprise/addForm",pe.VerifyAuth, l.VerifyToken, pe.VerifyWritePayment, c.RenderAddPaymentForm)
	py.Post("/:IdEnterprise/add",pe.VerifyAuth, l.VerifyToken, pe.VerifyWritePayment, c.AddPayment)
	py.Get("/:IdEnterprise/:IdPayment/file",pe.VerifyAuth, l.VerifyToken, c.RenderPaymentFile)
	py.Put("/:IdEnterprise/:IdPayment/edit",pe.VerifyAuth, l.VerifyToken, pe.VerifyWritePayment, c.EditPayment)
	py.Delete("/:IdEnterprise/:IdPayment/delete",pe.VerifyAuth, l.VerifyToken, pe.VerifyDeletePayment, c.DeletePayment)
	py.Get("/:IdEnterprise/paymentTable",pe.VerifyAuth, l.VerifyToken, c.RenderPaymentTable)

	app.Get("/error", er.RenderError)

	app.Get("/test", c.TestNull)
	app.Get("/createMembers", creators.CreateMembers)
	app.Get("/createParents", creators.CreateParents)
	app.Get("/createEnterprises", creators.CreateEnterprises)
	app.Get("/createPayments", creators.CreatePayments)
	app.Get("/renderElectoralMemberList",pe.VerifyAuth, l.VerifyToken, c.RenderElectoralMemberList)
	app.Get("/pruebaEmpresas", c.RenderPruebaEmpresas)
	app.Get("/backupDB",pe.VerifyAuth, l.VerifyToken, pe.VerifyAdmin, c.BackupDB)
}
