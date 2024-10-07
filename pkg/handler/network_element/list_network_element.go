package network_element

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	historyService "go-cli-mgt/pkg/service/history"
	neService "go-cli-mgt/pkg/service/network_elements"
	"go-cli-mgt/pkg/service/utils/response"
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
	neList, err := neService.GetListNetworkElement()
	if err != nil {
		logger.Logger.Error("Error get list ne: ", err)
		response.InternalError(c, "Error get list ne")
		historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	historyService.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, neList)
	return nil
}
