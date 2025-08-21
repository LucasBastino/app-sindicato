package controller

import (
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	// "syscall/js"
)

// ------------------------------------

func RenderIndex(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
}


func RenderMemberList(c *fiber.Ctx) error {
	data := fiber.Map{"withMemberTable": true, "withEnterpriseTable": false}
	data["admin"] = c.Locals("claims").(jwt.MapClaims)["admin"]
	data["canWriteMember"] = c.Locals("claims").(jwt.MapClaims)["writeMember"]
	data["canDeleteMember"] = c.Locals("claims").(jwt.MapClaims)["deleteMember"]
	data["canWriteEnterprise"] = c.Locals("claims").(jwt.MapClaims)["writeEnterprise"]
	data["canDeleteEnterprise"] = c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
	return c.Render("tablePage", data)
	// tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	// return tmpl.Execute(c, nil)
}

func RenderEnterpriseList(c *fiber.Ctx) error {
	data := fiber.Map{"withEnterpriseTable": true, "withMemberTable": false}
	data["admin"] = c.Locals("claims").(jwt.MapClaims)["admin"]
	data["canWriteMember"] = c.Locals("claims").(jwt.MapClaims)["writeMember"]
	data["canDeleteMember"] = c.Locals("claims").(jwt.MapClaims)["deleteMember"]
	data["canWriteEnterprise"] = c.Locals("claims").(jwt.MapClaims)["writeEnterprise"]
	data["canDeleteEnterprise"] = c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
	return c.Render("tablePage", data)
}

func RenderTablePage(c *fiber.Ctx) error {
	data := fiber.Map{}
	data["canDeleteMember"] = c.Locals("claims").(jwt.MapClaims)["deleteMember"]
	data["canDeleteEnterprise"] = c.Locals("claims").(jwt.MapClaims)["deleteEnterprise"]
	return c.Render("tablePage", data)
	// tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	// return tmpl.Execute(c, nil)
}

func GetPageFromPath(c *fiber.Ctx) int {
	params := struct {
		Page int `params:"page"`
	}{}

	c.ParamsParser(&params)
	if params.Page <= 1 {
		return 1
	} else {
		return params.Page
	}
}

func GetPaginationData(currentPage, totalRows int) (int, int, int, int) {
	// setting totalPages
	var totalPages int
	// si no hay filas no hay paginas, se pone 1 para que calcule bien el offset
	if totalRows == 0 {
		totalPages = 1
		// si la cantidad de filas es un multiplo de 10 entran justo y no sobran
	} else if totalRows%15 == 0 {
		totalPages = totalRows / 15
		// sino no entran justo y se agrega una pagina mas
	} else {
		totalPages = (totalRows / 15) + 1
	}

	// setting currentPage and offset
	var offset int

	// PODES HACER UNA FUNCION DE ESTO O METER UN SWITCH
	//  si currentPage es menor a 1, currentPage ahora es 1 y muestra los primeros 20
	if currentPage <= 1 {
		offset = 0
	}

	// si currentPage es mayor a totalPages, currentPage ahora es totalPages
	// y muestra los ultimos members
	if currentPage > totalPages {
		currentPage = totalPages
		offset = (currentPage - 1) * 15
	}

	// si currentPage es mayor a 1, muestra los miembros calculando el offset * 15
	if currentPage > 1 {
		offset = (currentPage - 1) * 15
	}

	// setting aproximador
	someBefore := totalPages / 6
	someAfter := totalPages / 6
	// si se pasa de la ultima que te lleve a la ultima
	if someAfter+currentPage > totalPages {
		someAfter = totalPages - currentPage
		// si se pasa de la primera que te lleve a la primera
	} else if currentPage-someBefore < 1 {
		someBefore = currentPage - 1
	}
	return totalPages, offset, someBefore, someAfter
}

func GetTotalPagesArray(totalPages int) []int {
	// devuelve el array para que se pueda recorrer en el template
	var totalPagesArray []int
	if totalPages <= 10 {
		for i := 1; i <= totalPages; i++ {
			totalPagesArray = append(totalPagesArray, i)
		}
	}
	return totalPagesArray
}

func getEnterpriseName(idEnterprise int) (string, error) {
	enterpriseName := ""
	result, err := database.DB.Query("SELECT Name FROM EnterpriseTable WHERE IdEnterprise = ?", idEnterprise)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return "", er.QueryError
	}
	for result.Next() {
		err = result.Scan(&enterpriseName)
		if err != nil {
			er.ScanError.Msg = err.Error()
			return "", er.ScanError
		}
	}
	return enterpriseName, nil
}

func RenderElectoralMemberList(c *fiber.Ctx) error {
	var m models.MemberWithEnterpriseName
	var mm []models.MemberWithEnterpriseName
	result, err := database.DB.Query("SELECT M.MemberNumber, M.LastName, M.Name, M.DNI, E.Name from MemberTable M INNER JOIN EnterpriseTable E ON M.IdEnterprise = E.IdEnterprise WHERE M.IdEnterprise != '1' AND M.Affiliated = true ORDER BY LastName ASC")
	if err != nil {
		// guardar el err en algun lado
		return er.CheckError(c, er.QueryError)
	}
	for result.Next() {
		err = result.Scan(&m.MemberNumber, &m.LastName, &m.Name, &m.DNI, &m.EnterpriseName)
		if err != nil {
			// guardar el err en algun lado
			return er.CheckError(c, er.ScanError)
		}
		mm = append(mm, m)
	}
	result.Close()
	return c.Render("electoralMemberList", fiber.Map{"members": mm})
}

func RenderPruebaEmpresas(c *fiber.Ctx) error {
	var enterprise models.Enterprise
	var enterprises []models.Enterprise
	result, err := database.DB.Query("SELECT IdEnterprise FROM EnterpriseTable ORDER BY Name")
	if err != nil {
		// guardar el err en algun lado
		er.QueryError.Msg = err.Error()
		return er.CheckError(c, er.QueryError)
	}
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name)
		if err != nil {
			// guardar el err en algun lado
			er.ScanError.Msg = err.Error()
			return er.CheckError(c, er.ScanError)
		}
		enterprises = append(enterprises, enterprise)
	}
	result.Close()
	return c.Render("pruebaEmpresas", fiber.Map{"enterprises": enterprises})
}

func RenderLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func formatTimeStamps(cA, uA time.Time) (string, string, error) {
	loc, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		er.FormatError.Msg = err.Error()
		return "", "", er.FormatError
	}

	return cA.In(loc).Format("02-01-2006 15:04:05"), uA.In(loc).Format("02-01-2006 15:04:05"), nil
}

func RenderRegisterUserForm(c *fiber.Ctx) error {
	// admin := c.Locals("claims").(jwt.MapClaims)["admin"]
	// return c.Render("register", fiber.Map{"admin": admin})
	// borrar el de abajo despues ↓
	return c.Render("register", fiber.Map{"admin": true})
}
