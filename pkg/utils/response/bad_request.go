package response

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func BadRequest(c *fiber.Ctx, message string) {
	if len(message) == 0 {
		message = "Bad Request"
	}

	err := c.Status(http.StatusBadRequest).JSON(response.RespError{
		Status:  false,
		Code:    http.StatusBadRequest,
		Message: message,
		Error:   "Bad Request",
	})
	if err != nil {
		logger.Logger.Error("Cannot send response with message: ", message)
		return
	}
}
