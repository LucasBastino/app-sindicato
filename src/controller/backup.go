package controller

import (
	"fmt"

	"github.com/JamesStewy/go-mysqldump"
	"github.com/LucasBastino/app-sindicato/src/config/logger"
	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/LucasBastino/app-sindicato/src/errors/errorHandler"
	"github.com/gofiber/fiber/v2"
)

func BackupDB(c *fiber.Ctx) error {
	DB := database.DB
	InfoDB := database.InfoDB
	dumpDir := "./src/database/dumps" // you should create this directory
	dumpFilenameFormat := fmt.Sprintf("%s-20060102T150405", InfoDB.DBName)

	// Register database with mysqldump.
	dumper, err := mysqldump.Register(DB, dumpDir, dumpFilenameFormat)
	if err != nil {
		customError.DatabaseError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.DatabaseError)
	}

	// Dump database to file.
	resultFilename, err := dumper.Dump()
	if err != nil {
		customError.FileError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.FileError)
	}

	logger.Log.Info(fmt.Sprintf("File is saved to %s", resultFilename))
	

	// // Close dumper, connected database and file stream.
	dumper.Close()

	database.CreateConnection()
	return c.Render("backup", fiber.Map{"done": true})
}
