package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_error"
	historyService "go-cli-mgt/pkg/service/history"
	userService "go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/bcrypt"
	"go-cli-mgt/pkg/service/utils/response"
)

func ChangePasswordHandler(c *fiber.Ctx) error {
	var userChangePassword models_api.ChangePassWord
	err := c.BodyParser(&userChangePassword)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")
	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Change password for user " + userChangePassword.Username,
		NeName:   "",
		Mode:     "cli-config",
	}
	logger.Logger.Info("Handler Change password for user: ", userChangePassword.Username)
	user, err := userService.GetProfileByUsername(userChangePassword.Username)
	if err != nil {
		if errors.Is(err, models_error.ErrNotFoundUser) {
			logger.Logger.Info("user does not existed")
			response.BadRequest(c, "user does not existed")
			historyService.SaveHistoryCommandFailure(historyCommand)
			return err
		}
		logger.Logger.Error("Error get user: ", err)
		response.InternalError(c, "Error get user")
		historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}
	user.Password = bcrypt.Encode(userChangePassword.Username + userChangePassword.NewPassword)
	userService.UpdatePassword(user)
	logger.Logger.Info("Update password success for user: ", userChangePassword.Username)
	return nil
}
