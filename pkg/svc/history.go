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
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/env"
)

type History struct {
	Id           uint64
	Username     string
	UserIp       string
	Command      string
	NeName       string
	Result       bool
	ExecutedTime time.Time
	Mode         string
}

type HistoryService struct {
	repo repository.IDatabaseStore
}

func NewHistoryService() *HistoryService {
	return &HistoryService{
		repo: repository.GetSingleton(),
	}
}

func (s *HistoryService) toDb(historySVC *History) *models_db.OperationHistory {
	res := "failed"
	if historySVC.Result {
		res = "success"
	}
	return &models_db.OperationHistory{
		ID:           historySVC.Id,
		Username:     historySVC.Username,
		UserIP:       historySVC.UserIp,
		Command:      historySVC.Command,
		NeName:       historySVC.NeName,
		Result:       res,
		ExecutedTime: historySVC.ExecutedTime,
		Mode:         historySVC.Mode,
	}
}

func (s *HistoryService) fromDb(historyDB *models_db.OperationHistory) *History {
	res := false
	if historyDB.Result == "success" {
		res = true
	}
	return &History{
		Id:           historyDB.ID,
		Username:     historyDB.Username,
		UserIp:       historyDB.UserIP,
		Command:      historyDB.Command,
		NeName:       historyDB.NeName,
		Result:       res,
		ExecutedTime: historyDB.ExecutedTime,
		Mode:         historyDB.Mode,
	}
}

func (s *HistoryService) DeleteHistoryById(id uint64) error {
	return s.repo.DeleteHistoryById(id)
}

func (s *HistoryService) GetHistoryById(id uint64) (*History, error) {
	historyDb, err := s.repo.GetHistoryById(id)
	if err != nil {
		return nil, err
	}
	return s.fromDb(historyDb), nil
}

func (s *HistoryService) GetListHistoryByMode(mode string, limit int) ([]*History, error) {
	historiesDb, err := s.repo.GetHistoryCommandByModeLimit(mode, limit)
	if err != nil {
		logger.Logger.Error("Cannot get list history, err: ", err)
		return nil, err
	}

	var historiesSVC []*History
	for _, historyDb := range historiesDb {
		historiesSVC = append(historiesSVC, s.fromDb(&historyDb))
	}

	return historiesSVC, nil
}

func (s *HistoryService) GetHistoryByCommand(command string) (*History, error) {
	historyDb, err := s.repo.GetRecordHistoryByCommand(command)
	if err != nil {
		return nil, err
	}
	return s.fromDb(historyDb), nil
}

func (s *HistoryService) SaveHistoryCommand(history *History) error {
	historyDb := s.toDb(history)
	err := s.repo.SaveHistory(historyDb)
	if err != nil {
		logger.Logger.Error("Cannot save history command to database, err: ", err)
		return err
	}
	return nil
}

func (s *HistoryService) SaveHistoryCommandSuccess(history *History) error {
	history.ExecutedTime = time.Now()
	history.Result = true

	historyDb := s.toDb(history)
	err := s.repo.SaveHistory(historyDb)
	if err != nil {
		logger.Logger.Error("Cannot save history command to database, err: ", err)
		return err
	}
	return nil
}

func (s *HistoryService) SaveHistoryCommandFailure(history *History) error {
	history.ExecutedTime = time.Now()
	history.Result = false

	historyDb := s.toDb(history)
	err := s.repo.SaveHistory(historyDb)
	if err != nil {
		logger.Logger.Error("Cannot save history command to database, err: ", err)
		return err
	}
	return nil
}

func (s *HistoryService) SavingLogHistory() {
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

	neList, err := s.repo.GetListNetworkElement()
	if err != nil {
		logger.Logger.Error("Cannot get ne list, err: ", err)
		return
	}

	dateStr := time.Now().Format("02_01_2006")

	for _, ne := range neList {
		histories, err := s.repo.GetHistorySavingLog(ne.Namespace)
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
			buff.WriteString(strconv.FormatUint(historyCommand.ID, 10))
			buff.WriteString(",")
			buff.WriteString(historyCommand.Username)
			buff.WriteString(",")
			buff.WriteString(historyCommand.UserIP)
			buff.WriteString(",")
			buff.WriteString(historyCommand.Command)
			buff.WriteString(",")
			buff.WriteString(historyCommand.NeName)
			buff.WriteString(",")
			buff.WriteString(historyCommand.Result)
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
