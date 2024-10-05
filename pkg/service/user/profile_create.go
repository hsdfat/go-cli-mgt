package user

import (
	"errors"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/service/utils/bcrypt"
	"go-cli-mgt/pkg/store/postgres"
	"go-cli-mgt/pkg/store/repository"
	"time"
)

func CreateProfile(user models_api.User) error {
	userDb, err := repository.GetSingleton().GetUserByUsername(user.Username)
	if err != nil {
		if !errors.Is(err, postgres.ErrNotFoundUser) {
			logger.Logger.Error("Cannot get user by username from db, username: ", user.Username, " err: ", err)
			return err
		}
	}

	if userDb != nil && userDb.Active == false {
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
	user.Password = bcrypt.Encode(user.Username + user.Password)

	err = repository.GetSingleton().CreateUser(&user)
	if err != nil {
		logger.Logger.Error("Cannot create user with username: ", user.Username)
		return err
	}

	return nil
}
