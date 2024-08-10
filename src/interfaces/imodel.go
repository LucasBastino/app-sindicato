package interfaces

import (
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

type IModel[M models.TypeModel] interface {
	InsertModel()
	DeleteModel()
	EditModel()
	GetIdModel(*fiber.Ctx) int
	SearchOneModelById(*fiber.Ctx) M
	SearchModels(*fiber.Ctx) []M
	ValidateFields(*fiber.Ctx) map[string]string
	CreateTemplateData(M, []M, string, map[string]string) models.TemplateData
}
