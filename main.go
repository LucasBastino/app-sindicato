package main

import (
	"log"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// Initializing and config app
	app := fiber.New(fiber.Config{
		Views: html.New("./src/views", ".html"),
	})

	// Middlewares
	app.Use(logger.New())

	// Routes
	router.RegisterRoutes(app)

	// Database connection
	database.CreateConnection()

	// Serve static files
	app.Static("/static", "./src/static")

	// Listen
	log.Fatal(app.Listen(":8085"))

	// muxer := http.NewServeMux()
	// svr := &http.Server{
	// 	Addr:    ":8085",
	// 	Handler: muxer,
	// }
	// c := api.Controller{}
	// c.DB = connection.CreateConnection()

	// // buscar como obtener el current PATH, y cambiarle las \ por / o eso de que depende el OS lo cambia solo
	// fileServerStatic := http.FileServer(http.Dir("src/static")) // con o sin barra al final es lo mismo
	// // si o si es muxer.Handle, con http.Handle no funciona
	// muxer.Handle("/static/", http.StripPrefix("/static/", fileServerStatic))
	// c.RegisterRoutes(muxer)

	// fmt.Println("Listen on port 8085")

	// svr.ListenAndServe()
}
