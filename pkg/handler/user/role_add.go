package user

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	historyService "go-cli-mgt/pkg/service/history"
	userService "go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/response"
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
	err = userService.RoleUserAdd(userRole.UserId, userRole.RoleId)
	if err != nil {
		logger.Logger.Error("Error add user role: ", err)
		response.InternalError(c, "Error add user role")
		historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	historyService.SaveHistoryCommandSuccess(historyCommand)
	return nil
}
