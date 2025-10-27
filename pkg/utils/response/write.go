package response

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Write(c *fiber.Ctx, data interface{}) {
	err := c.Status(http.StatusOK).JSON(response.RespSuccess{
		Status:   true,
		Code:     http.StatusOK,
		Message:  "success",
		DataResp: data,
	})
	if err != nil {
		logger.Logger.Error("Cannot send response with message")
		return
	}
}
