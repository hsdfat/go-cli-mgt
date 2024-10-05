package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

func GetProfile(username string) (*models_api.User, error) {
	user, err := repository.GetSingleton().GetUserByUsername(username)
	if err != nil {
		logger.Logger.Error("Cannot get user by username from database: ", err)
		return nil, err
	}
	return user, err
}
