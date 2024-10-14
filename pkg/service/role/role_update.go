package role

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

func UpdateRole(role *models_api.Role) {
	logger.Logger.Info("Update role from database with role name ", role.RoleName)
	logger.Logger.Debug("Description: ", role.Description)
	err := repository.GetSingleton().UpdateRole(role)
	if err != nil {
		logger.Logger.Error("Cannot update role, err: ", err)
	}
}
