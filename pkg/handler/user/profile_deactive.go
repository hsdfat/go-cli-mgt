package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_error"
	userService "go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/response"
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
	err = userService.DisableProfile(user.Username, username)
	if err != nil {
		if errors.Is(err, models_error.ErrDisableUser) {
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
