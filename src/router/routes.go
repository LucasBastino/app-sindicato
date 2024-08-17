package router

import (
	c "github.com/LucasBastino/app-sindicato/src/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", c.RenderIndex)

	app.Get("/member/renderTable", c.RenderMemberTable)
	app.Get("/form/createMember", c.RenderCreateMemberForm)
	app.Post("/member/create", c.CreateMember)
	app.Get("/member/:IdMember/file", c.RenderMemberFile)
	app.Get("/member/:IdMember/parentTable", c.RenderMemberParents)
	app.Put("/member/:IdMember/edit", c.EditMember)
	app.Delete("/member/:IdMember/delete", c.DeleteMember)
	app.Get("/member/:IdMember/form/createParent", c.RenderCreateParentForm)

	app.Get("/parent/renderTable", c.RenderParentTable)
	app.Post("/parent/create", c.CreateParent)
	app.Get("/parent/:IdParet/file", c.RenderParentFile)
	app.Delete("/parent/:IdMember/:IdParet/delete", c.DeleteParent)
	app.Put("/parent/:IdParet/edit", c.EditParent)

	app.Get("/enterprise/renderTable", c.RenderEnterpriseTable)
	app.Get("/form/createEnterprise", c.RenderCreateEnterpriseForm)
	app.Post("/enterprise/create", c.CreateEnterprise)
	app.Get("/enterprise/:IdEnterprise/file", c.RenderEnterpriseFile)
	app.Delete("/enterprise/:IdEnterprise/delete", c.DeleteEnterprise)
	// cambiar el de abajo a PUT
	app.Put("/enterprise/:IdEnterprise/edit", c.EditEnterprise)

	// app.Get("/enterPriseTable", c.renderEnterpriseTable)
	// app.Get("/parentTable", c.renderParentTable)

	// app.Get("/form/createParent", c.renderCreateParentForm)
	app.Get("/test/:Page", c.TestOffset)
	app.Get("/createMembers", c.CreateMembers)
}
