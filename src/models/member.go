package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

type Member struct {
	IdMember      int
	Name          string
	LastName      string
	DNI           string
	Birthday      string
	Gender        string
	MaritalStatus string
	Phone         string
	Email         string
	Address       string
	PostalCode    string
	District      string
	MemberNumber  string
	CUIL          string
	IdEnterprise  int
	Category      string
	EntryDate     string
}

func (m Member) Imprimir() {
	fmt.Println(m)
}

func (member Member) InsertModel() Member {
	insert, err := database.DB.Query(fmt.Sprintf(`
		INSERT INTO MemberTable
		(Name,
		LastName,
		DNI,
		Birthday,
		Gender,
		MaritalStatus,
		Phone,
		Email,
		Address,
		PostalCode,
		District,
		MemberNumber,
		CUIL,
		IdEnterprise,
		Category,
		EntryDate) 
		VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%d','%s','%s')`,
		member.Name,
		member.LastName,
		member.DNI,
		member.Birthday,
		member.Gender,
		member.MaritalStatus,
		member.Phone,
		member.Email,
		member.Address,
		member.PostalCode,
		member.District,
		member.MemberNumber,
		member.CUIL,
		member.IdEnterprise,
		member.Category,
		member.EntryDate))
	if err != nil {
		// DBError{"INSERT MEMBER"}.Error(err)
		fmt.Println("error insertando en la DB")
		panic(err)
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT * FROM MemberTable 
		WHERE IdMember = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		fmt.Print(err)
	}
	m, _ := member.ScanResult(result, true)
	return m
}

func (member Member) DeleteModel() {
	delete, err := database.DB.Query(fmt.Sprintf(`
		DELETE FROM MemberTable 
		WHERE IdMember = '%d'`,
		member.IdMember))
	if err != nil {
		// DBError{"DELETE MEMBER"}.Error(err)
		fmt.Println("error deleting member")
		panic(err)
	}
	defer delete.Close()

}

func (member Member) EditModel() {
	update, err := database.DB.Query(fmt.Sprintf(`
		UPDATE MemberTable
		SET
		Name = '%s',
		LastName = '%s',
		DNI = '%s',
		Birthday = '%s',
		Gender = '%s',
		MaritalStatus = '%s',
		Phone = '%s',
		Email = '%s',
		Address = '%s',
		PostalCode = '%s',
		District = '%s',
		MemberNumber = '%s',
		CUIL = '%s',
		IdEnterprise = '%d',
		Category = '%s',
		EntryDate = '%s'
		WHERE IdMember = '%d'`,
		member.Name,
		member.LastName,
		member.DNI,
		member.Birthday,
		member.Gender,
		member.MaritalStatus,
		member.Phone,
		member.Email,
		member.Address,
		member.PostalCode,
		member.District,
		member.MemberNumber,
		member.CUIL,
		member.IdEnterprise,
		member.Category,
		member.EntryDate,
		member.IdMember))
	if err != nil {
		// DBError{"UPDATE MEMBER"}.Error(err)
		fmt.Println("error updating member")
		panic(err)
	}
	defer update.Close()
}

func (member Member) GetIdModel(c *fiber.Ctx) int {
	params := struct {
		IdMember int `params:"IdMember"`
	}{}

	c.ParamsParser(&params)
	return params.IdMember
}

func (member Member) SearchOneModelById(c *fiber.Ctx) Member {
	IdMember := member.GetIdModel(c)
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT * FROM MemberTable 
		WHERE IdMember = '%d'`,
		IdMember))
	if err != nil {
		fmt.Println("error searching member by Id")
		panic(err)
	}
	m, _ := member.ScanResult(result, true)
	return m
}

func (member Member) SearchModels(c *fiber.Ctx, offset int) ([]Member, string) {
	searchKey := c.FormValue("search-key")
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT * FROM MemberTable 
		WHERE 
		Name LIKE '%%%s%%' OR DNI LIKE '%%%s%%'
		LIMIT 10 OFFSET %d`,
		searchKey, searchKey, offset))
	if err != nil {
		fmt.Println("error searching member in DB")
		panic(err)
	}
	_, members := member.ScanResult(result, false)
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
	if c.FormValue("id-enterprise") == "" {
		errorMap["enterprise"] = "hay que elegir una empresa"
	}
	return errorMap
}

func (member Member) GetTotalRows(c *fiber.Ctx) int {
	var totalRows int
	searchKey := c.FormValue("search-key")
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM MemberTable 
		WHERE Name LIKE '%%%s%%'`, searchKey))
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

func (member Member) GetAllModels() []Member {
	result, err := database.DB.Query("SELECT * FROM MemberTable")
	if err != nil {
		fmt.Println("error searching member in DB")
		panic(err)
	}
	_, members := member.ScanResult(result, false)
	return members
}

func CheckIdEnterprise(tempIdEnterprise sql.NullInt16) int {
	if tempIdEnterprise.Valid {
		return int(tempIdEnterprise.Int16)
	} else {
		return 0
	}
}

func (member Member) ScanResult(result *sql.Rows, onlyOne bool) (Member, []Member) {
	var m Member
	var members []Member
	var tempIdEnterprise sql.NullInt16
	for result.Next() {
		err := result.Scan(
			&m.IdMember,
			&m.Name,
			&m.LastName,
			&m.DNI,
			&m.Birthday,
			&m.Gender,
			&m.MaritalStatus,
			&m.Phone,
			&m.Email,
			&m.Address,
			&m.PostalCode,
			&m.District,
			&m.MemberNumber,
			&m.CUIL,
			&tempIdEnterprise,
			&m.Category,
			&m.EntryDate,
		)
		if err != nil {
			fmt.Println("error scanning member")
			panic(err)
		}
		m.IdEnterprise = CheckIdEnterprise(tempIdEnterprise)
		if !onlyOne {
			members = append(members, m)
		}
	}
	result.Close()
	return m, members
}
