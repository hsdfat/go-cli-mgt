package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	userService "go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/response"
)

func ProfileCreateHandler(c *fiber.Ctx) error {
	var user models_api.User
	err := c.BodyParser(&user)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	logger.Logger.Info("Handler create user with username: ", user.Username)
	err = userService.CreateProfile(user)
	if err != nil {
		if errors.Is(err, errors.New("username already existed")) {
			logger.Logger.Info("username already existed")
			response.BadRequest(c, "username already existed")
			return err
		}
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	logger.Logger.Info("Create User success with username: ", user.Username)
	return nil
}
