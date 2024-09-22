package controller

import (
	"database/sql"

	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

func parserCaller[M models.TypeModel](parser i.ModelParser[M], c *fiber.Ctx) M {
	return parser.ParseModel(c)
}

func insertModelCaller[M models.TypeModel](m i.IModel[M]) M {
	return m.InsertModel()
}

func deleteModelCaller[M models.TypeModel](m i.IModel[M]) {
	m.DeleteModel()
}

func editModelCaller[M models.TypeModel](m i.IModel[M]) {
	m.EditModel()
}

func getIdModelCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) int {
	return m.GetIdModel(c)
}

func searchOneModelByIdCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) M {
	return m.SearchOneModelById(c)
}

func searchModelsCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx, offset int) ([]M, string) {
	return m.SearchModels(c, offset)
}

func validateFieldsCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) map[string]string {
	return m.ValidateFields(c)
}

func getTotalRowsCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) int {
	return m.GetTotalRows(c)
}

func getFiberMapCaller[M models.TypeModel](m i.IModel[M], models []M, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return m.GetFiberMap(models, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)
}

func getAllModelsCaller[M models.TypeModel](m i.IModel[M]) []M {
	return m.GetAllModels()
}

func scanResultCaller[M models.TypeModel](m i.IModel[M], result *sql.Rows) M {
	return m.ScanResult(result)
}
