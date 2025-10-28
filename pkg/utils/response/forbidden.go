package response

import (
	"net/http"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	"github.com/hsdfat/go-cli-mgt/pkg/models/response"

	"github.com/gofiber/fiber/v2"
)

func Forbidden(c *fiber.Ctx, message string) {
	if len(message) == 0 {
		message = "forbidden"
	}

	err := c.Status(http.StatusForbidden).JSON(response.RespError{
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
