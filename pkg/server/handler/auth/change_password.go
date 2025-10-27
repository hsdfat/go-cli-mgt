package auth

import (
	"go-cli-mgt/pkg/logger"
	models_api "go-cli-mgt/pkg/models/api"
	"go-cli-mgt/pkg/svc"
	"go-cli-mgt/pkg/utils/bcrypt"
	"go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
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
	user, err := svc.GetProfileByUsername(userChangePassword.Username)
	if err != nil {
		logger.Logger.Error("Error get user: ", err)
		response.InternalError(c, "Error get user")
		svc.SaveHistoryCommandFailure(historyCommand)
		return err
	}
	user.Password = bcrypt.Encode(userChangePassword.Username + userChangePassword.NewPassword)
	svc.UpdatePassword(user)
	logger.Logger.Info("Update password success for user: ", userChangePassword.Username)
	return nil
}
