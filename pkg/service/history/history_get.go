package history

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
)

func GetHistoryById(id uint64) (*models_api.History, error) {
	return repository.GetSingleton().GetHistoryById(id)
}

func GetListHistoryByMode(mode string, limit int) ([]models_api.History, error) {
	histories, err := repository.GetSingleton().GetHistoryCommandByModeLimit(mode, limit)
	if err != nil {
		logger.Logger.Error("Cannot get list history, err: ", err)
		return nil, err
	}
	return histories, nil
}

func GetHistoryByCommand(command string) (*models_api.History, error) {
	return repository.GetSingleton().GetRecordHistoryByCommand(command)
}
