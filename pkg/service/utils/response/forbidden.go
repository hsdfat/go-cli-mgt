package response

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_response"
	"net/http"
)

func Forbidden(c *fiber.Ctx, message string) {
	if len(message) == 0 {
		message = "forbidden"
	}

	err := c.Status(http.StatusForbidden).JSON(models_response.RespError{
		Status:  false,
		Code:    http.StatusForbidden,
		Message: message,
		Error:   "forbidden",
	})
	if err != nil {
		logger.Logger.Error("Cannot send response with message: ", message)
		return
	}
}
