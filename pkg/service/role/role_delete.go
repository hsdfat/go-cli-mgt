package role

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

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
