package auth

import (
	"errors"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/service/utils/bcrypt"
	"go-cli-mgt/pkg/store/repository"
)

func Login(username, password string) (bool, error, uint) {
	user, err := repository.GetSingleton().GetUserByUsername(username)
	if err != nil {
		logger.Logger.Error("Cannot get user by username from db, username: ", username)
		return false, err, 0
	}

	if user.Active == false {
		logger.Logger.Info("username: ", username, " does not active")
		return false, errors.New("user not found"), 0
	}

	if !bcrypt.Matches(username+password, user.Password) {
		logger.Logger.Info("Login fail, wrong username or password")
		return false, nil, 0
	}

	return true, nil, user.Id
}

func GetRole(userId uint) (string, error) {
	roleList, err := repository.GetSingleton().GetRoleByUserId(userId)
	if err != nil {
		logger.Logger.Error("Cannot get role by user id: ", err)
		return "", err
	}
	roleStr := ""
	for _, role := range roleList {
		roleStr = roleStr + role.RoleName + " "
	}
	return roleStr, nil
}
