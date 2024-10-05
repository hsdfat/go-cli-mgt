package response

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_response"
	"net/http"
)

func Unauthorized(c *fiber.Ctx) {
	// Set Response Data to HTTP
	err := c.Status(http.StatusUnauthorized).JSON(models_response.RespError{
		Status:  false,
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
		Error:   "Unauthorized",
	})
	if err != nil {
		logger.Logger.Error("Cannot send response with message")
		return
	}
}
