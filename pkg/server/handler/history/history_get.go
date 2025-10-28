package history

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

func GetHistoryHandler(c *fiber.Ctx) error {
	var historyReq models_api.HistoryRequest
	err := c.BodyParser(&historyReq)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")
	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Get history command mode " + historyReq.Mode,
		NeName:   "",
		Mode:     "cli-config",
	}
	logger.Logger.Info("Handler get history list")
	histories, err := svc.GetListHistoryByMode(historyReq.Mode, historyReq.Limit)
	if err != nil {
		logger.Logger.Error("Cannot get history list, err: ", err)
		response.BadRequest(c, "cannot get history list")
		svc.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	svc.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, histories)
	logger.Logger.Info("Get history list success")
	return nil
}
