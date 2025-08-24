package interfaces

import (
	"database/sql"

	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

type IModel[M models.TypeModel] interface {
	InsertModel() (M, customError.CustomError)
	DeleteModel() customError.CustomError
	UpdateModel() (M, customError.CustomError)
	GetIdModel(*fiber.Ctx) (int, customError.CustomError)
	SearchOneModelById(*fiber.Ctx) (M, customError.CustomError)
	SearchModels(*fiber.Ctx, int) ([]M, string, customError.CustomError)
	ValidateFields(*fiber.Ctx) customError.CustomError
	GetTotalRows(*fiber.Ctx) (int, customError.CustomError)
	GetFiberMap([]M, string, int, int, int, int, []int) fiber.Map
	GetAllModels() ([]M, customError.CustomError)
	ScanResult(*sql.Rows, bool) (M, []M, customError.CustomError)
	CheckDeleted(int) (bool, customError.CustomError)
}
