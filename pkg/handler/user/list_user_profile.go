package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/store/repository"

	"github.com/gofiber/fiber/v2"
)

func ListUsersProfileHandler(c *fiber.Ctx) error {
	logger.Logger.Debugln("List Users")
	store := repository.GetSingleton()

	users, err := store.ListUsers()
	if err != nil {
		logger.Logger.Errorln("Error listing users:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(users)
}
