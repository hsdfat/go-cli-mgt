package history

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	historyService "go-cli-mgt/pkg/service/history"
	"go-cli-mgt/pkg/service/utils/response"
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
		historyService.SaveHistoryCommandSuccess(&history)
	} else {
		historyService.SaveHistoryCommandFailure(&history)
	}
	return nil
}
