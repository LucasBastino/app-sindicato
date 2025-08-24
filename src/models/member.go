package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/gofiber/fiber/v2"
)

type Member struct {
	IdMember      int		`json:"idmember"`
	Name          string	`json:"name"`
	LastName      string	`json:"lastname"`
	DNI           string	`json:"dni"`
	Birthday      string	`json:"birthday"`
	Gender        string	`json:"gender"`
	MaritalStatus string	`json:"maritalstatus"`
	Phone         string	`json:"phone"`
	Email         string	`json:"email"`
	Address       string	`json:"address"`
	PostalCode    string	`json:"postalcode"`
	District      string	`json:"district"`
	MemberNumber  string	`json:"membernumber"`
	Affiliated    bool		`json:"affiliated"`
	CUIL          string	`json:"cuil"`
	IdEnterprise  int		`json:"identerprise"`
	Category      string	`json:"category"`
	EntryDate     string	`json:"entrydate"`
	Observations  string	`json:"observations"`
	CreatedAt     time.Time	`json:"createdat"`
	UpdatedAt     time.Time	`json:"updatedat"`
}

func (member Member) InsertModel() (Member, customError.CustomError) {
	// formateo la fecha nac para que empiece con el año
	member.Birthday = FormatToYYYYMMDD(member.Birthday)
	member.EntryDate = FormatToYYYYMMDD(member.EntryDate)
	insert, err := database.DB.Query(`
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
		Affiliated,
		CUIL,
		IdEnterprise,
		Category,
		EntryDate,
		Observations) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
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
		member.Affiliated,
		member.CUIL,
		member.IdEnterprise,
		member.Category,
		member.EntryDate,
		member.Observations)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Member{}, customError.QueryError
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT * FROM MemberTable 
		WHERE IdMember = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Member{}, customError.QueryError
	}
	m, _, customErr := member.ScanResult(result, true)
	if err != nil {
		return Member{}, customErr
	}
	return m, customError.CustomError{}
}

func (member Member) DeleteModel() customError.CustomError {
	delete, err := database.DB.Query(`
		DELETE FROM MemberTable 
		WHERE IdMember = ?`,
		member.IdMember)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return customError.QueryError
	}
	defer delete.Close()
	return customError.CustomError{}
}

func (member Member) UpdateModel() (Member, customError.CustomError) {
	// formateo la fecha nac para que empiece con el año
	member.Birthday = FormatToYYYYMMDD(member.Birthday)
	member.EntryDate = FormatToYYYYMMDD(member.EntryDate)
	update, err := database.DB.Query(`
		UPDATE MemberTable
		SET
		Name = ?,
		LastName = ?,
		DNI = ?,
		Birthday = ?,
		Gender = ?,
		MaritalStatus = ?,
		Phone = ?,
		Email = ?,
		Address = ?,
		PostalCode = ?,
		District = ?,
		MemberNumber = ?,
		Affiliated = ?,
		CUIL = ?,
		IdEnterprise = ?,
		Category = ?,
		EntryDate = ?,
		Observations = ?
		WHERE IdMember = ?`,
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
		member.Affiliated,
		member.CUIL,
		member.IdEnterprise,
		member.Category,
		member.EntryDate,
		member.Observations,
		member.IdMember)
	if err != nil {
		fmt.Println("error updateando member")
		customError.QueryError.Msg = err.Error()
		return Member{}, customError.QueryError
	}
	update.Close()
	result, err := database.DB.Query(`
		SELECT * FROM MemberTable 
		WHERE IdMember = ?`, member.IdMember)

	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Member{}, customError.QueryError
	}
	m, _, customErr := member.ScanResult(result, true)
	if (customErr != customError.CustomError{}) {
		return Member{}, customErr
	}
	return m, customError.CustomError{}
}

func (member Member) GetIdModel(c *fiber.Ctx) (int, customError.CustomError) {
	// params := struct {
	// 	IdMember int `params:"IdMember"`
	// }{}

	// c.ParamsParser(&params)
	// return params.IdMember

	// hacerlos asi a partir de ahora
	idMember, err := c.ParamsInt("IdMember")
	if err != nil {
		customError.ParamsError.Msg = err.Error()
		return 0, customError.ParamsError
	}
	return idMember, customError.CustomError{}
}

func (member Member) SearchOneModelById(c *fiber.Ctx) (Member, customError.CustomError) {
	IdMember, customErr := member.GetIdModel(c)
	if (customErr != customError.CustomError{}) {
		return Member{}, customErr
	}
	result, err := database.DB.Query("SELECT * FROM MemberTable WHERE IdMember = ?", IdMember)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Member{}, customError.QueryError
	}
	m, _, customErr := member.ScanResult(result, true)
	if (customErr != customError.CustomError{}) {
		return Member{}, customErr
	}
	return m, customError.CustomError{}
}

func (member Member) SearchModels(c *fiber.Ctx, offset int) ([]Member, string, customError.CustomError) {
	var searchKey string
	// si estamos en deleteMode que el searchKey lo saque del header, ya que no se lo voy a mandar por el form
	// asi cuando elimino un miembro se quedan los miembros que busque antes menos el que elimine
	if c.Get("deleteMode") == "true" {
		searchKey = c.Get("searchKey")
	} else {
		// sino se lo mando por el form normalmente
		searchKey = c.FormValue("search-key")
	}
	result, err := database.DB.Query(`
		SELECT 
		*
		FROM MemberTable WHERE 
		Name LIKE concat('%', ?, '%') OR LastName LIKE concat('%', ?, '%') OR DNI LIKE concat('%', ?, '%') 
		ORDER BY LastName ASC LIMIT 15 OFFSET ?`,
		searchKey, searchKey, searchKey, offset)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return nil, "", customError.QueryError
	}
	_, mm, customErr := member.ScanResult(result, false)
	if err != nil {
		return nil, "", customErr
	}
	return mm, searchKey, customError.CustomError{}
}

func (member Member) ValidateFields(c *fiber.Ctx) customError.CustomError {
	validateFunctions := []func(*fiber.Ctx) error{
		ValidateName,
		ValidateLastName,
		ValidateDNI,
		ValidateBirthday,
		ValidateGender,
		ValidateMaritalStatus,
		ValidatePhone,
		ValidateEmail,
		ValidateAddress,
		ValidatePostalCode,
		ValidateDistrict,
		ValidateMemberNumber,
		ValidateAffiliated,
		ValidateCUIL,
		ValidateIdEnterprise,
		ValidateCategory,
		ValidateEntryDate,
		ValidateObservations,
	}

	

	for _, vF := range validateFunctions {
		if err := vF(c); err != nil {
			customError.ValidationError.Msg = err.Error()
			return customError.ValidationError
		} else {
			continue
		}
	}
	return customError.CustomError{}
}

func (member Member) GetTotalRows(c *fiber.Ctx) (int, customError.CustomError) {
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
	row := database.DB.QueryRow("SELECT COUNT(*) FROM MemberTable WHERE Name LIKE concat('%', ?, '%') OR LastName LIKE concat('%', ?, '%') OR DNI LIKE concat('%', ?, '%')", searchKey, searchKey, searchKey)
	err := row.Scan(&totalRows)
	if err != nil {
		customError.ScanError.Msg = err.Error()
		return 0, customError.ScanError
	}
	return totalRows, customError.CustomError{}
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

func (member Member) GetAllModels() ([]Member, customError.CustomError) {
	result, err := database.DB.Query("SELECT * FROM MemberTable")
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return nil, customError.QueryError
	}
	_, mm, customErr := member.ScanResult(result, false)
	if (customErr != customError.CustomError{}) {
		return nil, customErr
	}
	return mm, customError.CustomError{}
}

// func CheckIdEnterprise(tempIdEnterprise sql.NullInt16) int {
// 	if tempIdEnterprise.Valid {
// 		return int(tempIdEnterprise.Int16)
// 	} else {
// 		return 0
// 	}
// }

func (member Member) ScanResult(result *sql.Rows, onlyOne bool) (Member, []Member, customError.CustomError) {
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
			&m.Affiliated,
			&m.CUIL,
			// &tempIdEnterprise,
			&m.IdEnterprise,
			&m.Category,
			&m.EntryDate,
			&m.Observations,
			&m.CreatedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			customError.ScanError.Msg = err.Error()
			return Member{}, nil, customError.ScanError
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
	return m, mm, customError.CustomError{}
}

func (member Member) CheckDeleted(idMember int) (bool, customError.CustomError) {
	var totalRows int
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM MemberTable 
		WHERE IdMember = ?`, idMember)
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		customError.ScanError.Msg = err.Error()
		return false, customError.ScanError
	}
	if totalRows == 0 {
		return true, customError.CustomError{}
	} else {
		return false, customError.CustomError{}
	}
}
