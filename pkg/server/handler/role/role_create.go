package role

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
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

	roleGet, err := svc.GetRoleByName(role.RoleName)
	if roleGet == nil {
		logger.Logger.Info("Handler create role ", role.RoleName)
		err = svc.CreateRole(&role)
		if err != nil {
			logger.Logger.Error("Error create role: ", err)
			response.InternalError(c, "Error create role")
			svc.SaveHistoryCommandFailure(historyCommand)
			return err
		}
		svc.SaveHistoryCommandSuccess(historyCommand)
		logger.Logger.Infof("Create role %s success ", role.RoleName)
	} else {
		logger.Logger.Info("Handler update role ", role.RoleName)
		roleGet.Priority = role.Priority
		roleGet.Description = role.Description
		svc.UpdateRole(roleGet)
		svc.SaveHistoryCommandSuccess(historyCommand)
		logger.Logger.Infof("Update role %s success ", role.RoleName)
	}

	return nil
}
