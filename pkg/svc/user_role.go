package svc

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
)

func RoleUserAdd(userId, roleId uint) error {
	err := repository.GetSingleton().UserRoleAdd(&models_api.UserRole{
		UserId: userId,
		RoleId: roleId,
	})
	if err != nil {
		logger.Logger.Error("Cannot add user role into database: ", err)
		return err
	}
	return nil
}

func RoleUserDelete(userId, roleId uint) {
	logger.Logger.Infof("Delete user id %d with role id %d", userId, roleId)
	repository.GetSingleton().UserRoleDelete(userId, roleId)
}

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
