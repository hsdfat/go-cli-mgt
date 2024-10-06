package network_elements

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
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
