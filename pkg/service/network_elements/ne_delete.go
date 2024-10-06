package network_elements

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/store/repository"
)

func DeleteNetworkElement(neName string, namespace string) error {
	err := repository.GetSingleton().DeleteNetworkElementByName(neName, namespace)
	if err != nil {
		logger.Logger.Error("Cannot delete network element, err: ", err)
		return err
	}
	return nil
}
