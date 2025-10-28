package svc

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
)

func NeUserAdd(userId, neId uint) error {
	err := repository.GetSingleton().UserNeAdd(&models_api.UserNe{
		UserId: userId,
		NeId:   neId,
	})
	if err != nil {
		logger.Logger.Error("Cannot add user ne into database: ", err)
		return err
	}
	return nil
}

func NeUserDelete(userId, neId uint) error {
	err := repository.GetSingleton().UserNeDelete(userId, neId)
	if err != nil {
		logger.Logger.Error("Cannot delete user ne from database: ", err)
		return err
	}
	return nil
}

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
