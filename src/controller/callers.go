package controller

import (
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	i "github.com/LucasBastino/app-sindicato/src/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
	"github.com/gofiber/fiber/v2"
)

func parserCaller[M models.TypeModel](parser i.ModelParser[M], c *fiber.Ctx) (M, customError.CustomError) {
	return parser.ParseModel(c)
}

func insertModelCaller[M models.TypeModel](m i.IModel[M]) (M, customError.CustomError) {
	return m.InsertModel()
}

func deleteModelCaller[M models.TypeModel](m i.IModel[M]) customError.CustomError {
	return m.DeleteModel()
}

func updateModelCaller[M models.TypeModel](m i.IModel[M]) (M, customError.CustomError) {
	return m.UpdateModel()
}

func getIdModelCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) (int, customError.CustomError) {
	return m.GetIdModel(c)
}

func searchOneModelByIdCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) (M, customError.CustomError) {
	return m.SearchOneModelById(c)
}

func searchModelsCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx, offset int) ([]M, string, customError.CustomError) {
	return m.SearchModels(c, offset)
}

func validateFieldsCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) customError.CustomError {
	return m.ValidateFields(c)
}

func getTotalRowsCaller[M models.TypeModel](m i.IModel[M], c *fiber.Ctx) (int, customError.CustomError) {
	return m.GetTotalRows(c)
}

func getFiberMapCaller[M models.TypeModel](m i.IModel[M], models []M, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return m.GetFiberMap(models, searchKey, currentPage, someBefore, someAfter, totalPages, totalPagesArray)
}

func getAllModelsCaller[M models.TypeModel](m i.IModel[M]) ([]M, customError.CustomError) {
	return m.GetAllModels()
}

func checkDeletedCaller[M models.TypeModel](m i.IModel[M], idModel int) (bool, customError.CustomError) {
	return m.CheckDeleted(idModel)
}
