package main

import (
	"github.com/gofiber/fiber/v2"
	logger "github.com/kmesiab/go-klogger"
	"gorm.io/gorm"

	"github.com/kmesiab/equilibria-api/controllers"
	"github.com/kmesiab/equilibria-api/lib/nrclex"
	"github.com/kmesiab/equilibria-api/utils"
)

func main() {

	var err error
	var db *gorm.DB
	var cfg *utils.EQConfig

	logger.Logf("API Server: started.")

	// Load configuration
	if cfg, err = utils.GetConfig(); err != nil {
		logger.Logf("API Server: failed to load config: %s", err.Error())

		return
	}

	logger.Logf("API Server: loaded config.")

	if db, err = utils.InitDB(cfg); err != nil {
		logger.Logf("API Server: failed to initialize database: %s", err.Error())

		return
	}

	// Create the app
	app := fiber.New()

	emotionController := controllers.EmotionsController{
		MaxLimit:  50,
		MaxOffset: 10,

		NrcLexService: nrclex.NewService(nrclex.NewRepository(db)),
	}

	AddGetRoute("/emotions/:userid", emotionController.NrcLex)

	// Add the middleware, which
	// MUST come before loading routes.
	LoadMiddleware(app)

	// Load the routes
	LoadRoutes(app)

	// Start the server
	err = app.Listen(":443")

	if err != nil {

		logger.Logf("API Server: failed to start: %s", err.Error())

		return
	}
}
