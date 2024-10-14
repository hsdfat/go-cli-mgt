package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
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
