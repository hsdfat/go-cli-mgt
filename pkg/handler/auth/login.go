package auth

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/service/auth"
	"go-cli-mgt/pkg/service/utils/response"
	"go-cli-mgt/pkg/service/utils/token"
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
	checkPass, err, userId := auth.Login(userLogin.Username, userLogin.Password)
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

	roles, err := auth.GetRole(userId)
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
