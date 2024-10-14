package role

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
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
