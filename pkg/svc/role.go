package svc

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
)

func CreateRole(role *models_api.Role) error {
	logger.Logger.Info("Create role: ", role.RoleName)
	err := repository.GetSingleton().CreateRole(role)
	if err != nil {
		logger.Logger.Error("Cannot create role, err: ", err)
		return err
	}
	logger.Logger.Infof("Create role %s success ", role.RoleName)
	return nil
}

func DeleteRole(role *models_api.Role) error {
	logger.Logger.Info("Delete role: ", role.RoleName)
	err := repository.GetSingleton().DeleteRole(role)
	if err != nil {
		logger.Logger.Error("Cannot delete role, err: ", err)
		return err
	}
	logger.Logger.Infof("Delete role %s success ", role.RoleName)
	return nil
}

func GetRoleByName(roleName string) (*models_api.Role, error) {
	logger.Logger.Info("Get role from database with role name ", roleName)
	role, err := repository.GetSingleton().GetRoleByName(roleName)
	if err != nil {
		logger.Logger.Error("Cannot get role by role name from database, err: ", err)
		return nil, err
	}
	logger.Logger.Info("Get role success from database with role name ", roleName)
	return role, nil
}

func GetListRole() ([]models_api.Role, error) {
	roleList, err := repository.GetSingleton().GetListRole()
	if err != nil {
		logger.Logger.Error("Cannot get role list, err: ", err)
		return nil, err
	}
	return roleList, nil
}

func UpdateRole(role *models_api.Role) {
	logger.Logger.Info("Update role from database with role name ", role.RoleName)
	logger.Logger.Debug("Description: ", role.Description)
	err := repository.GetSingleton().UpdateRole(role)
	if err != nil {
		logger.Logger.Error("Cannot update role, err: ", err)
	}
}
