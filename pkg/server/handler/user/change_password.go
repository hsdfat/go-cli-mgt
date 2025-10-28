package user

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/bcrypt"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"

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
	if !bcrypt.Matches(userChangePassword.Username+userChangePassword.OldPassword, user.Password) {
		logger.Logger.Info("Wrong old password user")
		response.BadRequest(c, "old password incorrect")
		svc.SaveHistoryCommandFailure(historyCommand)
		return nil
	}

	user.Password = bcrypt.Encode(userChangePassword.Username + userChangePassword.NewPassword)
	svc.UpdatePassword(user)
	logger.Logger.Info("Update password success for user: ", userChangePassword.Username)
	return nil
}
