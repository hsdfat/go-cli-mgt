package response

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Created(c *fiber.Ctx) {
	var resp response.RespSuccess

	resp.Status = true
	resp.Code = http.StatusCreated
	resp.Message = "Created"

	err := c.Status(http.StatusCreated).JSON(resp)
	if err != nil {
		logger.Logger.Error("Cannot send response")
		return
	}
}
