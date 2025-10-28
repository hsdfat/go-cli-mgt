package user

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

func ProfileDeactivateHandler(c *fiber.Ctx) error {
	var user models_api.User
	err := c.BodyParser(&user)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}
	username := c.Get("username")
	logger.Logger.Info("Handler disable user with username: ", user.Username)

	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Deactivate user " + user.Username,
		NeName:   "",
		Mode:     "cli-config",
	}

	err = svc.DisableProfile(user.Username, username)
	if err != nil {
		svc.SaveHistoryCommandFailure(historyCommand)
		logger.Logger.Error("Error Disable user: ", err)
		response.InternalError(c, "Error Disable user")
		return err
	}

	logger.Logger.Info("Deactivate user success, username: ", user.Username)
	svc.SaveHistoryCommandSuccess(historyCommand)
	return nil
}
