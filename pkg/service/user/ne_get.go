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

func GetListNeByUser(username string) ([]models_api.NeData, error) {
	logger.Logger.Info("get list ne of user: ", username)
	neList, err := repository.GetSingleton().GetNetworkElementByUserName(username)
	if err != nil {
		logger.Logger.Error("Cannot get ne from database: ", err)
		return nil, err
	}
	logger.Logger.Infof("get list ne of user %s success", username)
	return neList, nil
}
