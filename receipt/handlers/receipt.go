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

func isLeapYear(year int) bool {
	return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
}

func PointsEarned(receipt models.Purchases) (int, error) {
	points := 0

	// Check if there's item in the receipt
	if len(receipt.Items) == 0 {
		return 0,
		fmt.Errorf("No items found in the receipt.")
	}

	// One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points ++
		}
	}

	// 50 points if the total is a round number
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, 
		fmt.Errorf("invalid total format: %v", err)
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

	// Processing each item in the receipt w length % 3 == 0
	for _, item := range receipt.Items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0, 
			fmt.Errorf("invalid price format for item '%s': %v", item.ShortDescription, err)
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
	yrStr := dateParts[0]
	year, err := strconv.Atoi(yrStr)
	if err != nil {
		return 0,
		fmt.Errorf("Invalid year format: %s", receipt.PurchaseDate)
	}

	monthStr := dateParts[1]
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return 0,
		fmt.Errorf("Invalid month format: %s", receipt.PurchaseDate)
	}
	if month < 1 || month > 12 {
		return 0,
		fmt.Errorf("Invalid month", month)
	}

	dayStr := dateParts[2]
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return 0,
		fmt.Errorf("Invalid date: %s", receipt.PurchaseDate)
	}
	var maxDay int
	switch month {
	case 1, 3, 5, 7, 8, 10, 12: maxDay = 31
	case 4, 6, 9, 11: maxDay = 30
	case 2:
		if isLeapYear(year) {
			maxDay = 29
		} else {
			maxDay = 28
		}

	default:
		return 0, fmt.Errorf("Invalid month:", month)
	}

	if day < 1 || day > maxDay {
		return 0,
		fmt.Errorf("Invalid date format: %s", receipt.PurchaseDate)
	}

	if day % 2 != 0 {
		points += 6
	}

	// 10 points for purchasin between 14:00-16:00
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0,
		fmt.Errorf("Invalid time format: expected HH:MM, got '%s'", receipt.PurchaseTime)
	}
	
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}
	
	return points, nil
}

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

	pointsEarned, err := PointsEarned(*purchase)
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

