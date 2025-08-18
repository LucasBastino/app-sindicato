package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	pe "github.com/LucasBastino/app-sindicato/src/permissions"
	"github.com/LucasBastino/app-sindicato/src/router"
	"github.com/Masterminds/sprig"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

//go:embed src/views/*
var viewFiles embed.FS

//go:embed src/static/*
var staticFiles embed.FS

func embedfsSub(fsys embed.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}
	return sub
}

func main() {

	// engine config
	// engine := html.New("./src/views", ".html")
	// engine := html.NewFileSystem(http.FS(viewFiles), ".html")
	engine := html.NewFileSystem(http.FS(embedfsSub(viewFiles, "src/views")), ".html")

	// sprig es un paquete con funciones para el template
	engine.AddFuncMap(sprig.FuncMap())

	// Initializing and config app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middlewares
	app.Use(logger.New())

	// Routes
	router.RegisterRoutes(app)

	// Database connection
	database.CreateConnection()

	// Serve static files
	// app.Static("/static", "./src/static")
	// app.StaticFS("/static", http.FS(staticFiles))
	app.Use("/static", adaptor.HTTPHandler(http.StripPrefix("/static", http.FileServer(http.FS(embedfsSub(staticFiles, "src/static"))))))

	// Loading .env file
	godotenv.Load()

	go func() {
		for {
			result, err := database.AuthDB.Query("SELECT Authorized FROM ClientTable WHERE Name ='Sindicato'")
			if err != nil {
				pe.Authorized = false
			}
			for result.Next() {
				err = result.Scan(&pe.Authorized)
				if err != nil {
					pe.Authorized = false
				}
			}
			result.Close()
			fmt.Println(pe.Authorized)
			time.Sleep(5 * time.Hour)
		}
	}()
	// 192.168.100.2
	// Listen
	// cambiar el port este por uno mas profesional ↓
	log.Fatal(app.Listen("0.0.0.0:8080"))
}
