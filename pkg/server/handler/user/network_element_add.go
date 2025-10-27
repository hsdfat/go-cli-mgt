package user

import (
	"go-cli-mgt/pkg/logger"
	models_api "go-cli-mgt/pkg/models/api"
	"go-cli-mgt/pkg/svc"
	"go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

func NetworkElementAddHandler(c *fiber.Ctx) error {
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
		Command:  "Add permission for user",
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler add user ne: ")
	err = svc.NeUserAdd(userNe.UserId, userNe.NeId)
	if err != nil {
		logger.Logger.Error("Error add user ne: ", err)
		response.InternalError(c, "Error add user ne")
		svc.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	svc.SaveHistoryCommandSuccess(historyCommand)
	return nil
}
