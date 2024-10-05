package user

import (
	"errors"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/store/postgres"
	"go-cli-mgt/pkg/store/repository"
)

func DeleteProfile(username string) error {
	_, err := repository.GetSingleton().GetUserByUsername(username)
	if err != nil {
		if !errors.Is(err, postgres.ErrNotFoundUser) {
			logger.Logger.Error("Cannot get user by username from database: ", err)
			return err
		}
	}
	err = repository.GetSingleton().DeleteUser(username)
	if err != nil {
		logger.Logger.Error("Cannot delete user by username from database: ", err)
		return err
	}
	logger.Logger.Info("Delete user complete, username: ", username)
	return nil
}

func DisableProfile(username string) error {
	user, err := repository.GetSingleton().GetUserByUsername(username)
	if err != nil {
		logger.Logger.Error("Cannot get user by username from database: ", err)
		return err
	}

	if user.Active == false {
		logger.Logger.Error(postgres.ErrDisableUser)
		return postgres.ErrDisableUser
	}

	user.Active = false
	err = repository.GetSingleton().UpdateUser(user)
	if err != nil {
		logger.Logger.Error("Cannot update user to database: ", err)
		return err
	}
	return nil
}
