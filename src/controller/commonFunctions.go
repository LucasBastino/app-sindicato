package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	// "syscall/js"
)

// ------------------------------------

func RenderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
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
	// si no hay filas no hay paginas
	if totalRows == 0 {
		totalPages = 0
		// si la cantidad de filas es un multiplo de 10 entran justo y no sobran
	} else if totalRows%10 == 0 {
		totalPages = totalRows / 10
		// sino no entran justo y se agrega una pagina mas
	} else {
		totalPages = (totalRows / 10) + 1
	}
	fmt.Println("totalPages at line 44:", totalPages)
	fmt.Println("currentPage at line 44:", currentPage)

	// setting currentPage and offset
	var offset int

	// PODES HACER UNA FUNCION DE ESTO O METER UN SWITCH
	//  si currentPage es menor a 1, currentPage ahora es 1 y muestra los primeros 10
	if currentPage <= 1 {
		offset = 0
	}

	// si currentPage es mayor a totalPages, currentPage ahora es totalPages
	// y muestra los ultimos members
	if currentPage > totalPages {
		currentPage = totalPages
		offset = (currentPage - 1) * 10
	}

	// si currentPage es mayor a 1, muestra los miembros calculando el offset * 10
	if currentPage > 1 {
		offset = (currentPage - 1) * 10
	}

	fmt.Println("totalPages at line 56:", totalPages)
	fmt.Println("currentPage at line 56:", currentPage)

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
	var totalPagesArray []int
	if totalPages <= 10 {
		for i := 1; i <= totalPages; i++ {
			totalPagesArray = append(totalPagesArray, i)
		}
	}
	return totalPagesArray
}
