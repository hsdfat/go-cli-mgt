package role

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
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
	roleList, err := svc.GetListRole()
	if err != nil {
		logger.Logger.Error("Error get list role: ", err)
		response.InternalError(c, "Error get list role")
		svc.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	svc.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, roleList)
	return nil
}
