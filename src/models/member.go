package models

import (
	"fmt"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

type Member struct {
	IdMember     int
	Name         string
	DNI          string
	IdEnterprise int
}

func (m Member) Imprimir() {
	fmt.Println(m)
}

func (newMember Member) InsertModel() Member {
	insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO MemberTable (Name, DNI, IdEnterprise) VALUES ('%s','%s', '%d')", newMember.Name, newMember.DNI, newMember.IdEnterprise))
	if err != nil {
		// DBError{"INSERT MEMBER"}.Error(err)
		fmt.Println("error insertando en la DB")
		panic(err)
	}
	insert.Close()
	var member Member
	result, err := database.DB.Query("SELECT * FROM MemberTable WHERE IdMember = (SELECT LAST_INSERT_ID())")
	if err != nil {
		fmt.Print(err)
	}
	result.Next()
	err = result.Scan(&member.IdMember, &member.Name, &member.DNI, &member.IdEnterprise)
	if err != nil {
		fmt.Println(err)
	}
	result.Close()
	return member
}

func (m Member) DeleteModel() {
	delete, err := database.DB.Query(fmt.Sprintf("DELETE FROM MemberTable WHERE IdMember = '%v'", m.IdMember))
	if err != nil {
		// DBError{"DELETE MEMBER"}.Error(err)
		fmt.Println("error deleting member")
		panic(err)
	}
	defer delete.Close()

}

func (m Member) EditModel() {
	update, err := database.DB.Query(fmt.Sprintf("UPDATE MemberTable SET Name = '%s', DNI = '%s', IdEnterprise = '%d' WHERE IdMember = '%v'", m.Name, m.DNI, m.IdEnterprise, m.IdMember))
	if err != nil {
		// DBError{"UPDATE MEMBER"}.Error(err)
		fmt.Println("error updating member")
		panic(err)
	}
	defer update.Close()
}

func (m Member) GetIdModel(c *fiber.Ctx) int {
	params := struct {
		IdMember int `params:"IdMember"`
	}{}

	c.ParamsParser(&params)
	return params.IdMember
}

func (m Member) SearchOneModelById(c *fiber.Ctx) Member {
	IdMember := m.GetIdModel(c)
	result, err := database.DB.Query(fmt.Sprintf("SELECT IdMember, Name, DNI, IdEnterprise FROM MemberTable WHERE IdMember = '%v'", IdMember))
	if err != nil {
		fmt.Println("error searching member by Id")
		panic(err)
	}

	var member Member
	for result.Next() {
		err = result.Scan(&member.IdMember, &member.Name, &member.DNI, &member.IdEnterprise)
		if err != nil {
			fmt.Println("error scanning member")
			panic(err)
		}
	}
	defer result.Close()
	return member
}

func (m Member) SearchModels(c *fiber.Ctx, offset int) ([]Member, string) {
	searchKey := c.FormValue("search-key")
	var members []Member
	var member Member
	result, err := database.DB.Query(fmt.Sprintf(`SELECT * FROM MemberTable WHERE Name LIKE '%%%s%%' OR DNI LIKE '%%%s%%' LIMIT 10 OFFSET %d`, searchKey, searchKey, offset))
	if err != nil {
		fmt.Println("error searching member in DB")
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning member")
			panic(err)
		}
		members = append(members, member)
	}
	defer result.Close()
	return members, searchKey
}

func (m Member) ValidateFields(c *fiber.Ctx) map[string]string {
	errorMap := map[string]string{}

	if strings.TrimSpace(c.FormValue("name")) == "" {
		errorMap["name"] = "el campo Nombre no puede estar vacio"
	}
	// consultar que sea alfanumerico
	if strings.TrimSpace(c.FormValue("dni")) == "" {
		errorMap["dni"] = "el campo DNI no puede estar vacio"
	}
	if utf8.RuneCountInString(c.FormValue("dni")) > 8 {
		errorMap["dni"] = "el DNI no puede tener mas de 8 caracteres"
	}
	return errorMap
}

func (m Member) GetTotalRows(c *fiber.Ctx) int {
	var totalRows int
	searchKey := c.FormValue("search-key")
	row := database.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM MemberTable WHERE Name LIKE '%%%s%%'", searchKey))
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}
	return totalRows
}

func (m Member) GetFiberMap(members []Member, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return fiber.Map{
		"model":           "member",
		"members":         members,
		"searchKey":       searchKey,
		"currentPage":     currentPage,
		"firstPage":       1,
		"previousPage":    currentPage - 1,
		"someBefore":      currentPage - someBefore,
		"sixBefore":       currentPage - 6,
		"fiveBefore":      currentPage - 5,
		"fourBefore":      currentPage - 4,
		"threeBefore":     currentPage - 3,
		"twoBefore":       currentPage - 2,
		"twoAfter":        currentPage + 2,
		"threeAfter":      currentPage + 3,
		"fourAfter":       currentPage + 4,
		"fiveAfter":       currentPage + 5,
		"sixAfter":        currentPage + 6,
		"someAfter":       currentPage + someAfter,
		"nextPage":        currentPage + 1,
		"lastPage":        totalPages,
		"totalPages":      totalPages,
		"totalPagesArray": totalPagesArray,
	}
}
