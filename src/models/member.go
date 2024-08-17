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
	IdMember int
	Name     string
	DNI      string
}

func (m Member) Imprimir() {
	fmt.Println(m)
}

func (newMember Member) InsertModel() {
	insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO MemberTable (Name, DNI) VALUES ('%s','%s')", newMember.Name, newMember.DNI))
	if err != nil {
		// DBError{"INSERT MEMBER"}.Error(err)
		fmt.Println("error insertando en la DB")
		panic(err)
	}
	defer insert.Close()
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
	update, err := database.DB.Query(fmt.Sprintf("UPDATE MemberTable SET Name = '%s', DNI = '%s' WHERE IdMember = '%v'", m.Name, m.DNI, m.IdMember))
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
	result, err := database.DB.Query(fmt.Sprintf("SELECT IdMember, Name, DNI FROM MemberTable WHERE IdMember = '%v'", IdMember))
	if err != nil {
		fmt.Println("error searching member by Id")
		panic(err)
	}

	var member Member
	for result.Next() {
		err = result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning member")
			panic(err)
		}
	}
	defer result.Close()
	return member
}

func (m Member) SearchModels(c *fiber.Ctx) []Member {
	searchKey := c.FormValue("search-key")
	var members []Member
	var member Member

	// getting page from url
	params := struct {
		Page int `params:"page"`
	}{}
	c.ParamsParser(&params)
	// currentPage := params.Page

	// getting number of rows
	var totalRows int
	row := database.DB.QueryRow("SELECT COUNT(*) FROM MemberTable WHERE Name LIKE '%%%s%%' OR DNI LIKE '%%%s%%'", searchKey, searchKey)
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}

	result, err := database.DB.Query(fmt.Sprintf(`SELECT * FROM MemberTable WHERE Name LIKE '%%%s%%' OR DNI LIKE '%%%s%%'`, searchKey, searchKey))
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
	return members
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
