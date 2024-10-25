package repo_test

import (
	"testing"

	"github.com/AngelinaLe021/fetchReward_Assessment/receipt/models"
	"github.com/AngelinaLe021/fetchReward_Assessment/receipt/handlers"
	"github.com/stretchr/testify/assert"
)

func TestValidInput(t *testing.T) {
	receipt := models.Purchases{
		Retailer: "SuperStore123", 
		Total: "120.00",
		PurchaseDate: "2024-10-24",
		PurchaseTime: "14:30",
		Items: []models.Item{
			{ShortDescription: "Apples", Price: "5.00"},
			{ShortDescription: "Bananas", Price: "5.00"},
			{ShortDescription: "Oranges", Price: "10.00"},
		},
	}

	points, err := handlers.PointsEarned(receipt)
	assert.NoError(t, err, "Expected no error for valid input")
	assert.Equal(t, 104, points, "Expected correct points calculation")
}

func TestInvalidTotal(t *testing.T) {
	receipt := models.Purchases{
		Retailer: "Invalid",
		Total: "abc",
		PurchaseDate: "2024-10-23",
		PurchaseTime: "14:20",
		Items: []models.Item{
			{ShortDescription: "Apples", Price: "5.00"},
			{ShortDescription: "Oranges", Price: "5.00"},
			{ShortDescription: "Bananas", Price: "2.00"},
		},
	}

	_, err := handlers.PointsEarned(receipt)
	assert.Error(t, err, "Expected an error for invalid total format")
	assert.Equal(t, "invalid total format: strconv.ParseFloat: parsing \"abc\": invalid syntax", err.Error())
}

func TestInvalidDate(t *testing.T) {
	receipt := models.Purchases{
		Retailer: "Invalid",
		Total: "120.0",
		PurchaseDate: "2024-10-34",
		PurchaseTime: "14:20",
		Items: []models.Item{
			{ShortDescription: "Apples", Price: "5.00"},
			{ShortDescription: "Oranges", Price: "5.00"},
			{ShortDescription: "Bananas", Price: "2.00"},
		},
	}

	_, err := handlers.PointsEarned(receipt)
	assert.Error(t, err, "Expected an error for invalid date")
	assert.Equal(t, "Invalid date format: 2024-10-34", err.Error())	
}

func TestInvalidTime(t *testing.T) {
	receipt := models.Purchases{
		Retailer: "Invalid",
		Total: "120.0",
		PurchaseDate: "2024-10-20",
		PurchaseTime: "26:00",
		Items: []models.Item{
			{ShortDescription: "Apples", Price: "5.00"},
			{ShortDescription: "Oranges", Price: "5.00"},
			{ShortDescription: "Bananas", Price: "2.00"},
		},
	}

	_, err := handlers.PointsEarned(receipt)
	assert.Error(t, err, "Expected an error for invalid time format")
	assert.Equal(t, "Invalid time format: expected HH:MM, got '26:00'", err.Error())	
}

func TestNoItems(t *testing.T) {
	receipt := models.Purchases{
		Retailer: "Invalid",
		Total: "120.0",
		PurchaseDate: "2024-10-20",
		PurchaseTime: "14:20",
		Items: []models.Item{},
	}

	_, err := handlers.PointsEarned(receipt)
	assert.Error(t, err, "Expected an error for missing items")
	assert.Equal(t, "No items found in the receipt.", err.Error())	
}

func TestRoundTotal(t *testing.T) {
	receipt := models.Purchases{
		Retailer: "RoundNum",
		Total: "100.00",
		PurchaseDate: "2024-10-20",
		PurchaseTime: "14:20",
		Items: []models.Item{
			{ShortDescription: "Apples", Price: "5.00"},
			{ShortDescription: "Oranges", Price: "5.00"},
		},
	}

	points, err := handlers.PointsEarned(receipt)
	assert.NoError(t, err, "Expected no error for valid round total")
	assert.Equal(t, 99, points, "Expected correct points calculation for round total")	
}

func TestMultipleOfQuarter(t *testing.T) {
	receipt := models.Purchases{
		Retailer: "Quarters",
		Total: "125.25",
		PurchaseDate: "2024-10-20",
		PurchaseTime: "14:20",
		Items: []models.Item{
			{ShortDescription: "Apples", Price: "5.00"},
			{ShortDescription: "Oranges", Price: "5.00"},
		},
	}

	points, err := handlers.PointsEarned(receipt)
	assert.NoError(t, err, "Expected no error for valid total")
	assert.Equal(t, 49, points, "Expected correct points calculation for total multiple of 0.25")	
}