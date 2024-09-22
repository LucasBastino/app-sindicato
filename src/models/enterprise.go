package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

type Enterprise struct {
	IdEnterprise int
	Name         string
	Address      string
}

func (enterprise Enterprise) InsertModel() Enterprise {
	insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO EnterpriseTable (Name, Address) VALUES ('%s','%s')", enterprise.Name, enterprise.Address))
	if err != nil {
		// DBError{"INSERT Enterprise"}.Error(err)
		fmt.Println("error insertando en la DB")
		panic(err)
	}
	insert.Close()
	var e Enterprise
	result, err := database.DB.Query("SELECT * FROM EnterpriseTable WHERE IdEnterprise = (SELECT LAST_INSERT_ID())")
	if err != nil {
		fmt.Print(err)
	}
	result.Next()
	err = result.Scan(&e.IdEnterprise, &e.Name, &e.Address)
	if err != nil {
		fmt.Println(err)
	}
	result.Close()
	return e
}

func (enterprise Enterprise) DeleteModel() {
	delete, err := database.DB.Query(fmt.Sprintf("DELETE FROM EnterpriseTable WHERE IdEnterprise = '%v'", enterprise.IdEnterprise))
	if err != nil {
		// DBError{"DELETE Enterprise"}.Error(err)
		fmt.Println("error deleting Enterprise")
		panic(err)
	}
	defer delete.Close()

}

func (enterprise Enterprise) EditModel() {
	update, err := database.DB.Query(fmt.Sprintf("UPDATE EnterpriseTable SET Name = '%s', Address = '%s' WHERE IdEnterprise = '%v'", enterprise.Name, enterprise.Address, enterprise.IdEnterprise))
	if err != nil {
		// DBError{"UPDATE Enterprise"}.Error(err)
		fmt.Println("error updating Enterprise")
		panic(err)
	}
	defer update.Close()
}

func (enterprise Enterprise) GetIdModel(c *fiber.Ctx) int {
	params := struct {
		IdEnterprise int `params:"IdEnterprise"`
	}{}

	c.ParamsParser(&params)

	return params.IdEnterprise

	// IdEnterpriseStr := c.PathValue("IdEnterprise")
	// IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
	// if err != nil {
	// 	fmt.Println("error converting type")
	// 	panic(err)
	// }
	// return IdEnterprise
}

func (enterprise Enterprise) SearchOneModelById(c *fiber.Ctx) Enterprise {
	IdEnterprise := enterprise.GetIdModel(c)
	result, err := database.DB.Query(fmt.Sprintf("SELECT IdEnterprise, Name, Address FROM EnterpriseTable WHERE IdEnterprise = '%v'", IdEnterprise))
	if err != nil {
		fmt.Println("error searching enterprise by id")
		panic(err)
	}

	var e Enterprise
	for result.Next() {
		err = result.Scan(&e.IdEnterprise, &e.Name, &e.Address)
		if err != nil {
			fmt.Println("error scanning Enterprise")
			panic(err)
		}
	}
	defer result.Close()
	return e
}

func (enterprise Enterprise) SearchModels(c *fiber.Ctx, offset int) ([]Enterprise, string) {
	searchKey := c.FormValue("search-key")
	var enterprises []Enterprise
	var e Enterprise

	result, err := database.DB.Query(fmt.Sprintf(`SELECT * FROM EnterpriseTable WHERE Name LIKE '%%%s%%' OR Address LIKE '%%%s%%' LIMIT 10 OFFSET %d`, searchKey, searchKey, offset))
	if err != nil {
		fmt.Println("error searching Enterprise in DB")
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&e.IdEnterprise, &e.Name, &e.Address)
		if err != nil {
			fmt.Println("error scanning Enterprise")
			panic(err)
		}
		enterprises = append(enterprises, e)
	}
	defer result.Close()
	return enterprises, searchKey
}

func (enterprise Enterprise) ValidateFields(c *fiber.Ctx) map[string]string {
	errorMap := map[string]string{}
	if strings.TrimSpace(c.FormValue("name")) == "" {
		errorMap["name"] = "el campo Nombre no puede estar vacio"
	}
	if strings.TrimSpace(c.FormValue("address")) == "" {
		errorMap["address"] = "el campo Direccion no puede estar vacio"
	}
	return errorMap
}

func (enterprise Enterprise) GetTotalRows(c *fiber.Ctx) int {
	var totalRows int
	searchKey := c.FormValue("search-key")
	row := database.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM EnterpriseTable WHERE Name LIKE '%%%s%%'", searchKey))
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}
	return totalRows
}

func (enterprise Enterprise) GetFiberMap(enterprises []Enterprise, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return fiber.Map{
		"model":           "enterprise",
		"enterprises":     enterprises,
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

func (enterprise Enterprise) GetAllModels() []Enterprise {
	var enterprises []Enterprise
	var e Enterprise

	result, err := database.DB.Query("SELECT * FROM EnterpriseTable")
	if err != nil {
		fmt.Println("error searching enterprise in DB")
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&e.IdEnterprise, &e.Name, &e.Address)
		if err != nil {
			fmt.Println("error scanning enterprise")
			panic(err)
		}
		enterprises = append(enterprises, e)
	}
	defer result.Close()
	return enterprises
}

func (e Enterprise) ScanResult(result *sql.Rows) Enterprise {
	return Enterprise{}
}
