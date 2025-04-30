package interfaces

import (
	"database/sql"

	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

type IModel[M models.TypeModel] interface {
	InsertModel() (M, error)
	DeleteModel() error
	UpdateModel() (M, error)
	GetIdModel(*fiber.Ctx) (int, error)
	SearchOneModelById(*fiber.Ctx) (M, error)
	SearchModels(*fiber.Ctx, int) ([]M, string, error)
	ValidateFields(*fiber.Ctx) error
	GetTotalRows(*fiber.Ctx) (int, error)
	GetFiberMap([]M, string, int, int, int, int, []int) fiber.Map
	GetAllModels() ([]M, error)
	ScanResult(*sql.Rows, bool) (M, []M, error)
	CheckDeleted(int) (bool, error)
}
