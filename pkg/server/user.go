package server

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"
)

// UserAPI represents the API model for user
type UserAPI struct {
	Id           uint   `json:"id,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	Active       bool   `json:"active,omitempty"`
	FailAttempt  int32  `json:"fail-attempt,omitempty"`
	Email        string `json:"email,omitempty"`
	DisableDate  uint64 `json:"disable-date,omitempty"`
	CreatedDate  uint64 `json:"created-date,omitempty"`
	DeActivateBy string `json:"deactivate-by"`
}

// UserRoleAPI represents user-role relationship
type UserRoleAPI struct {
	Id     uint `json:"id"`
	UserId uint `json:"user-id,omitempty"`
	RoleId uint `json:"role-id,omitempty"`
}

// UserNeAPI represents user-network element relationship
type UserNeAPI struct {
	Id     uint `json:"id"`
	UserId uint `json:"userId"`
	NeId   uint `json:"neId"`
}

// UserHandler struct contains all required services
type UserHandler struct {
	userService     *svc.UserService
	historyService  *svc.HistoryService
	userRoleService *svc.UserRoleService
	userNeService   *svc.UserNeService
}

// NewHandler initializes handler with services
func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService:     svc.NewUserService(),
		historyService:  svc.NewHistoryService(),
		userRoleService: svc.NewUserRoleService(),
		userNeService:   svc.NewUserNeService(),
	}
}

// apiToSvcUser converts API model to service model
func (h *UserHandler) apiToSvcUser(userAPI *UserAPI) *svc.User {
	return &svc.User{
		Id:           userAPI.Id,
		Username:     userAPI.Username,
		Password:     userAPI.Password,
		Active:       userAPI.Active,
		Email:        userAPI.Email,
		CreatedDate:  userAPI.CreatedDate,
		DisableDate:  userAPI.DisableDate,
		DeActivateBy: userAPI.DeActivateBy,
	}
}

// ProfileCreateHandler handles creating new user profile
func (h *UserHandler) ProfileCreateHandler(c *fiber.Ctx) error {
	var userAPI UserAPI
	err := c.BodyParser(&userAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")

	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Create user " + userAPI.Username + " password xxx",
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler create user with username: ", userAPI.Username)

	userSVC := h.apiToSvcUser(&userAPI)
	err = h.userService.CreateProfile(userSVC)
	if err != nil {
		if errors.Is(err, errors.New("username already existed")) {
			logger.Logger.Info("username already existed")
			response.BadRequest(c, "username already existed")
			h.historyService.SaveHistoryCommandFailure(historyCommand)
			return err
		}
		logger.Logger.Error("Error create user: ", err)
		response.InternalError(c, "Error create user")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	logger.Logger.Info("Create User success with username: ", userAPI.Username)
	return nil
}

// ProfileDeactivateHandler handles deactivating user profile
func (h *UserHandler) ProfileDeactivateHandler(c *fiber.Ctx) error {
	var userAPI UserAPI
	err := c.BodyParser(&userAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")
	logger.Logger.Info("Handler disable user with username: ", userAPI.Username)

	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Deactivate user " + userAPI.Username,
		NeName:   "",
		Mode:     "cli-config",
	}

	err = h.userService.DisableProfile(userAPI.Username, username)
	if err != nil {
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		logger.Logger.Error("Error Disable user: ", err)
		response.InternalError(c, "Error Disable user")
		return err
	}

	logger.Logger.Info("Deactivate user success, username: ", userAPI.Username)
	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	return nil
}

// ListUsersProfileHandler handles getting all user profiles
func (h *UserHandler) ListUsersProfileHandler(c *fiber.Ctx) error {
	logger.Logger.Debug("Handler request List Users")

	users, err := h.userService.GetListProfile()
	if err != nil {
		logger.Logger.Error("Cannot get list user: ", err)
		response.InternalError(c, "cannot get list user")
		return err
	}

	logger.Logger.Info("Get list user success, total length: ", len(users))
	response.Write(c, users)
	return nil
}

// RoleAddHandler handles adding role to user
func (h *UserHandler) RoleAddHandler(c *fiber.Ctx) error {
	var userRoleAPI UserRoleAPI
	err := c.BodyParser(&userRoleAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")
	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Add role for user",
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler add user role: ")
	err = h.userRoleService.RoleUserAdd(userRoleAPI.UserId, userRoleAPI.RoleId)
	if err != nil {
		logger.Logger.Error("Error add user role: ", err)
		response.InternalError(c, "Error add user role")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	return nil
}

// RoleDeleteHandler handles removing role from user
func (h *UserHandler) RoleDeleteHandler(c *fiber.Ctx) error {
	var userRoleAPI UserRoleAPI
	err := c.BodyParser(&userRoleAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")
	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Delete role for user",
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler add user role")
	h.userRoleService.RoleUserDelete(userRoleAPI.UserId, userRoleAPI.RoleId)
	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	return nil
}

// NetworkElementAddHandler handles adding network element to user
func (h *UserHandler) NetworkElementAddHandler(c *fiber.Ctx) error {
	var userNeAPI UserNeAPI
	err := c.BodyParser(&userNeAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")
	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Add permission for user",
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler add user ne: ")
	err = h.userNeService.NeUserAdd(userNeAPI.UserId, userNeAPI.NeId)
	if err != nil {
		logger.Logger.Error("Error add user ne: ", err)
		response.InternalError(c, "Error add user ne")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	return nil
}

// NetworkElementDeleteHandler handles removing network element from user
func (h *UserHandler) NetworkElementDeleteHandler(c *fiber.Ctx) error {
	var userNeAPI UserNeAPI
	err := c.BodyParser(&userNeAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")
	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Delete permission for user",
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler add user ne: ")
	err = h.userNeService.NeUserDelete(userNeAPI.UserId, userNeAPI.NeId)
	if err != nil {
		logger.Logger.Error("Error delete user ne: ", err)
		response.InternalError(c, "Error delete user ne")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	return nil
}

// NetworkElementsListHandler handles getting list of network elements for user
func (h *UserHandler) NetworkElementsListHandler(c *fiber.Ctx) error {
	var userAPI UserAPI
	err := c.BodyParser(&userAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	var username string
	if userAPI.Username == "" {
		username = c.Get("username")
	} else {
		username = userAPI.Username
	}

	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Add permission for user",
		NeName:   "",
		Mode:     "cli-config",
	}

	neList, err := h.userNeService.GetListNeByUser(username)
	if err != nil {
		logger.Logger.Error("Cannot get list ne form database, err: ", err)
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		response.BadRequest(c, "cannot get list ne of user")
		return err
	}

	logger.Logger.Info("Get list Ne of User success")
	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, neList)
	return nil
}

// PermissionGetHandler handles getting user permissions (placeholder)
func (h *UserHandler) PermissionGetHandler(c *fiber.Ctx) error {
	return nil
}

// ListUsersPermissionHandler handles getting all users permissions (placeholder)
func (h *UserHandler) ListUsersPermissionHandler(c *fiber.Ctx) error {
	return nil
}

// ListUsersNetworkElementHandler handles getting all users network elements (placeholder)
func (h *UserHandler) ListUsersNetworkElementHandler(c *fiber.Ctx) error {
	return nil
}

// NetworkElementsListDeleteHandler handles deleting list of network elements (placeholder)
func (h *UserHandler) NetworkElementsListDeleteHandler(c *fiber.Ctx) error {
	return nil
}
