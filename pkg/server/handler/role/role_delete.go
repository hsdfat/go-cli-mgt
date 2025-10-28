package role

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
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
	err = svc.DeleteRole(&role)
	if err != nil {
		logger.Logger.Error("Error delete role: ", err)
		response.InternalError(c, "Error delete role")
		svc.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	svc.SaveHistoryCommandSuccess(historyCommand)
	logger.Logger.Infof("Delete role %s success ", role.RoleName)
	return nil
}
