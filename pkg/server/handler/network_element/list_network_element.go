package network_element

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

func ListNetworkElementHandler(c *fiber.Ctx) error {
	username := c.Get("username")
	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Get list ne",
		NeName:   "",
		Mode:     "cli-config",
	}
	logger.Logger.Info("Handler get ne list")
	neList, err := svc.GetListNetworkElement()
	if err != nil {
		logger.Logger.Error("Error get list ne: ", err)
		response.InternalError(c, "Error get list ne")
		svc.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	svc.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, neList)
	return nil
}
