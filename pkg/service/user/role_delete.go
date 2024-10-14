package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/store/repository"
)

func RoleUserDelete(userId, roleId uint) {
	logger.Logger.Infof("Delete user id %d with role id %d", userId, roleId)
	repository.GetSingleton().UserRoleDelete(userId, roleId)
}
