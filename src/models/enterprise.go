package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

type Enterprise struct {
	IdEnterprise     int
	Name             string
	EnterpriseNumber string
	Address          string
	CUIT             string
	District         string
	PostalCode       string
	Phone            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (enterprise Enterprise) InsertModel() Enterprise {
	insert, err := database.DB.Query(fmt.Sprintf(`
		INSERT INTO EnterpriseTable 
		(Name,
		EnterpriseNumber,
		Address, 
		CUIT, 
		District, 
		PostalCode, 
		Phone)
		VALUES ('%s','%s','%s','%s','%s','%s', '%s')`,
		enterprise.Name,
		enterprise.EnterpriseNumber,
		enterprise.Address,
		enterprise.CUIT,
		enterprise.District,
		enterprise.PostalCode,
		enterprise.Phone))
	if err != nil {
		// DBError{"INSERT Enterprise"}.Error(err)
		fmt.Println("error insertando en la DB")
		panic(err)
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT
		*
		FROM EnterpriseTable
		WHERE IdEnterprise = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		fmt.Print(err)
	}
	e, _ := enterprise.ScanResult(result, true)
	return e
}

func (enterprise Enterprise) DeleteModel() {
	if enterprise.IdEnterprise == 1 {
		fmt.Println("internal error: cannot delete enterprise 1")
		return
	}
	delete, err := database.DB.Query(fmt.Sprintf(`
		DELETE FROM EnterpriseTable
		WHERE IdEnterprise = '%d'`, enterprise.IdEnterprise))
	if err != nil {
		// DBError{"DELETE Enterprise"}.Error(err)
		fmt.Println("error deleting Enterprise")
		panic(err)
	}
	defer delete.Close()

}

func (enterprise Enterprise) UpdateModel() Enterprise {
	update, err := database.DB.Query(fmt.Sprintf(`
		UPDATE EnterpriseTable 
		SET 
		Name = '%s', 
		EnterpriseNumber = '%s',
		Address = '%s', 
		CUIT = '%s', 
		District = '%s', 
		PostalCode = '%s', 
		Phone = '%s' 
		WHERE IdEnterprise = '%d'`,
		enterprise.Name,
		enterprise.EnterpriseNumber,
		enterprise.Address,
		enterprise.CUIT,
		enterprise.District,
		enterprise.PostalCode,
		enterprise.Phone,
		enterprise.IdEnterprise))
	if err != nil {
		// DBError{"UPDATE Enterprise"}.Error(err)
		fmt.Println("error updating Enterprise")
		panic(err)
	}
	update.Close()
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT
		*
		FROM EnterpriseTable
		WHERE IdEnterprise = %d`, enterprise.IdEnterprise))
	if err != nil {
		fmt.Print(err)
	}
	e, _ := enterprise.ScanResult(result, true)
	return e
}

func (enterprise Enterprise) GetIdModel(c *fiber.Ctx) int {
	params := struct {
		IdEnterprise int `params:"IdEnterprise"`
	}{}

	c.ParamsParser(&params)

	return params.IdEnterprise
}

func (enterprise Enterprise) SearchOneModelById(c *fiber.Ctx) Enterprise {
	IdEnterprise := enterprise.GetIdModel(c)
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT
		*
		FROM EnterpriseTable 
		WHERE IdEnterprise = '%d'`, IdEnterprise))
	if err != nil {
		fmt.Println("error searching enterprise by id")
		panic(err)
	}

	e, _ := enterprise.ScanResult(result, true)
	return e

}

func (enterprise Enterprise) SearchModels(c *fiber.Ctx, offset int) ([]Enterprise, string) {
	searchKey := c.FormValue("search-key")
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT
		*
		FROM EnterpriseTable 
		WHERE 
		Name LIKE '%%%s%%' OR Address LIKE '%%%s%%' 
		ORDER BY Name ASC
		LIMIT 15 OFFSET %d`,
		searchKey, searchKey, offset))
	if err != nil {
		fmt.Println("error searching Enterprise in DB")
		panic(err)
	}
	_, ee := enterprise.ScanResult(result, false)
	return ee, searchKey
}

func (enterprise Enterprise) ValidateFields(c *fiber.Ctx) map[string]string {
	errorMap := map[string]string{}
	var valid bool
	var err string
	if valid, err = ValidateEnterpriseName(c); !valid {
		errorMap["enterpriseName"] = err
	}
	if valid, err = ValidateEnterpriseNumber(c); !valid {
		errorMap["enterpriseNumber"] = err
	}
	if valid, err = ValidateAddress(c); !valid {
		errorMap["address"] = err
	}
	if valid, err = ValidateCUIT(c); !valid {
		errorMap["cuit"] = err
	}
	if valid, err = ValidateDistrict(c); !valid {
		errorMap["district"] = err
	}
	if valid, err = ValidatePostalCode(c); !valid {
		errorMap["postalCode"] = err
	}
	if valid, err = ValidatePhone(c); !valid {
		errorMap["phone"] = err
	}
	return errorMap
}

func (enterprise Enterprise) GetTotalRows(c *fiber.Ctx) int {
	var totalRows int
	searchKey := c.FormValue("search-key")
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM EnterpriseTable 
		WHERE Name LIKE '%%%s%%'`, searchKey))
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
	result, err := database.DB.Query(`
		SELECT 
		*
		FROM EnterpriseTable`)
	if err != nil {
		fmt.Println("error searching enterprise in DB")
		panic(err)
	}
	_, ee := enterprise.ScanResult(result, false)
	return ee
}

func (enterprise Enterprise) ScanResult(result *sql.Rows, onlyOne bool) (Enterprise, []Enterprise) {
	var e Enterprise
	var ee []Enterprise
	for result.Next() {
		err := result.Scan(
			&e.IdEnterprise,
			&e.Name,
			&e.EnterpriseNumber,
			&e.Address,
			&e.CUIT,
			&e.District,
			&e.PostalCode,
			&e.Phone,
			&e.CreatedAt,
			&e.UpdatedAt,
		)
		if err != nil {
			fmt.Println("error scanning enterprise")
			panic(err)
		}
		if !onlyOne {
			ee = append(ee, e)
		}
	}
	result.Close()
	return e, ee
}

func (enterprise Enterprise) CheckDeleted(idEnterprise int) bool {
	var totalRows int
	// row := database.DB.QueryRow(fmt.Sprintf(`
	// 	SELECT COUNT(*) FROM EnterpriseTable
	// 	WHERE IdEnterprise = '%d'`, enterprise.IdEnterprise))
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM EnterpriseTable 
		WHERE IdEnterprise = '%d'`, idEnterprise))
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}
	if totalRows == 0 {
		return true
	} else {
		return false
	}
}
