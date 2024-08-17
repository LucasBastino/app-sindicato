package controller

import (
	"fmt"
	"log"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

func TestOffset(c *fiber.Ctx) error {
	var member models.Member
	var members []models.Member

	// getting page from url
	params := struct {
		Page int `params:"page"`
	}{}
	c.ParamsParser(&params)
	currentPage := params.Page
	fmt.Println("currentPage at line 22:", currentPage)

	// getting number of rows
	var totalRows int
	row := database.DB.QueryRow("SELECT COUNT(*) FROM MemberTable WHERE Name LIKE '%%%a%%'")
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}

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
		currentPage = 1
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

	// getting members from database
	result, err := database.DB.Query(fmt.Sprintf("SELECT Name, DNI from MemberTable WHERE NAME LIKE '%%a%%' ORDER BY Name LIMIT 10 OFFSET %d", offset))
	if err != nil {
		log.Println(err)
	}
	for result.Next() {
		err = result.Scan(&member.Name, &member.DNI)
		if err != nil {
			log.Println(err)
		}
		members = append(members, member)
	}

	var totalPagesArray []int
	if totalPages <= 10 {
		for i := 1; i <= totalPages; i++ {
			totalPagesArray = append(totalPagesArray, i)
		}
	}

	return c.Render("test", fiber.Map{
		"members":         members,
		"currentPage":     currentPage,
		"firstPage":       1,
		"previousPage":    currentPage - 1,
		"someBefore":      currentPage - someBefore,
		"threeBefore":     currentPage - 3,
		"twoBefore":       currentPage - 2,
		"twoAfter":        currentPage + 2,
		"threeAfter":      currentPage + 3,
		"someAfter":       currentPage + someAfter,
		"nextPage":        currentPage + 1,
		"lastPage":        totalPages,
		"totalPages":      totalPages,
		"totalPagesArray": totalPagesArray,
	})
}
