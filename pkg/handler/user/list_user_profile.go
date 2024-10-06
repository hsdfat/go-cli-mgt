package user

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	userService "go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/response"
)

func ListUsersProfileHandler(c *fiber.Ctx) error {
	logger.Logger.Debugln("Handler request List Users")

	users, err := userService.GetListProfile()
	if err != nil {
		logger.Logger.Error("Cannot get list user: ", err)
		response.InternalError(c, "cannot get list user")
		return err
	}

	logger.Logger.Info("Get list user success, total length: ", len(users))
	response.Write(c, users)
	return nil
}
