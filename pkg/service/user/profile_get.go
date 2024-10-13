package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

func GetProfileByUsername(username string) (*models_api.User, error) {
	logger.Logger.Info("Get profile username: ", username)
	user, err := repository.GetSingleton().GetUserByUsername(username)
	if err != nil {
		logger.Logger.Error("Cannot get user by username from database: ", err)
		return nil, err
	}
	logger.Logger.Infof("Get profile success with username %s and id %d ", username, user.Id)
	return user, err
}

func GetListProfile() ([]models_api.User, error) {
	users, err := repository.GetSingleton().ListUsers()
	if err != nil {
		logger.Logger.Error("Cannot get list user from database: ", err)
		return nil, err
	}
	return users, nil
}
