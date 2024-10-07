package network_element

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	historyService "go-cli-mgt/pkg/service/history"
	neService "go-cli-mgt/pkg/service/network_elements"
	"go-cli-mgt/pkg/service/utils/response"
)

func DeleteHandler(c *fiber.Ctx) error {
	var ne models_api.NeData
	err := c.BodyParser(&ne)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")

	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Create network element " + ne.Name + " namespace " + ne.Namespace,
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Infof("Handler create ne %s namespace %s", ne.Name, ne.Namespace)
	err = neService.CreateNetworkElement(&ne)
	if err != nil {
		logger.Logger.Error("Error create ne: ", err)
		response.InternalError(c, "Error create ne")
		historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	historyService.SaveHistoryCommandSuccess(historyCommand)
	logger.Logger.Infof("Create ne %s namespace %s success", ne.Name, ne.Namespace)
	return nil
}
