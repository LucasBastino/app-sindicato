package controller

import (
	"fmt"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	// "syscall/js"
)

// ------------------------------------

func RenderIndex(c *fiber.Ctx) error {
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	return c.Render("index", fiber.Map{"role": role})
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

func getEnterpriseName(idEnterprise int) string {
	enterpriseName := ""
	result, err := database.DB.Query(fmt.Sprintf("SELECT Name FROM EnterpriseTable WHERE IdEnterprise = '%d'", idEnterprise))
	if err != nil {
		fmt.Println("error searching for enterprise name")
	}
	for result.Next() {
		err = result.Scan(&enterpriseName)
		if err != nil {
			fmt.Println("error scanning enterprise name")
		}
	}
	return enterpriseName
}

func RenderElectoralMemberList(c *fiber.Ctx) error {
	var member models.Member
	var members []models.Member
	result, err := database.DB.Query("SELECT Name, DNI from MemberTable")
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		err = result.Scan(&member.Name, &member.DNI)
		if err != nil {
			fmt.Println(err)
		}
		members = append(members, member)
	}
	defer result.Close()
	return c.Render("electoralMemberList", fiber.Map{"members": members})
}

func RenderPruebaEmpresas(c *fiber.Ctx) error {
	var enterprise models.Enterprise
	var enterprises []models.Enterprise
	result, err := database.DB.Query("SELECT IdEnterprise, Name FROM EnterpriseTable ORDER BY Name")
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name)
		if err != nil {
			fmt.Println(err)
		}
		enterprises = append(enterprises, enterprise)
	}
	return c.Render("pruebaEmpresas", fiber.Map{"enterprises": enterprises})
}

func RenderLogin(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{})
}

func formatTimeStamps(cA, uA time.Time) (string, string) {
	loc, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		panic(err)
	}

	return cA.In(loc).Format("02-01-2006 15:04:05"), uA.In(loc).Format("02-01-2006 15:04:05")
}

// Creado: 2024-11-18 19:12:38 +0000 UTC

// func formatTimeStamps(m *models.Member) {
// 	fmt.Println(m)
// 	fmt.Println(&m)
// 	fmt.Println(*m)
// 	m.Name = "Lukitas"
// }
