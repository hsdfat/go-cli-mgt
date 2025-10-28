package user

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

func RoleAddHandler(c *fiber.Ctx) error {
	var userRole models_api.UserRole
	err := c.BodyParser(&userRole)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")
	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Add role for user",
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler add user role: ")
	err = svc.RoleUserAdd(userRole.UserId, userRole.RoleId)
	if err != nil {
		logger.Logger.Error("Error add user role: ", err)
		response.InternalError(c, "Error add user role")
		svc.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	svc.SaveHistoryCommandSuccess(historyCommand)
	return nil
}
