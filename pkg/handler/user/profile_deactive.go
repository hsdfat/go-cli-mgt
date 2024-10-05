package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	userService "go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/response"
	"go-cli-mgt/pkg/store/postgres"
)

func ProfileDeactivateHandler(c *fiber.Ctx) error {
	var user models_api.User
	err := c.BodyParser(&user)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	logger.Logger.Info("Handler disable user with username: ", user.Username)
	err = userService.DisableProfile(user.Username)
	if err != nil {
		if errors.Is(err, postgres.ErrDisableUser) {
			logger.Logger.Info("username already disable")
			response.BadRequest(c, "username already disable")
			return err
		}
		logger.Logger.Error("Error Disable user: ", err)
		response.InternalError(c, "Error Disable user")
		return err
	}
	return nil
}
