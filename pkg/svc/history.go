package svc

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/env"
)

func DeleteHistoryById(id uint64) error {
	return repository.GetSingleton().DeleteHistoryById(id)
}

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

func SaveHistoryCommand(history *models_api.History) error {
	err := repository.GetSingleton().SaveHistory(history)
	if err != nil {
		logger.Logger.Error("Cannot save history command to database, err: ", err)
		return err
	}
	return nil
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

func SavingLogHistory() {
	dir := env.GetEnv("SAVING_LOG_DIR", "mme/history")
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		logger.Logger.Info("Directory created successfully:", dir)
	}
	logger.Logger.Info("Directory already exists:", dir)

	neList, err := repository.GetSingleton().GetListNetworkElement()
	if err != nil {
		logger.Logger.Error("Cannot get ne list, err: ", err)
		return
	}

	dateStr := time.Now().Format("02_01_2006")

	for _, ne := range neList {
		histories, err := repository.GetSingleton().GetHistorySavingLog(ne.Namespace)
		if err != nil {
			logger.Logger.Error("Cannot get histories list, err: ", err)
			continue
		}
		filePath := dir + "/" + ne.Namespace + "_" + dateStr + ".txt"
		isTestFile := env.GetEnv("SAVING_LOG_TEST", "false")
		if strings.EqualFold(isTestFile, "true") {
			filePath = filePath + ".test"
		}
		file, err := os.Create(filePath)
		if err != nil {
			logger.Logger.Errorf("Cannot create file %s err %s", filePath, err.Error())
			continue
		}
		writer := bufio.NewWriter(file)
		buff := bytes.NewBufferString("")
		for _, historyCommand := range histories {
			buff.WriteString(strconv.FormatUint(historyCommand.Id, 10))
			buff.WriteString(",")
			buff.WriteString(historyCommand.Username)
			buff.WriteString(",")
			buff.WriteString(historyCommand.UserIp)
			buff.WriteString(",")
			buff.WriteString(historyCommand.Command)
			buff.WriteString(",")
			buff.WriteString(historyCommand.NeName)
			buff.WriteString(",")
			buff.WriteString(fromBoolToString(historyCommand.Result))
			buff.WriteString(",")
			buff.WriteString(historyCommand.ExecutedTime.Format("05:04:15 02_01_2006"))
			buff.WriteString("\n")
		}
		writer.Write(buff.Bytes())
		writer.Flush()
		file.Close()
		logger.Logger.Info("Saving history log, file path: ", filePath)
	}
}

func fromBoolToString(flag bool) string {
	if flag {
		return "success"
	}
	return "failure"
}
