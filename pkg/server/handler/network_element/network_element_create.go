package network_element

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

func CreateOrUpdateHandler(c *fiber.Ctx) error {
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
	err = svc.CreateNetworkElement(&ne)
	if err != nil {
		logger.Logger.Error("Error create ne: ", err)
		response.InternalError(c, "Error create ne")
		svc.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	svc.SaveHistoryCommandSuccess(historyCommand)
	logger.Logger.Infof("Create ne %s namespace %s success", ne.Name, ne.Namespace)
	return nil
}
