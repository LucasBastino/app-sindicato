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

	app.Get("/member/renderTable", c.VerifyToken, c.RenderMemberTable)
	app.Get("/member/renderTable/:Page", c.VerifyToken, c.RenderMemberTable)
	app.Get("/form/createMember", c.VerifyToken, c.RenderCreateMemberForm)
	app.Post("/member/create", c.VerifyToken, c.CreateMember)
	app.Get("/member/:IdMember/file", c.VerifyToken, c.RenderMemberFile)
	app.Get("/member/:IdMember/parentTable", c.VerifyToken, c.RenderMemberParents)
	app.Put("/member/:IdMember/edit", c.VerifyToken, c.EditMember)
	app.Delete("/member/:IdMember/delete", c.VerifyToken, c.DeleteMember)
	app.Get("/member/:IdMember/form/createParent", c.VerifyToken, c.RenderCreateParentForm)

	app.Get("/parent/renderTable", c.VerifyToken, c.RenderParentTable)
	app.Get("/parent/renderTable/:Page", c.VerifyToken, c.RenderParentTable)
	app.Post("/parent/create", c.VerifyToken, c.CreateParent)
	app.Get("/parent/:IdParent/file", c.VerifyToken, c.RenderParentFile)
	app.Delete("/parent/:IdMember/:IdParent/delete", c.VerifyToken, c.DeleteParent)
	app.Put("/parent/:IdParent/edit", c.VerifyToken, c.EditParent)

	app.Get("/enterprise/renderTable", c.VerifyToken, c.RenderEnterpriseTable)
	app.Get("/enterprise/renderTable/:Page", c.VerifyToken, c.RenderEnterpriseTable)
	app.Get("/form/createEnterprise", c.VerifyToken, c.RenderCreateEnterpriseForm)
	app.Post("/enterprise/create", c.VerifyToken, c.CreateEnterprise)
	app.Get("/enterprise/:IdEnterprise/file", c.VerifyToken, c.RenderEnterpriseFile)
	app.Get("/enterprise/:IdEnterprise/memberTable", c.VerifyToken, c.RenderEnterpriseMembers)
	app.Delete("/enterprise/:IdEnterprise/delete", c.VerifyToken, c.DeleteEnterprise)
	// cambiar el de abajo a PUT
	app.Put("/enterprise/:IdEnterprise/edit", c.VerifyToken, c.EditEnterprise)

	// app.Get("/enterPriseTable", c.VerifyToken, c.renderEnterpriseTable)
	// app.Get("/parentTable", c.VerifyToken, c.renderParentTable)

	// app.Get("/form/createParent", c.VerifyToken, c.renderCreateParentForm)
	app.Get("/test", c.TestNull)
	app.Get("/createMembers", c.CreateMembers)
	app.Get("/createParents", c.CreateParents)
	app.Get("/createEnterprises", c.CreateEnterprises)
	app.Get("/prueba", c.Prueba)
	app.Get("/renderElectoralMemberList", c.VerifyToken, c.RenderElectoralMemberList)
	app.Get("/pruebaEmpresas", c.RenderPruebaEmpresas)
}
