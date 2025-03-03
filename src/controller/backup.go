package controller

import (
	"fmt"

	"github.com/JamesStewy/go-mysqldump"
	"github.com/LucasBastino/app-sindicato/src/database"
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
		fmt.Println("Error registering database:", err)
		return c.Render("backup", fiber.Map{"done": false, "error": err})
	}

	// Dump database to file.
	resultFilename, err := dumper.Dump()
	if err != nil {
		fmt.Println("Error dumping:", err)
		return c.Render("backup", fiber.Map{"done": false, "error": err})
	}

	fmt.Printf("File is saved to %s", resultFilename)

	// // Close dumper, connected database and file stream.
	dumper.Close()

	database.CreateConnection()
	return c.Render("backup", fiber.Map{"done": true})
}
