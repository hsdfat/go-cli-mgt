package role

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	historyService "go-cli-mgt/pkg/service/history"
	roleService "go-cli-mgt/pkg/service/role"
	"go-cli-mgt/pkg/service/utils/response"
)

func ListRoleHandler(c *fiber.Ctx) error {
	username := c.Get("username")
	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Get list role",
		NeName:   "",
		Mode:     "cli-config",
	}
	logger.Logger.Info("Handler get role list")
	roleList, err := roleService.GetListRole()
	if err != nil {
		logger.Logger.Error("Error get list role: ", err)
		response.InternalError(c, "Error get list role")
		historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	historyService.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, roleList)
	return nil
}
