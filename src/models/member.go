package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
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
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (member Member) InsertModel() (Member, error) {
	// formateo la fecha nac para que empiece con el año
	member.Birthday = FormatToYYYYMMDD(member.Birthday)
	member.EntryDate = FormatToYYYYMMDD(member.EntryDate)
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
		er.QueryError.Msg = err.Error()
		return Member{}, er.QueryError
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT * FROM MemberTable 
		WHERE IdMember = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Member{}, er.QueryError
	}
	m, _, err := member.ScanResult(result, true)
	if err != nil {
		return Member{}, err
	}
	return m, nil
}

func (member Member) DeleteModel() error {
	delete, err := database.DB.Query(fmt.Sprintf(`
		DELETE FROM MemberTable 
		WHERE IdMember = '%d'`,
		member.IdMember))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return er.QueryError
	}
	defer delete.Close()
	return nil
}

func (member Member) UpdateModel() (Member, error) {
	// formateo la fecha nac para que empiece con el año
	member.Birthday = FormatToYYYYMMDD(member.Birthday)
	member.EntryDate = FormatToYYYYMMDD(member.EntryDate)
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
		MemberNumber = %s,
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
		er.QueryError.Msg = err.Error()
		return Member{}, er.QueryError
	}
	update.Close()
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT * FROM MemberTable 
		WHERE IdMember = %d`, member.IdMember))

	if err != nil {
		er.QueryError.Msg = err.Error()
		return Member{}, er.QueryError
	}
	m, _, err := member.ScanResult(result, true)
	if err != nil {
		return Member{}, err
	}
	return m, nil
}

func (member Member) GetIdModel(c *fiber.Ctx) (int, error) {
	// params := struct {
	// 	IdMember int `params:"IdMember"`
	// }{}

	// c.ParamsParser(&params)
	// return params.IdMember

	// hacerlos asi a partir de ahora
	idMember, err := c.ParamsInt("IdMember")
	if err != nil {
		er.ParamsError.Msg = err.Error()
		return 0, er.ParamsError
	}
	return idMember, nil
}

func (member Member) SearchOneModelById(c *fiber.Ctx) (Member, error) {
	IdMember, err := member.GetIdModel(c)
	if err != nil {
		return Member{}, err
	}
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT 
		*
		FROM MemberTable 
		WHERE IdMember = '%d'`,
		IdMember))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Member{}, er.QueryError
	}
	m, _, err := member.ScanResult(result, true)
	if err != nil {
		return Member{}, err
	}
	return m, nil
}

func (member Member) SearchModels(c *fiber.Ctx, offset int) ([]Member, string, error) {
	var searchKey string
	// si estamos en deleteMode que el searchKey lo saque del header, ya que no se lo voy a mandar por el form
	// asi cuando elimino un miembro se quedan los miembros que busque antes menos el que elimine
	if c.Get("deleteMode") == "true" {
		searchKey = c.Get("searchKey")
	} else {
		// sino se lo mando por el form normalmente
		searchKey = c.FormValue("search-key")
	}
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT 
		*
		FROM MemberTable WHERE 
		Name LIKE '%%%s%%' OR LastName LIKE '%%%s%%' OR DNI LIKE '%%%s%%' 
		ORDER BY LastName ASC LIMIT 15 OFFSET %d`,
		searchKey, searchKey, searchKey, offset))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, "", er.QueryError
	}
	_, mm, err := member.ScanResult(result, false)
	if err != nil {
		return nil, "", err
	}
	return mm, searchKey, nil
}

func (member Member) ValidateFields(c *fiber.Ctx) (map[string]string, error) {
	errorMap := map[string]string{}
	var valid bool
	var err string

	if valid, err = ValidateName(c); !valid {
		errorMap["name"] = err
	}
	if valid, err = ValidateLastName(c); !valid {
		errorMap["lastName"] = err
	}
	if valid, err = ValidateDNI(c); !valid {
		errorMap["dni"] = err
	}
	if valid, err = ValidateBirthday(c); !valid {
		errorMap["birthday"] = err
	}
	if valid, err = ValidateGender(c); !valid {
		errorMap["gender"] = err
	}
	if valid, err = ValidateMaritalStatus(c); !valid {
		errorMap["maritalStatus"] = err
	}
	if valid, err = ValidatePhone(c); !valid {
		errorMap["phone"] = err
	}
	if valid, err = ValidateEmail(c); !valid {
		errorMap["email"] = err
	}
	if valid, err = ValidateAddress(c); !valid {
		errorMap["address"] = err
	}
	if valid, err = ValidatePostalCode(c); !valid {
		errorMap["postalCode"] = err
	}
	if valid, err = ValidateDistrict(c); !valid {
		errorMap["district"] = err
	}
	if valid, err = ValidateMemberNumber(c); !valid {
		errorMap["memberNumber"] = err
	}
	if valid, err = ValidateCUIL(c); !valid {
		errorMap["cuil"] = err
	}
	valid, err, err2 := ValidateIdEnterprise(c)
	if err2 != nil {
		return nil, err2
	}
	if !valid {
		errorMap["idEnterprise"] = err
	}
	if valid, err = ValidateCategory(c); !valid {
		errorMap["category"] = err
	}
	if valid, err = ValidateEntryDate(c); !valid {
		errorMap["entryDate"] = err
	}

	if len(errorMap) > 1 {

		return errorMap, er.ValidationError
	}
	return errorMap, nil
}

func (member Member) GetTotalRows(c *fiber.Ctx) (int, error) {
	var totalRows int
	var searchKey string
	// si estamos en deleteMode que el searchKey lo saque del header, ya que no se lo voy a mandar por el form
	// asi cuando elimino un miembro se quedan los miembros que busque antes menos el que elimine
	if c.Get("deleteMode") == "true" {
		searchKey = c.Get("searchKey")
	} else {
		// sino se lo mando por el form normalmente
		searchKey = c.FormValue("search-key")
	}
	// no puedo hacer asi porque sino los afiliados con id enterprise null no aparecen
	// row := database.DB.QueryRow(fmt.Sprintf(`
	// 	SELECT COUNT(*) FROM MemberTable M INNER JOIN EnterpriseTable E ON M.IdEnterprise = E.IdEnterprise
	// 	WHERE
	// 	M.Name LIKE '%%%s%%' OR M.LastName LIKE '%%%s%%' OR M.DNI LIKE '%%%s%%'
	// 	OR E.Name LIKE '%%%s%%' OR E.EnterpriseNumber LIKE '%%%s%%'`,
	// 	searchKey, searchKey, searchKey, searchKey, searchKey))
	// row.Scan copia el numero de fila en la variable count
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM MemberTable WHERE
		Name LIKE '%%%s%%' OR LastName LIKE '%%%s%%' OR DNI LIKE '%%%s%%'`,
		searchKey, searchKey, searchKey))
	err := row.Scan(&totalRows)
	if err != nil {
		er.ScanError.Msg = err.Error()
		return 0, er.ScanError
	}
	return totalRows, nil
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

func (member Member) GetAllModels() ([]Member, error) {
	result, err := database.DB.Query("SELECT * FROM MemberTable")
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, er.QueryError
	}
	_, mm, err := member.ScanResult(result, false)
	if err != nil {
		return nil, err
	}
	return mm, nil
}

// func CheckIdEnterprise(tempIdEnterprise sql.NullInt16) int {
// 	if tempIdEnterprise.Valid {
// 		return int(tempIdEnterprise.Int16)
// 	} else {
// 		return 0
// 	}
// }

func (member Member) ScanResult(result *sql.Rows, onlyOne bool) (Member, []Member, error) {
	var m Member
	var mm []Member
	// var tempIdEnterprise sql.NullInt16
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
			// &tempIdEnterprise,
			&m.IdEnterprise,
			&m.Category,
			&m.EntryDate,
			&m.CreatedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			er.ScanError.Msg = err.Error()
			return Member{}, nil, er.ScanError
		}
		// formateo las fechas en formato argentino
		m.Birthday = FormatToDDMMYYYY(m.Birthday)
		m.EntryDate = FormatToDDMMYYYY(m.EntryDate)
		// m.IdEnterprise = CheckIdEnterprise(tempIdEnterprise)
		if !onlyOne {
			mm = append(mm, m)
		}
	}
	result.Close()
	return m, mm, nil
}

func (member Member) CheckDeleted(idMember int) (bool, error) {
	var totalRows int
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM MemberTable 
		WHERE IdMember = '%d'`, idMember))
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
