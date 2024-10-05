package response

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_response"
	"net/http"
)

func LoginSuccess(c *fiber.Ctx, tokenStr string) {
	var resp models_response.RespSuccess

	resp.Status = true
	resp.Code = http.StatusOK
	resp.Message = tokenStr

	err := c.Status(http.StatusOK).JSON(resp)
	if err != nil {
		logger.Logger.Error("Cannot send response")
		return
	}
}
