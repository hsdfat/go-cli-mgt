package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/bcrypt"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/token"
)

type RequestUser struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type ChangePassWord struct {
	Username    string `json:"username"`
	NewPassword string `json:"new_password"`
}

type AuthHandler struct {
	authService    *svc.AuthService
	userService    *svc.UserService
	historyService *svc.HistoryService
}

func NewAuthHandlerHandler() *AuthHandler {
	return &AuthHandler{
		authService:    svc.NewAuthService(),
		userService:    svc.NewUserService(),
		historyService: svc.NewHistoryService(),
	}
}

func (h *AuthHandler) LoginHandler(c *fiber.Ctx) error {
	var userLogin RequestUser
	err := c.BodyParser(&userLogin)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	logger.Logger.Info("Handler login for user: ", userLogin.Username)

	checkPass, err, userId := h.authService.Login(userLogin.Username, userLogin.Password)
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

	roles, err := h.authService.GetRole(userId)
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

func (h *AuthHandler) ChangePasswordHandler(c *fiber.Ctx) error {
	var userChangePassword ChangePassWord
	err := c.BodyParser(&userChangePassword)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")

	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Change password for user " + userChangePassword.Username,
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler Change password for user: ", userChangePassword.Username)

	user, err := h.userService.GetProfileByUsername(userChangePassword.Username)
	if err != nil {
		logger.Logger.Error("Error get user: ", err)
		response.InternalError(c, "Error get user")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	user.Password = bcrypt.Encode(userChangePassword.Username + userChangePassword.NewPassword)
	h.userService.UpdatePassword(user)

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	logger.Logger.Info("Update password success for user: ", userChangePassword.Username)
	return nil
}
