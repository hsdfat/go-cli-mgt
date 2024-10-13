package user

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

func NeUserGet(userId, neId uint) (*models_api.UserNe, error) {
	logger.Logger.Infof("get user ne with userId %d and neId %d", userId, neId)
	userNe, err := repository.GetSingleton().UserNeGet(userId, neId)
	if err != nil {
		logger.Logger.Error("Cannot get user ne from database: ", err)
		return nil, err
	}
	logger.Logger.Infof("get user ne with userId %d and neId %d success", userId, neId)
	return userNe, nil
}
