package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.ctx) error {
	return c.SendString("Hello, Fetch!")
}