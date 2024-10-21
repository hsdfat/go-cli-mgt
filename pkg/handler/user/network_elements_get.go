package user

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	historyService "go-cli-mgt/pkg/service/history"
	userService "go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/response"
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

	neList, err := userService.GetListNeByUser(username)
	if err != nil {
		logger.Logger.Error("Cannot get list ne form database, err: ", err)
		historyService.SaveHistoryCommandFailure(historyCommand)
		response.BadRequest(c, "cannot get list ne of user")
		return err
	}

	logger.Logger.Info("Get list Ne of User success")
	historyService.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, neList)
	return nil
}
