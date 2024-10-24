package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/AngelinaLe021/fetchReward_Assessment/receipt/handlers"
	"github.com/AngelinaLe021/fetchReward_Assessment/receipt/models"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	storage := &models.Storage{
		Store: make(map[string]models.Response),
	}
	app.Post("/receipts/process", handlers.ProcessReceipt)
	app.Get("/receipts/:id/points", handlers.Point(storage))
}