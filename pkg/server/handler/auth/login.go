package auth

import (
	"go-cli-mgt/pkg/logger"
	models_api "go-cli-mgt/pkg/models/api"
	"go-cli-mgt/pkg/svc"
	"go-cli-mgt/pkg/utils/response"
	"go-cli-mgt/pkg/utils/token"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var userLogin models_api.RequestUser
	err := c.BodyParser(&userLogin)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	logger.Logger.Info("Handler login for user: ", userLogin.Username)
	checkPass, err, userId := svc.Login(userLogin.Username, userLogin.Password)
	if err != nil {
		logger.Logger.Error("Cannot check password: ", err)
		response.InternalError(c, "Cannot check password")
		return err
	}
	if checkPass == false {
		logger.Logger.Error("Wrong username or password with username: ", userLogin.Username)
		response.Unauthorized(c)
		return err
	}

	roles, err := svc.GetRole(userId)
	if err != nil {
		logger.Logger.Error("Cannot get role from user: ", err)
		response.InternalError(c, "Cannot get role from user")
		return err
	}

	tokenStr, err := token.CreateToken(userLogin.Username, roles)
	logger.Logger.Info("Login success for user: ", userLogin.Username)
	response.LoginSuccess(c, tokenStr)
	return nil
}
