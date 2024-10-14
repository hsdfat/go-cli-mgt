package role

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

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
