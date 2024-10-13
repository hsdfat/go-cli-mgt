package user

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	historyService "go-cli-mgt/pkg/service/history"
	userService "go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/response"
)

func NetworkElementDeleteHandler(c *fiber.Ctx) error {
	var userNe models_api.UserNe
	err := c.BodyParser(&userNe)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")
	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Delete permission for user",
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler add user ne: ")
	err = userService.NeUserDelete(userNe.UserId, userNe.NeId)
	if err != nil {
		logger.Logger.Error("Error delete user ne: ", err)
		response.InternalError(c, "Error delete user ne")
		historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	historyService.SaveHistoryCommandSuccess(historyCommand)
	return nil
}
