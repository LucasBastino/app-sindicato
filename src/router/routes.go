package router

import (
	c "github.com/LucasBastino/app-sindicato/src/controller"
	"github.com/gofiber/fiber/v2"
)

// FALTA HACER LOS GROUPS
func RegisterRoutes(app *fiber.App) {
	app.Get("/", c.VerifyToken, c.RenderIndex)
	app.Get("/login", c.RenderLogin)
	app.Post("/login", c.LoginUser)

	app.Get("/member/renderTable", c.VerifyToken, c.VerifyAdminOrGuest, c.RenderMemberTable)
	app.Get("/member/renderTable/:Page", c.VerifyToken, c.VerifyAdminOrGuest, c.RenderMemberTable)
	app.Get("/form/createMember", c.VerifyToken, c.VerifyAdmin, c.RenderCreateMemberForm)
	app.Post("/member/create", c.VerifyToken, c.VerifyAdmin, c.CreateMember)
	app.Get("/member/:IdMember/file", c.VerifyToken, c.VerifyAdminOrGuest, c.RenderMemberFile)
	app.Get("/member/:IdMember/parentTable", c.VerifyToken, c.VerifyAdminOrGuest, c.RenderMemberParents)
	app.Put("/member/:IdMember/edit", c.VerifyToken, c.VerifyAdmin, c.EditMember)
	app.Delete("/member/:IdMember/delete", c.VerifyToken, c.VerifyAdmin, c.DeleteMember)
	app.Get("/member/:IdMember/form/createParent", c.VerifyToken, c.VerifyAdmin, c.RenderCreateParentForm)

	app.Post("/parent/:IdMember/create", c.VerifyToken, c.VerifyAdmin, c.CreateParent)
	app.Get("/parent/:IdMember/:IdParent/file", c.VerifyToken, c.VerifyAdminOrGuest, c.RenderParentFile)
	app.Delete("/parent/:IdMember/:IdParent/delete", c.VerifyToken, c.VerifyAdmin, c.DeleteParent)
	app.Put("/parent/:IdMember/:IdParent/edit", c.VerifyToken, c.VerifyAdmin, c.EditParent)

	app.Get("/enterprise/renderTable", c.VerifyToken, c.VerifyAdminOrGuest, c.RenderEnterpriseTable)
	app.Get("/enterprise/renderTable/:Page", c.VerifyToken, c.VerifyAdminOrGuest, c.RenderEnterpriseTable)
	app.Get("/form/createEnterprise", c.VerifyToken, c.VerifyAdmin, c.RenderCreateEnterpriseForm)
	app.Post("/enterprise/create", c.VerifyToken, c.VerifyAdmin, c.CreateEnterprise)
	app.Get("/enterprise/:IdEnterprise/file", c.VerifyToken, c.VerifyAdminOrGuest, c.RenderEnterpriseFile)
	app.Get("/enterprise/:IdEnterprise/memberTable", c.VerifyToken, c.VerifyAdminOrGuest, c.RenderEnterpriseMembers)
	app.Delete("/enterprise/:IdEnterprise/delete", c.VerifyToken, c.VerifyAdmin, c.DeleteEnterprise)
	app.Put("/enterprise/:IdEnterprise/edit", c.VerifyToken, c.VerifyAdmin, c.EditEnterprise)

	app.Get("/test", c.TestNull)
	app.Get("/createMembers", c.CreateMembers)
	app.Get("/createParents", c.CreateParents)
	app.Get("/createEnterprises", c.CreateEnterprises)
	app.Get("/renderElectoralMemberList", c.VerifyToken, c.RenderElectoralMemberList)
	app.Get("/pruebaEmpresas", c.RenderPruebaEmpresas)
}
