package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

type Parent struct {
	IdParent int
	Name     string
	LastName string
	Rel      string
	Gender   string
	Birthday string
	CUIL     string
	IdMember int
}

func (parent Parent) InsertModel() Parent {
	insert, err := database.DB.Query(fmt.Sprintf(`
		INSERT INTO ParentTable 
		(Name,
		LastName,
		Rel,
		Gender,
		Birthday,
		CUIL,
		IdMember)
		VALUES ('%s','%s','%s', '%s', '%s', '%s', '%d')`,
		parent.Name,
		parent.LastName,
		parent.Rel,
		parent.Gender,
		parent.Birthday,
		parent.CUIL,
		parent.IdMember))
	if err != nil {
		// DBError{"INSERT Parent"}.Error(err)
		fmt.Println("error inserting parent")
		panic(err)
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT * FROM ParentTable 
		WHERE IdParent = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		fmt.Print(err)
	}
	p, _ := parent.ScanResult(result, true)
	return p
}

func (parent Parent) DeleteModel() {
	delete, err := database.DB.Query(fmt.Sprintf(`
		DELETE FROM ParentTable 
		WHERE IdParent = '%d'`,
		parent.IdParent))
	if err != nil {
		// DBError{"DELETE Parent"}.Error(err)
		fmt.Println("error deleting parent")
		panic(err)
	}
	defer delete.Close()

}

func (parent Parent) EditModel() {
	update, err := database.DB.Query(fmt.Sprintf(`
		UPDATE ParentTable 
		SET Name = '%s',
		LastName = '%s',
		Rel = '%s',
		Gender = '%s',
		Birthday = '%s',
		CUIL = '%s',
		IdMember = '%d'
		WHERE IdParent = '%d'`,
		parent.Name,
		parent.LastName,
		parent.Rel,
		parent.Gender,
		parent.Birthday,
		parent.CUIL,
		parent.IdMember,
		parent.IdParent))
	if err != nil {
		// DBError{"UPDATE Parent"}.Error(err)
		fmt.Println("error updating parent")
		panic(err)
	}
	defer update.Close()
}

func (parent Parent) GetIdModel(c *fiber.Ctx) int {
	params := struct {
		IdParent int `params:"IdParent"`
	}{}
	c.ParamsParser(&params)
	return params.IdParent
}

func (parent Parent) SearchOneModelById(c *fiber.Ctx) Parent {
	IdParent := parent.GetIdModel(c)
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT * FROM ParentTable
		WHERE IdParent = '%d'`, IdParent))
	if err != nil {
		fmt.Println("error searching parent by id")
		panic(err)
	}
	p, _ := parent.ScanResult(result, true)
	return p
}

func (parent Parent) SearchModels(c *fiber.Ctx, offset int) ([]Parent, string) {
	searchKey := c.FormValue("search-key")
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT IdParent,
		Name,
		Rel
		FROM ParentTable 
		WHERE Name LIKE '%%%s%%' OR Rel LIKE '%%%s%%'
		LIMIT 10 OFFSET %d`,
		searchKey, searchKey, offset))
	if err != nil {
		fmt.Println("error searching Parent in DB")
		panic(err)
	}
	_, parents := parent.ScanResult(result, false)
	return parents, searchKey
}

func (parent Parent) ValidateFields(c *fiber.Ctx) map[string]string {
	errorMap := map[string]string{}
	if strings.TrimSpace(c.FormValue("name")) == "" {
		errorMap["name"] = "el campo Nombre no puede estar vacio"
	}
	if strings.TrimSpace(c.FormValue("rel")) == "" {
		errorMap["rel"] = "el campo Parentesco no puede estar vacio"
	}
	return errorMap
}

func (parent Parent) GetTotalRows(c *fiber.Ctx) int {
	var totalRows int
	searchKey := c.FormValue("search-key")
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM ParentTable 
		WHERE Name LIKE '%%%s%%'`, searchKey))
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}
	return totalRows
}

func (parent Parent) GetFiberMap(parents []Parent, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return fiber.Map{
		"model":           "parent",
		"parents":         parents,
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

func (parent Parent) GetAllModels() []Parent {
	result, err := database.DB.Query("SELECT * FROM ParentTable")
	if err != nil {
		fmt.Println("error searching parent in DB")
		panic(err)
	}
	_, parents := parent.ScanResult(result, false)
	return parents
}

func (parent Parent) ScanResult(result *sql.Rows, onlyOne bool) (Parent, []Parent) {
	var p Parent
	var parents []Parent
	for result.Next() {
		err := result.Scan(
			&p.IdParent,
			&p.Name,
			&p.LastName,
			&p.Rel,
			&p.Gender,
			&p.Birthday,
			&p.CUIL,
			&p.IdMember,
		)
		if err != nil {
			fmt.Println("error scanning parent")
			panic(err)
		}
		if !onlyOne {
			parents = append(parents, p)
		}
	}
	result.Close()
	return p, parents
}
