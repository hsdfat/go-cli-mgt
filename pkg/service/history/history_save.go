package history

import (
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/store/repository"
	"time"
)

func SaveHistoryCommand(history *models_api.History) {
	err := repository.GetSingleton().SaveHistory(history)
	if err != nil {
		logger.Logger.Error("Cannot save history command to database, err: ", err)
	}
}

func SaveHistoryCommandSuccess(history *models_api.History) error {
	history.ExecutedTime = time.Now()
	history.Result = true

	err := repository.GetSingleton().SaveHistory(history)
	if err != nil {
		logger.Logger.Error("Cannot save history command to database, err: ", err)
		return err
	}
	return nil
}

func SaveHistoryCommandFailure(history *models_api.History) error {
	history.ExecutedTime = time.Now()
	history.Result = false

	err := repository.GetSingleton().SaveHistory(history)
	if err != nil {
		logger.Logger.Error("Cannot save history command to database, err: ", err)
		return err
	}
	return nil
}
