package models

import (
	"database/sql"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
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

func (enterprise Enterprise) InsertModel() (Enterprise, error) {
	insert, err := database.DB.Query(`
		INSERT INTO EnterpriseTable 
		(Name,
		EnterpriseNumber,
		Address, 
		CUIT, 
		District, 
		PostalCode, 
		Phone)
		VALUES ('?','?','?','?','?','?', '?')`,
		enterprise.Name,
		enterprise.EnterpriseNumber,
		enterprise.Address,
		enterprise.CUIT,
		enterprise.District,
		enterprise.PostalCode,
		enterprise.Phone)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Enterprise{}, er.QueryError
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT
		*
		FROM EnterpriseTable
		WHERE IdEnterprise = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Enterprise{}, er.QueryError
	}
	e, _, err := enterprise.ScanResult(result, true)
	if err != nil {
		return Enterprise{}, err
	}
	return e, nil
}

func (enterprise Enterprise) DeleteModel() error {
	if enterprise.IdEnterprise == 1 {
		// cambiar este log
		er.InsufficientPermisionsError.Msg = "permisos insuficientes"
		return er.InsufficientPermisionsError
	}
	delete, err := database.DB.Query(`
		DELETE FROM EnterpriseTable
		WHERE IdEnterprise = '?'`, enterprise.IdEnterprise)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return er.QueryError
	}
	defer delete.Close()
	return nil
}

func (enterprise Enterprise) UpdateModel() (Enterprise, error) {
	update, err := database.DB.Query(`
		UPDATE EnterpriseTable 
		SET 
		Name = '?', 
		EnterpriseNumber = '?',
		Address = '?', 
		CUIT = '?', 
		District = '?', 
		PostalCode = '?', 
		Phone = '?' 
		WHERE IdEnterprise = '?'`,
		enterprise.Name,
		enterprise.EnterpriseNumber,
		enterprise.Address,
		enterprise.CUIT,
		enterprise.District,
		enterprise.PostalCode,
		enterprise.Phone,
		enterprise.IdEnterprise)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Enterprise{}, er.QueryError
	}
	update.Close()
	result, err := database.DB.Query(`
		SELECT
		*
		FROM EnterpriseTable
		WHERE IdEnterprise = ?`, enterprise.IdEnterprise)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Enterprise{}, er.QueryError
	}
	e, _, err := enterprise.ScanResult(result, true)
	if err != nil {
		return Enterprise{}, err
	}
	return e, nil
}

func (enterprise Enterprise) GetIdModel(c *fiber.Ctx) (int, error) {
	params := struct {
		IdEnterprise int `params:"IdEnterprise"`
	}{}

	err := c.ParamsParser(&params)
	if err != nil {
		er.ParamsError.Msg = err.Error()
		return 0, er.ParamsError
	}

	return params.IdEnterprise, nil
}

func (enterprise Enterprise) SearchOneModelById(c *fiber.Ctx) (Enterprise, error) {
	IdEnterprise, err := enterprise.GetIdModel(c)
	if err != nil {
		return Enterprise{}, err
	}
	result, err := database.DB.Query(`
		SELECT
		*
		FROM EnterpriseTable 
		WHERE IdEnterprise = '?'`, IdEnterprise)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Enterprise{}, er.QueryError
	}

	e, _, err := enterprise.ScanResult(result, true)
	if err != nil {
		return Enterprise{}, err
	}
	return e, nil

}

func (enterprise Enterprise) SearchModels(c *fiber.Ctx, offset int) ([]Enterprise, string, error) {
	searchKey := c.FormValue("search-key")
	result, err := database.DB.Query(`
		SELECT
		*
		FROM EnterpriseTable 
		WHERE 
		Name LIKE '%?%' OR Address LIKE '%?%' 
		ORDER BY Name ASC
		LIMIT 15 OFFSET ?`,
		searchKey, searchKey, offset)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, "", er.QueryError
	}
	_, ee, err := enterprise.ScanResult(result, false)
	if err != nil {
		return nil, "", err
	}
	return ee, searchKey, nil
}

// v1
/* func (enterprise Enterprise) ValidateFields(c *fiber.Ctx) (map[string]string, error) {
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
	if len(errorMap) > 1 {

		return errorMap, er.ValidationError
	}
	return errorMap, nil
} */

// v2
func (enterprise Enterprise) ValidateFields(c *fiber.Ctx) error {
	validateFunctions := []func(*fiber.Ctx) error{
		ValidateEnterpriseName,
		ValidateEnterpriseNumber,
		ValidateAddress,
		ValidateCUIT,
		ValidateDistrict,
		ValidatePostalCode,
		ValidatePhone,
	}

	for _, vF := range validateFunctions {
		if err := vF(c); err != nil {
			return err
		} else {
			continue
		}
	}
	return nil
}

func (enterprise Enterprise) GetTotalRows(c *fiber.Ctx) (int, error) {
	var totalRows int
	searchKey := c.FormValue("search-key")
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM EnterpriseTable 
		WHERE Name LIKE '%?%'`, searchKey)
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		er.ScanError.Msg = err.Error()
		return 0, er.ScanError
	}
	return totalRows, nil
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

func (enterprise Enterprise) GetAllModels() ([]Enterprise, error) {
	result, err := database.DB.Query(`
		SELECT 
		*
		FROM EnterpriseTable`)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, er.QueryError
	}
	_, ee, err := enterprise.ScanResult(result, false)
	if err != nil {
		return nil, err
	}
	return ee, nil
}

func (enterprise Enterprise) ScanResult(result *sql.Rows, onlyOne bool) (Enterprise, []Enterprise, error) {
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
			er.ScanError.Msg = err.Error()
			return Enterprise{}, nil, er.ScanError
		}
		if !onlyOne {
			ee = append(ee, e)
		}
	}
	result.Close()
	return e, ee, nil
}

func (enterprise Enterprise) CheckDeleted(idEnterprise int) (bool, error) {
	var totalRows int
	// row := database.DB.QueryRow(fmt.Sprintf(`
	// 	SELECT COUNT(*) FROM EnterpriseTable
	// 	WHERE IdEnterprise = '%d'`, enterprise.IdEnterprise))
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM EnterpriseTable 
		WHERE IdEnterprise = '?'`, idEnterprise)
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		er.ScanError.Msg = err.Error()
		return false, er.ScanError
	}
	if totalRows == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
