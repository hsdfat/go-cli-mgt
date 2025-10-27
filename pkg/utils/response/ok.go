package response

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Ok(c *fiber.Ctx) {
	var resp response.RespSuccess

	resp.Status = true
	resp.Code = http.StatusOK
	resp.Message = "Ok"

	err := c.Status(http.StatusOK).JSON(resp)
	if err != nil {
		logger.Logger.Error("Cannot send response")
		return
	}
}
