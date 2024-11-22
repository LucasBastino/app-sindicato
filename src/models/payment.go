package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

type Payment struct {
	IdPayment    int
	Month        string
	Year         string
	Status       string
	Amount       int
	PaymentDate  string
	Commentary   string
	IdEnterprise int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (payment Payment) InsertModel() Payment {
	insert, err := database.DB.Query(fmt.Sprintf(`
		INSERT INTO PaymentTable 
		(Month,
		Year,
		Status, 
		Amount, 
		PaymentDate, 
		Commentary,
		IdEnterprise)
		VALUES ('%s','%s','%s','%d','%s', '%s', '%d')`,
		payment.Month,
		payment.Year,
		payment.Status,
		payment.Amount,
		payment.PaymentDate,
		payment.Commentary,
		payment.IdEnterprise))
	if err != nil {
		fmt.Println("error insertando Payment en la DB")
		panic(err)
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT *
		FROM PaymentTable
		WHERE IdPayment = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		fmt.Print(err)
	}
	p, _ := payment.ScanResult(result, true)
	return p
}

func (payment Payment) DeleteModel() {
	delete, err := database.DB.Query(fmt.Sprintf(`
		DELETE FROM PaymentTable
		WHERE IdPayment = '%d'`, payment.IdPayment))
	if err != nil {
		fmt.Println("error deleting Payment")
		panic(err)
	}
	defer delete.Close()

}

func (payment Payment) UpdateModel() {
	update, err := database.DB.Query(fmt.Sprintf(`
		UPDATE PaymentTable 
		SET 
		Month = '%s', 
		Year = '%s',
		Status = '%s', 
		Amount = '%d', 
		PaymentDate = '%s', 
		Commentary = '%s' 
		WHERE IdPayment = '%d'`,
		payment.Month,
		payment.Year,
		payment.Status,
		payment.Amount,
		payment.PaymentDate,
		payment.Commentary,
		payment.IdPayment))
	if err != nil {
		fmt.Println("error updating Payment")
		panic(err)
	}
	defer update.Close()
}

func (payment Payment) GetIdModel(c *fiber.Ctx) int {
	params := struct {
		IdPayment int `params:"IdPayment"`
	}{}

	c.ParamsParser(&params)

	return params.IdPayment
}

func (payment Payment) SearchOneModelById(c *fiber.Ctx) Payment {
	IdPayment := payment.GetIdModel(c)
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT
		*
		FROM PaymentTable
		WHERE IdPayment = '%d'`, IdPayment))
	if err != nil {
		fmt.Println("error searching Payment by id")
		panic(err)
	}

	p, _ := payment.ScanResult(result, true)
	return p

}

func (payment Payment) SearchModels(c *fiber.Ctx, year int) ([]Payment, string) {
	IdEnterprise := Enterprise{}.GetIdModel(c)
	yearStr := strconv.Itoa(year)
	result, err := database.DB.Query(fmt.Sprintf(`
	SELECT
	*
	FROM PaymentTable 
	WHERE Year = '%s' AND IdEnterprise = '%d'
	ORDER BY Month ASC
	`, yearStr, IdEnterprise))
	if err != nil {
		fmt.Println("error searching payments per year")
		panic(err)
	}
	_, pp := payment.ScanResult(result, false)
	return pp, ""
}

func (payment Payment) ValidateFields(c *fiber.Ctx) map[string]string {
	errorMap := map[string]string{}
	var valid bool
	var err string
	if valid, err = validatePayment(c); !valid {
		errorMap["month"] = err
	}
	if valid, err = validatePayment(c); !valid {
		errorMap["year"] = err
	}
	if valid, err = validateStatus(c); !valid {
		errorMap["status"] = err
	}
	if valid, err = validatePaymentAmount(c); !valid {
		errorMap["amount"] = err
	}
	if valid, err = validatePaymentDate(c); !valid {
		errorMap["payment-date"] = err
	}
	if valid, err = validateCommentary(c); !valid {
		errorMap["commentary"] = err
	}
	return errorMap
}

func (payment Payment) GetTotalRows(c *fiber.Ctx) int {
	return 0
}

func (payment Payment) GetFiberMap(Payments []Payment, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return nil
}

func (payment Payment) GetAllModels() []Payment {
	return nil
}

func (payment Payment) ScanResult(result *sql.Rows, onlyOne bool) (Payment, []Payment) {
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
			&p.Commentary,
			&p.IdEnterprise,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			fmt.Println("error scanning Payment")
			panic(err)
		}
		if !onlyOne {
			pp = append(pp, p)
		}
	}
	result.Close()
	return p, pp
}

func (payment Payment) CheckDeleted(idPayment int) bool {
	var totalRows int
	// row := database.DB.QueryRow(fmt.Sprintf(`
	// 	SELECT COUNT(*) FROM PaymentTable
	// 	WHERE IdPayment = '%d'`, p.IdPayment))
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM PaymentTable 
		WHERE IdPayment = '%d'`, idPayment))
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
