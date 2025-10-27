package response

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func InternalError(c *fiber.Ctx, message string) {
	if len(message) == 0 {
		message = "Internal Server Error"
	}

	resp := response.RespError{
		Status:  false,
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Error:   message,
	}

	err := c.Status(http.StatusInternalServerError).JSON(resp)
	if err != nil {
		logger.Logger.Error("Cannot send response with message: ", message)
		return
	}
}
