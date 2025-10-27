package svc

import (
	"go-cli-mgt/pkg/logger"
	models_api "go-cli-mgt/pkg/models/api"
	"go-cli-mgt/pkg/store/repository"
)

func CreateNetworkElement(ne *models_api.NeData) error {
	err := repository.GetSingleton().CreateNetworkElement(ne)
	if err != nil {
		logger.Logger.Error("Cannot create network element, err: ", err)
		return err
	}
	return nil
}

func DeleteNetworkElement(neName string, namespace string) error {
	err := repository.GetSingleton().DeleteNetworkElementByName(neName, namespace)
	if err != nil {
		logger.Logger.Error("Cannot delete network element, err: ", err)
		return err
	}
	return nil
}

func GetNetworkElement(neName, namespace string) (*models_api.NeData, error) {
	logger.Logger.Infof("Get network element info with ne %s namespace %s", neName, namespace)
	ne, err := repository.GetSingleton().GetNetworkElementByName(neName, namespace)
	if err != nil {
		logger.Logger.Error("Cannot get ne, err: ", err)
		return nil, err
	}
	logger.Logger.Infof("Get network element info with ne %s namespace %s success with id %d", neName, namespace, ne.NeId)
	return ne, nil
}

func GetListNetworkElement() ([]models_api.NeData, error) {
	neList, err := repository.GetSingleton().GetListNetworkElement()
	if err != nil {
		logger.Logger.Error("Cannot get ne list, err: ", err)
		return nil, err
	}
	return neList, nil
}
