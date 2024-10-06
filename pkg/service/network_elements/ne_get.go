package network_elements

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

func GetNetworkElement(neName, namespace string) (*models_api.NeData, error) {
	ne, err := repository.GetSingleton().GetNetworkElementByName(neName, namespace)
	if err != nil {
		logger.Logger.Error("Cannot get ne, err: ", err)
		return nil, err
	}
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
