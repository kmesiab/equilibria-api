package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/kmesiab/equilibria-api/middleware"
)

func LoadMiddleware(app *fiber.App) {

	app.Use(middleware.Authorize)

	EnableCORS(app)
}

func EnableCORS(app *fiber.App) {

	fmt.Println("Enabling CORS")

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://my-eq.com, http://localhost:4200",
		AllowMethods:     "GET,POST,HEAD,PUT,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
	}))

}
