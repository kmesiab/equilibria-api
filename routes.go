package main

import (
	"github.com/gofiber/fiber/v2"
)

// Controllers

// Routes

var getRoutes = map[string]fiber.Handler{}

var postRoutes = map[string]fiber.Handler{}

var putRoutes = map[string]fiber.Handler{}

var patchRoutes = map[string]fiber.Handler{}

var deleteRoutes = map[string]fiber.Handler{}

func AddGetRoute(path string, handler fiber.Handler) {
	getRoutes[path] = handler
}

func LoadRoutes(app *fiber.App) {

	for path, handler := range getRoutes {
		app.Get(path, handler)
	}

	for path, handler := range postRoutes {
		app.Post(path, handler)
	}

	for path, handler := range putRoutes {
		app.Put(path, handler)
	}

	for path, handler := range patchRoutes {
		app.Patch(path, handler)
	}

	for path, handler := range deleteRoutes {
		app.Delete(path, handler)
	}
}
