package role

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	historyService "go-cli-mgt/pkg/service/history"
	roleService "go-cli-mgt/pkg/service/role"
	"go-cli-mgt/pkg/service/utils/response"
)

func DeleteHandler(c *fiber.Ctx) error {
	var role models_api.Role
	err := c.BodyParser(&role)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")

	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Delete role " + role.RoleName,
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler delete role ", role.RoleName)
	err = roleService.DeleteRole(&role)
	if err != nil {
		logger.Logger.Error("Error delete role: ", err)
		response.InternalError(c, "Error delete role")
		historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	historyService.SaveHistoryCommandSuccess(historyCommand)
	logger.Logger.Infof("Delete role %s success ", role.RoleName)
	return nil
}
