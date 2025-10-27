package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/svc"
	"go-cli-mgt/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
)

func ListUsersProfileHandler(c *fiber.Ctx) error {
	logger.Logger.Debug("Handler request List Users")

	users, err := svc.GetListProfile()
	if err != nil {
		logger.Logger.Error("Cannot get list user: ", err)
		response.InternalError(c, "cannot get list user")
		return err
	}

	logger.Logger.Info("Get list user success, total length: ", len(users))
	response.Write(c, users)
	return nil
}
