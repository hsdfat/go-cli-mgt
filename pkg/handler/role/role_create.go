package role

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	historyService "go-cli-mgt/pkg/service/history"
	roleService "go-cli-mgt/pkg/service/role"
	"go-cli-mgt/pkg/service/utils/response"
)

func CreateOrUpdateHandler(c *fiber.Ctx) error {
	var role models_api.Role
	err := c.BodyParser(&role)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}
	logger.Logger.Debug("Description: ", role.Description)

	username := c.Get("username")

	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Create role " + role.RoleName,
		NeName:   "",
		Mode:     "cli-config",
	}

	roleGet, err := roleService.GetRoleByName(role.RoleName)
	if roleGet == nil {
		logger.Logger.Info("Handler create role ", role.RoleName)
		err = roleService.CreateRole(&role)
		if err != nil {
			logger.Logger.Error("Error create role: ", err)
			response.InternalError(c, "Error create role")
			historyService.SaveHistoryCommandFailure(historyCommand)
			return err
		}
		historyService.SaveHistoryCommandSuccess(historyCommand)
		logger.Logger.Infof("Create role %s success ", role.RoleName)
	} else {
		logger.Logger.Info("Handler update role ", role.RoleName)
		roleGet.Priority = role.Priority
		roleGet.Description = role.Description
		roleService.UpdateRole(roleGet)
		historyService.SaveHistoryCommandSuccess(historyCommand)
		logger.Logger.Infof("Update role %s success ", role.RoleName)
	}

	return nil
}
