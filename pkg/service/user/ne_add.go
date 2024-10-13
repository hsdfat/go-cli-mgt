package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

func NeUserAdd(userId, neId uint) error {
	err := repository.GetSingleton().UserNeAdd(&models_api.UserNe{
		UserId: userId,
		NeId:   neId,
	})
	if err != nil {
		logger.Logger.Error("Cannot add user ne from database: ", err)
		return err
	}
	return nil
}
