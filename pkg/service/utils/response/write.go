package response

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_response"
	"net/http"
)

func Write(c *fiber.Ctx, data interface{}) {
	err := c.Status(http.StatusOK).JSON(models_response.RespSuccess{
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
