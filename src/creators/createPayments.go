package creators

import (
	"fmt"
	"math/rand/v2"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePayments(c *fiber.Ctx) error {
	months := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	pp := []models.Payment{}
	p := models.Payment{}

	for i := range 50 {
		for _, month := range months {
			p.Month = month
			p.Year = "2024"
			p.Commentary = ""
			p.IdEnterprise = i + 1
			random := rand.IntN(10)
			if random == 1 {
				p.Status = "IMPAGO"
			} else {
				p.Status = "PAGO"
			}
			p.Amount = rand.IntN(50000) + 50000
			if p.Status == "IMPAGO" {

				p.PaymentDate = ""
			} else {
				p.PaymentDate = fmt.Sprintf("05/%s/%s", p.Month, p.Year)
			}
			pp = append(pp, p)
		}
	}

	for _, p := range pp {
		insert, err := database.DB.Query(`
		INSERT INTO PaymentTable(
		Month,
		Year,
		Status,
		Amount,
		PaymentDate,
		Commentary,
		IdEnterprise
		)
		VALUES ('?','?','?','?','?', '?', '?')`,
			p.Month, p.Year, p.Status, p.Amount, p.PaymentDate, p.Commentary, p.IdEnterprise)
		if err != nil {
			fmt.Println("error inserting payment in db")
			panic(err)
		}
		insert.Close()
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ps added"})
}
