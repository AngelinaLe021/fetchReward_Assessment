package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/AngelinaLe021/fetchReward_Assessment/receipt/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}