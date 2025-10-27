package history

import (
	"go-cli-mgt/pkg/logger"
	models_api "go-cli-mgt/pkg/models/api"
	"go-cli-mgt/pkg/svc"
	"go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

func SaveHistoryHandler(c *fiber.Ctx) error {
	var history models_api.History
	err := c.BodyParser(&history)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}
	logger.Logger.Info("Saving history command: ", history.Command)
	if history.Result {
		svc.SaveHistoryCommandSuccess(&history)
	} else {
		svc.SaveHistoryCommandFailure(&history)
	}
	return nil
}
