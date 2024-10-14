package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

func UpdateProfile(user *models_api.User) error {
	err := repository.GetSingleton().UpdateUser(user)
	if err != nil {
		logger.Logger.Error("Cannot update user to database: ", err)
		return err
	}
	return nil
}

func UpdatePassword(user *models_api.User) {
	logger.Logger.Info("udpate password for user: ", user.Username)
	repository.GetSingleton().UpdatePasswordUser(user)
}
