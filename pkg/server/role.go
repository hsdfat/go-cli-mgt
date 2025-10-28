package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"
)

// RoleAPI represents the API model for role
type RoleAPI struct {
	RoleId      uint   `json:"role-id,omitempty"`
	RoleName    string `json:"role-name,omitempty"`
	Priority    string `json:"priority,omitempty"`
	Description string `json:"description,omitempty"`
}

// RoleHandler struct contains role and history services
type RoleHandler struct {
	roleService    *svc.RoleService
	historyService *svc.HistoryService
}

// NewHandler initializes handler with services
func NewRoleHandler() *RoleHandler {
	return &RoleHandler{
		roleService:    svc.NewRoleService(),
		historyService: svc.NewHistoryService(),
	}
}

// apiToSvcRole converts API model to service model
func (h *RoleHandler) apiToSvcRole(roleAPI *RoleAPI) *svc.Role {
	return &svc.Role{
		RoleId:      roleAPI.RoleId,
		RoleName:    roleAPI.RoleName,
		Description: roleAPI.Description,
	}
}

// ListRoleHandler handles getting role list
func (h *RoleHandler) ListRoleHandler(c *fiber.Ctx) error {
	username := c.Get("username")

	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Get list role",
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler get role list")
	roleList, err := h.roleService.GetListRole()
	if err != nil {
		logger.Logger.Error("Error get list role: ", err)
		response.InternalError(c, "Error get list role")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, roleList)
	return nil
}

// CreateOrUpdateHandler handles creating or updating role
func (h *RoleHandler) CreateOrUpdateHandler(c *fiber.Ctx) error {
	var roleAPI RoleAPI
	err := c.BodyParser(&roleAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}
	logger.Logger.Debug("Description: ", roleAPI.Description)

	username := c.Get("username")

	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Create role " + roleAPI.RoleName,
		NeName:   "",
		Mode:     "cli-config",
	}

	// Check if role already exists
	roleGet, err := h.roleService.GetRoleByName(roleAPI.RoleName)
	if roleGet == nil {
		// Create new role
		logger.Logger.Info("Handler create role ", roleAPI.RoleName)
		roleSVC := h.apiToSvcRole(&roleAPI)
		err = h.roleService.CreateRole(roleSVC)
		if err != nil {
			logger.Logger.Error("Error create role: ", err)
			response.InternalError(c, "Error create role")
			h.historyService.SaveHistoryCommandFailure(historyCommand)
			return err
		}
		h.historyService.SaveHistoryCommandSuccess(historyCommand)
		logger.Logger.Infof("Create role %s success ", roleAPI.RoleName)
	} else {
		// Update existing role
		logger.Logger.Info("Handler update role ", roleAPI.RoleName)
		roleGet.Description = roleAPI.Description
		h.roleService.UpdateRole(roleGet)
		h.historyService.SaveHistoryCommandSuccess(historyCommand)
		logger.Logger.Infof("Update role %s success ", roleAPI.RoleName)
	}

	return nil
}

// DeleteHandler handles deleting role
func (h *RoleHandler) DeleteHandler(c *fiber.Ctx) error {
	var roleAPI RoleAPI
	err := c.BodyParser(&roleAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")

	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Delete role " + roleAPI.RoleName,
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler delete role ", roleAPI.RoleName)

	roleSVC := h.apiToSvcRole(&roleAPI)
	err = h.roleService.DeleteRole(roleSVC)
	if err != nil {
		logger.Logger.Error("Error delete role: ", err)
		response.InternalError(c, "Error delete role")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	logger.Logger.Infof("Delete role %s success ", roleAPI.RoleName)
	return nil
}
