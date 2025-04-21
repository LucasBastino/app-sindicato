package interfaces

import (
	"database/sql"

	er "github.com/LucasBastino/app-sindicato/src/errors"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

type IModel[M models.TypeModel] interface {
	InsertModel() (M, er.CustomError)
	DeleteModel() er.CustomErrorr
	UpdateModel() (M, er.CustomError)
	GetIdModel(*fiber.Ctx) (int, er.CustomError)
	SearchOneModelById(*fiber.Ctx) (M, er.CustomError)
	SearchModels(*fiber.Ctx, int) ([]M, string, er.CustomError)
	ValidateFields(*fiber.Ctx) error
	GetTotalRows(*fiber.Ctx) (int, error)
	GetFiberMap([]M, string, int, int, int, int, []int) fiber.Map
	GetAllModels() ([]M, error)
	ScanResult(*sql.Rows, bool) (M, []M, error)
	CheckDeleted(int) (bool, error)
}
