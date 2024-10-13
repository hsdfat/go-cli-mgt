package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/store/repository"
)

func NeUserDelete(userId, neId uint) error {
	err := repository.GetSingleton().UserNeDelete(userId, neId)
	if err != nil {
		logger.Logger.Error("Cannot delete user ne from database: ", err)
		return err
	}
	return nil
}
