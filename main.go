package main

import (
	"e-com/config"
	database "e-com/pkg/databases"
	"e-com/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		log.Println("Error .env")
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "POST,GET,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	db := database.ConnentDB()

	routes.SetupRoutes(app, db)

	config.InitGoogleOAuth()

	if err := app.Listen(":8080"); err != nil {
		log.Println(err)
	}
}
