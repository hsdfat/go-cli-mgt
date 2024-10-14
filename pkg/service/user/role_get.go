package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

func RoleUserGet(userId, roleId uint) (*models_api.UserRole, error) {
	logger.Logger.Infof("get user ne with userId %d and roleId %d", userId, roleId)
	userRole, err := repository.GetSingleton().UserRoleGet(userId, roleId)
	if err != nil {
		logger.Logger.Error("Cannot get user role from database: ", err)
		return nil, err
	}
	logger.Logger.Infof("get user ne with userId %d and role %d success", userId, roleId)
	return userRole, nil
}
