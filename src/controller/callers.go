package controller

import (
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

func parserCaller[M models.TypeModel](parser i.ModelParser[M], c *fiber.Ctx) (M, error) {
	return parser.ParseModel(c)
}

func insertModelCaller[M models.TypeModel](m i.IModel[M]) (M, error) {
	return m.InsertModel()
}

func deleteModelCaller[M models.TypeModel](m i.IModel[M]) error {
	return m.DeleteModel()
}

func updateModelCaller[M models.TypeModel](m i.IModel[M]) (M, error) {
	return m.UpdateModel()
}

func getIdModelCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) (int, error) {
	return m.GetIdModel(c)
}

func searchOneModelByIdCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) (M, error) {
	return m.SearchOneModelById(c)
}

func searchModelsCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx, offset int) ([]M, string, error) {
	return m.SearchModels(c, offset)
}

func validateFieldsCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) error {
	return m.ValidateFields(c)
}

func getTotalRowsCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) (int, error) {
	return m.GetTotalRows(c)
}

func getFiberMapCaller[M models.TypeModel](m i.IModel[M], models []M, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return m.GetFiberMap(models, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)
}

func getAllModelsCaller[M models.TypeModel](m i.IModel[M]) ([]M, error) {
	return m.GetAllModels()
}

func checkDeletedCaller[M models.TypeModel](m i.IModel[M], idModel int) (bool, error) {
	return m.CheckDeleted(idModel)
}
