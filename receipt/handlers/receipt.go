package handlers

import (
	"strconv"
	"strings"
	"fmt"
	"time"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/AngelinaLe021/fetchReward_Assessment/receipt/models"
	"github.com/google/uuid"
)


func converterStringToFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func PointsEarned(receipt models.Purchases) (int, error) {
	points := 0

	// One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points ++
		}
	}

	// 50 points if the total is a round number
	total, err := converterStringToFloat(receipt.Total)
	if err != nil {
		return 0, fmt.Errorf("invalid total format: %v", err)
	}
	if total == float64(int(total)) {
		points += 50
	}
	
	// 25 if the total is a multiple of 0.25
	if int(total*100)%25 == 0 {
		points += 25
	}

	// 5 point for every 2 items on the receipt
	points += 5 * (len(receipt.Items) / 2)

	// Trimmed Length of ShortDescription
	for _, item := range receipt.Items {
		trimmed := strings.TrimSpace(item.ShortDescription)

		price, err := converterStringToFloat(item.Price)
		if err != nil {
			return 0, fmt.Errorf("invalid price format for item: %v", err)
		}

		if len(trimmed) % 3 == 0{
			points += int((price * 0.2) + 0.999999)
		}
	}

	// 6 points for purchasing on odd days
	dateParts := strings.Split(receipt.PurchaseDate, "-")
	if len(dateParts) != 3 { 
		println("Invalid date format") 
	}
	dayStr := dateParts[2]
	day, _ := strconv.Atoi(dayStr)
	if day % 2 != 0 {
		points += 6
	}

	// 10 points for purchasin between 14:00-16:00
	if purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
		if purchaseTime.After(time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)) &&
		purchaseTime.Before(time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)) {
			points += 10
		}
	
	}
	return points, nil
}

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, Fetch!")
}


func ProcessReceipt(c *fiber.Ctx) error {
	purchase := new(models.Purchases)
	if err := c.BodyParser(purchase); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	pointsEarned, _ := PointsEarned(*purchase)
	newID := uuid.New()
	response := models.Response{
		ID: newID.String(),
		Points: pointsEarned,
	}

	record := models.NewRecord()
	record.Add(response)

	return c.Status(200).JSON(fiber.Map{
		"id": response.ID,
		"points": response.Points,
	})
}

func Point(c *fiber.Ctx) error {
	id := c.Params("id")

	record, check := models.Fetch(id)
	
	if !check {
		return c.Status(404).JSON(fiber.Map{ "error": "Receipt not found." })
	}

	return c.Status(200).JSON(fiber.Map{
		"points": record.Points,
	})
}

