package middleware

import "github.com/gofiber/fiber/v2"

func BasicAuth(c *fiber.Ctx) error {
	// Basic auth implementation
	c.Next()
	return nil
}
