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

	// getting number of rows
	var totalRows int
	row := database.DB.QueryRow("SELECT COUNT(*) FROM MemberTable")
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
		fmt.Println("resto cero")
		// sino no entran justo y se agrega una pagina mas
	} else {
		totalPages = (totalRows / 10) + 1
		fmt.Println("tiene resto")
	}

	// setting currentPage and offset
	var offset int
	if currentPage <= 1 {
		currentPage = 1
		offset = 0
	} else if currentPage > 1 {
		offset = (currentPage - 1) * 10
	} else if currentPage > totalPages {
		currentPage = totalPages
	}

	// getting members from database
	result, err := database.DB.Query(fmt.Sprintf("SELECT Name, DNI from MemberTable ORDER BY Name LIMIT 10 OFFSET %d", offset))
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
	return c.Render("test", fiber.Map{
		"members":      members,
		"currentPage":  currentPage,
		"previousPage": currentPage - 1,
		"nextPage":     currentPage + 1,
		"twoBefore":    currentPage - 2,
		"twoAfter":     currentPage + 2,
		"threeBefore":  currentPage - 3,
		"threeAfter":   currentPage + 3,
		"totalPages":   totalPages,
	})
}
