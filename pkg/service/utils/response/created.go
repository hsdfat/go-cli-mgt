package response

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_response"
	"net/http"
)

func Created(c *fiber.Ctx) {
	var resp models_response.RespSuccess

	resp.Status = true
	resp.Code = http.StatusCreated
	resp.Message = "Created"

	err := c.Status(http.StatusCreated).JSON(resp)
	if err != nil {
		logger.Logger.Error("Cannot send response")
		return
	}
}
