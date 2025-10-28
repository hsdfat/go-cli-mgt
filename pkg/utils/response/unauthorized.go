package response

import (
	"net/http"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	"github.com/hsdfat/go-cli-mgt/pkg/models/response"

	"github.com/gofiber/fiber/v2"
)

func Unauthorized(c *fiber.Ctx) {
	// Set Response Data to HTTP
	err := c.Status(http.StatusUnauthorized).JSON(response.RespError{
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
