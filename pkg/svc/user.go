package svc

import (
	"errors"
	"go-cli-mgt/pkg/logger"
	models_api "go-cli-mgt/pkg/models/api"
	models_error "go-cli-mgt/pkg/models/error"
	"go-cli-mgt/pkg/store/repository"
	"go-cli-mgt/pkg/utils/bcrypt"
	"time"
)

func CreateProfile(user models_api.User) error {
	userDb, err := repository.GetSingleton().GetUserByUsername(user.Username)
	if err != nil {
		if !errors.Is(err, models_error.ErrNotFoundUser) {
			logger.Logger.Error("Cannot get user by username from db, username: ", user.Username, " err: ", err)
			return err
		}
	}

	if userDb != nil && !userDb.Active {
		logger.Logger.Info("user disable")
		userDb.Active = true
		userDb.Password = bcrypt.Encode(user.Username + user.Password)
		err = repository.GetSingleton().UpdateUser(userDb)
		if err != nil {
			logger.Logger.Error("Cannot update user to database: ", err)
			return err
		}
		return nil
	}

	user.Active = true
	user.CreatedDate = uint64(time.Now().Unix())
	user.DisableDate = 1
	user.Password = bcrypt.Encode(user.Username + user.Password)

	err = repository.GetSingleton().CreateUser(&user)
	if err != nil {
		logger.Logger.Error("Cannot create user with username: ", user.Username)
		return err
	}

	return nil
}

func DeleteProfile(username string) error {
	_, err := repository.GetSingleton().GetUserByUsername(username)
	if err != nil {
		if !errors.Is(err, models_error.ErrNotFoundUser) {
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

func DisableProfile(username string, userDeactivate string) error {
	user, err := repository.GetSingleton().GetUserByUsername(username)
	if err != nil {
		logger.Logger.Error("Cannot get user by username from database: ", err)
		return err
	}

	if !user.Active {
		return models_error.ErrDisableUser
	}

	user.Active = false
	user.DisableDate = uint64(time.Now().Unix())
	user.DeActivateBy = userDeactivate
	err = repository.GetSingleton().UpdateUser(user)
	if err != nil {
		logger.Logger.Error("Cannot update user to database: ", err)
		return err
	}
	return nil
}

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
