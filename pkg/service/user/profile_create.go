package user

import (
	"errors"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

func CreateProfile(user models_api.User) error {
	userFromDb, err := repository.GetSingleton().GetUserByUsername(user.Username)
	if err != nil {
		if !errors.Is(err, errors.New("user not found")) {
			logger.Logger.Error("Cannot get user by username from db, username: ", user.Username)
			return err
		}
	}

	if userFromDb.Username == user.Username {
		return errors.New("username already existed")
	}

	err = repository.GetSingleton().CreateUser(&user)
	if err != nil {
		logger.Logger.Error("Cannot create user with username: ", user.Username)
		return err
	}

	return nil
}
