package interfaces

import (
	"database/sql"

	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

type IModel[M models.TypeModel] interface {
	InsertModel() M
	DeleteModel()
	UpdateModel()
	GetIdModel(*fiber.Ctx) int
	SearchOneModelById(*fiber.Ctx) M
	SearchModels(*fiber.Ctx, int) ([]M, string)
	ValidateFields(*fiber.Ctx) map[string]string
	GetTotalRows(*fiber.Ctx) int
	GetFiberMap([]M, string, int, int, int, int, []int) fiber.Map
	GetAllModels() []M
	ScanResult(*sql.Rows, bool) (M, []M)
	CheckDeleted(int) bool
}
