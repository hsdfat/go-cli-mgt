package user

import (
	"go-cli-mgt/pkg/logger"
	models_api "go-cli-mgt/pkg/models/api"
	"go-cli-mgt/pkg/svc"
	"go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

func NetworkElementsListHandler(c *fiber.Ctx) error {
	var user models_api.User
	err := c.BodyParser(&user)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}
	var username string
	if user.Username == "" {
		username = c.Get("username")
	} else {
		username = user.Username
	}

	historyCommand := &models_api.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Add permission for user",
		NeName:   "",
		Mode:     "cli-config",
	}

	neList, err := svc.GetListNeByUser(username)
	if err != nil {
		logger.Logger.Error("Cannot get list ne form database, err: ", err)
		svc.SaveHistoryCommandFailure(historyCommand)
		response.BadRequest(c, "cannot get list ne of user")
		return err
	}

	logger.Logger.Info("Get list Ne of User success")
	svc.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, neList)
	return nil
}
