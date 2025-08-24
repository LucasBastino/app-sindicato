package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/gofiber/fiber/v2"
)

type Payment struct {
	IdPayment    int		`json:"idpayment"`
	Month        string		`json:"month"`
	Year         string		`json:"year"`
	Status       bool		`json:"status"`
	Amount       int		`json:"amount"`
	PaymentDate  string		`json:"paymentdate"`
	Observations string		`json:"observations"`
	IdEnterprise int		`json:"identerprise"`
	CreatedAt    time.Time	`json:"createdat"`
	UpdatedAt    time.Time	`json:"updatedat"`
}

func (payment Payment) InsertModel() (Payment, customError.CustomError) {
	insert, err := database.DB.Query(`
		INSERT INTO PaymentTable 
		(Month,
		Year,
		Status, 
		Amount, 
		PaymentDate, 
		Observations,
		IdEnterprise)
		VALUES (?,?,?,?,?, ?, ?)`,
		payment.Month,
		payment.Year,
		payment.Status,
		payment.Amount,
		payment.PaymentDate,
		payment.Observations,
		payment.IdEnterprise)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Payment{}, customError.QueryError
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT *
		FROM PaymentTable
		WHERE IdPayment = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Payment{}, customError.QueryError
	}
	p, _, customErr := payment.ScanResult(result, true)
	if (customErr != customError.CustomError{}) {
		return Payment{}, customErr
	}
	return p, customError.CustomError{}
}

func (payment Payment) DeleteModel() customError.CustomError {
	delete, err := database.DB.Query(`
		DELETE FROM PaymentTable
		WHERE IdPayment = ?`, payment.IdPayment)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return customError.QueryError
	}
	defer delete.Close()
	return customError.CustomError{}
}

func (payment Payment) UpdateModel() (Payment, customError.CustomError) {
	update, err := database.DB.Query(`
		UPDATE PaymentTable 
		SET 
		Month = ?, 
		Year = ?,
		Status = ?, 
		Amount = ?, 
		PaymentDate = ?, 
		Observations = ? 
		WHERE IdPayment = ?`,
		payment.Month,
		payment.Year,
		payment.Status,
		payment.Amount,
		payment.PaymentDate,
		payment.Observations,
		payment.IdPayment)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Payment{}, customError.QueryError
	}
	update.Close()
	result, err := database.DB.Query("SELECT * FROM PaymentTable WHERE IdPayment = ?", payment.IdPayment)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Payment{}, customError.QueryError
	}
	p, _, customErr := payment.ScanResult(result, true)
	if (customErr != customError.CustomError{}) {
		return Payment{}, customErr
	}
	return p, customError.CustomError{}
}

func (payment Payment) GetIdModel(c *fiber.Ctx) (int, customError.CustomError) {
	params := struct {
		IdPayment int `params:"IdPayment"`
	}{}

	err := c.ParamsParser(&params)
	if err != nil {
		customError.ParamsError.Msg = err.Error()
		return 0, customError.ParamsError
	}

	return params.IdPayment, customError.CustomError{}
}

func (payment Payment) SearchOneModelById(c *fiber.Ctx) (Payment, customError.CustomError) {
	IdPayment, customErr := payment.GetIdModel(c)
	if (customErr != customError.CustomError{}) {
		return Payment{}, customErr
	}
	result, err := database.DB.Query(`
		SELECT
		*
		FROM PaymentTable
		WHERE IdPayment = ?`, IdPayment)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Payment{}, customError.QueryError
	}

	p, _, customErr := payment.ScanResult(result, true)
	if (customErr != customError.CustomError{}) {
		return Payment{}, customErr
	}
	return p, customError.CustomError{}

}

func (payment Payment) SearchModels(c *fiber.Ctx, year int) ([]Payment, string, customError.CustomError) {
	IdEnterprise, customErr := Enterprise{}.GetIdModel(c)
	if (customErr != customError.CustomError{}) {
		return nil, "", customErr
	}
	yearStr := strconv.Itoa(year)
	result, err := database.DB.Query(`
	SELECT
	*
	FROM PaymentTable 
	WHERE Year = ? AND IdEnterprise = ?
	ORDER BY Month DESC
	`, yearStr, IdEnterprise)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return nil, "", customError.QueryError
	}
	_, pp, customErr := payment.ScanResult(result, false)
	if (customErr != customError.CustomError{}) {
		return nil, "", customErr
	}
	return pp, "", customError.CustomError{}
}

func (payment Payment) ValidateFields(c *fiber.Ctx) customError.CustomError {
	validateFunctions := []func(*fiber.Ctx) error{
		ValidatePayment,
		ValidateStatus,
		ValidatePaymentAmount,
		ValidatePaymentDate,
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

func (payment Payment) GetTotalRows(c *fiber.Ctx) (int, customError.CustomError) {
	return 0, customError.CustomError{}
}

func (payment Payment) GetFiberMap(Payments []Payment, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return nil
}

func (payment Payment) GetAllModels() ([]Payment, customError.CustomError) {
	return nil, customError.CustomError{}
}

func (payment Payment) ScanResult(result *sql.Rows, onlyOne bool) (Payment, []Payment, customError.CustomError) {
	var p Payment
	var pp []Payment
	for result.Next() {
		err := result.Scan(
			&p.IdPayment,
			&p.Month,
			&p.Year,
			&p.Status,
			&p.Amount,
			&p.PaymentDate,
			&p.Observations,
			&p.IdEnterprise,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			customError.ScanError.Msg = err.Error()
			return Payment{}, nil, customError.ScanError
		}
		if !onlyOne {
			pp = append(pp, p)
		}
	}
	result.Close()
	return p, pp, customError.CustomError{}
}

func (payment Payment) CheckDeleted(idPayment int) (bool, customError.CustomError) {
	var totalRows int
	// row := database.DB.QueryRow(fmt.Sprintf(`
	// 	SELECT COUNT(*) FROM PaymentTable
	// 	WHERE IdPayment = '%d'`, p.IdPayment))
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM PaymentTable 
		WHERE IdPayment = ?`, idPayment)
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
