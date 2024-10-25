package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/AngelinaLe021/fetchReward_Assessment/receipt/models"
	"github.com/AngelinaLe021/fetchReward_Assessment/receipt/repo"
	"github.com/google/uuid"
)


func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, Fetch!")
}

var storeRecord = models.NewRecord()

func ProcessReceipt(c *fiber.Ctx) error {
	purchase := new(models.Purchases)
	if err := c.BodyParser(purchase); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	pointsEarned, err := repo.PointsEarned(*purchase)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	newID := uuid.New()
	response := models.Response{
		ID: newID.String(),
		Points: pointsEarned,
	}

	storeRecord.Add(response)

	return c.Status(200).JSON(fiber.Map{
		"id": response.ID,
	})
}

func Point(storage *models.Storage) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		record, check := storeRecord.Fetch(id)
	
		if !check {
			return c.Status(404).JSON(fiber.Map{ "error": "Receipt not found." })
		}

		return c.Status(200).JSON(fiber.Map{
			"points": record.Points,
		})
	}
}

