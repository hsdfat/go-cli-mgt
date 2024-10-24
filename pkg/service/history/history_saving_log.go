package history

import (
	"bufio"
	"bytes"
	"fmt"
	"go-cli-mgt/pkg/logger"
	neService "go-cli-mgt/pkg/service/network_elements"
	"go-cli-mgt/pkg/service/utils/env"
	"go-cli-mgt/pkg/store/repository"
	"os"
	"strconv"
	"strings"
	"time"
)

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

	neList, err := neService.GetListNetworkElement()
	if err != nil {
		logger.Logger.Error("Cannot get list ne, err: ", err)
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
