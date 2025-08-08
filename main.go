package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	pe "github.com/LucasBastino/app-sindicato/src/permissions"
	"github.com/LucasBastino/app-sindicato/src/router"
	"github.com/Masterminds/sprig"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {

	// KeyAuthApp.Api(
	// 	"appsindicato", // App name
	// 	"qPDHOjEGZl",   // Account ID
	// 	"1.0",          // Application version. Used for automatic downloads see video here https://www.youtube.com/watch?v=kW195PLCBKs
	// 	"",             // Token Path (PUT "null" OR LEAVE BLANK IF YOU DO NOT WANT TO USE THE TOKEN VALIDATION SYSTEM! MUST DISABLE VIA APP SETTINGS)
	// )

	// KeyAuthApp.Login("sindicato", "Contraparasindicato123,,")
	// // fmt.Println(KeyAuthApp.Var("pago"))
	// // if KeyAuthApp.Var("pago") == "no" {
	// // 	fmt.Println("no permitido")
	// // }

	// ticker := time.NewTicker(5 * time.Second)

	// go func() {
	// 	for t := range ticker.C {
	// 		log.Printf("Tick at: %v\n", t.UTC())
	// 		// do something
	// 		fmt.Println(KeyAuthApp.Var("pago"))
	// 	}
	// }()

	// engine config
	engine := html.New("./src/views", ".html")
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
	app.Static("/static", "./src/static")

	// Loading .env file
	godotenv.Load()

	// go func(){
	// 	for {
	// 		fmt.Println("go routina ejecutandose")
	// 		resp, err := http.Get("http://localhost:8080/validateAuth")
	// 		if err!=nil{
	// 			fmt.Println("error")
	// 		}
	// 		fmt.Println(resp.Body)
	// 		time.Sleep(5*time.Second)
	// 		defer resp.Body.Close()
	// 	}
	// }()



	go func(){
		for{
			result, err := database.DB.Query("SELECT Authorized FROM ClientTable WHERE Name ='Sindicato'")
			if err!=nil{
				pe.Authorized = false
			}
			for result.Next(){
				err = result.Scan(&pe.Authorized)
				if err!=nil{
					pe.Authorized = false
				}
			}
			result.Close()
			fmt.Println(pe.Authorized)
			time.Sleep(5*time.Hour)
		}	
	}()


	
	// Listen
	// cambiar el port este por uno mas profesional ↓
	log.Fatal(app.Listen(":8080"))
}
