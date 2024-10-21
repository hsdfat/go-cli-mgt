package unit

import (
	"github.com/stretchr/testify/require"
	"go-cli-mgt/pkg/models/models_api"
	historyService "go-cli-mgt/pkg/service/history"
	"go-cli-mgt/pkg/service/utils/random"
	"testing"
	"time"
)

func TestSaveHistorySuccess(t *testing.T) {
	// history data
	historyCommand := &models_api.History{
		Username: random.StringRandom(10),
		UserIp:   random.StringRandom(10),
		Command:  random.StringRandom(10),
		NeName:   random.StringRandom(10),
		Mode:     random.StringRandom(10),
	}

	// Save history
	err := historyService.SaveHistoryCommandSuccess(historyCommand)
	require.NoError(t, err)

	// Get history
	historyCommandGet, err := historyService.GetHistoryById(historyCommand.Id)
	require.NoError(t, err)
	require.NotEmpty(t, historyCommandGet)
	require.NotEmpty(t, historyCommandGet.ExecutedTime)
	require.Equal(t, historyCommand.Id, historyCommandGet.Id)
	require.Equal(t, historyCommand.Username, historyCommandGet.Username)
	require.Equal(t, historyCommand.UserIp, historyCommandGet.UserIp)
	require.Equal(t, historyCommand.Command, historyCommandGet.Command)
	require.Equal(t, historyCommand.NeName, historyCommandGet.NeName)
	require.Equal(t, historyCommand.Mode, historyCommandGet.Mode)
	require.Equal(t, true, historyCommandGet.Result)

	// Delete History
	err = historyService.DeleteHistoryById(historyCommand.Id)
}

func TestSaveHistoryFailure(t *testing.T) {
	// history data
	historyCommand := &models_api.History{
		Username: random.StringRandom(10),
		UserIp:   random.StringRandom(10),
		Command:  random.StringRandom(10),
		NeName:   random.StringRandom(10),
		Mode:     random.StringRandom(10),
	}

	// Save history
	err := historyService.SaveHistoryCommandFailure(historyCommand)
	require.NoError(t, err)

	// Get history
	historyCommandGet, err := historyService.GetHistoryById(historyCommand.Id)
	require.NoError(t, err)
	require.NotEmpty(t, historyCommandGet)
	require.NotEmpty(t, historyCommandGet.ExecutedTime)
	require.Equal(t, historyCommand.Id, historyCommandGet.Id)
	require.Equal(t, historyCommand.Username, historyCommandGet.Username)
	require.Equal(t, historyCommand.UserIp, historyCommandGet.UserIp)
	require.Equal(t, historyCommand.Command, historyCommandGet.Command)
	require.Equal(t, historyCommand.NeName, historyCommandGet.NeName)
	require.Equal(t, historyCommand.Mode, historyCommandGet.Mode)
	require.Equal(t, false, historyCommandGet.Result)

	// Delete History
	err = historyService.DeleteHistoryById(historyCommand.Id)
}

func TestSaveHistory(t *testing.T) {
	// history data
	historyCommand := models_api.History{
		Username:     "userTest1",
		UserIp:       random.Ipv4Random(),
		Command:      random.StringRandom(10),
		NeName:       random.StringRandom(10),
		Result:       random.BooleanRandom(),
		ExecutedTime: time.Now(),
		Mode:         random.StringRandom(10),
	}

	// Save history
	err := historyService.SaveHistoryCommand(&historyCommand)
	require.NoError(t, err)

	// Get History
	historyCommandGet, err := historyService.GetHistoryByCommand(historyCommand.Command)
	require.NoError(t, err)
	require.NotEmpty(t, historyCommandGet)
	require.Equal(t, historyCommand.Id, historyCommandGet.Id)
	require.Equal(t, historyCommand.Username, historyCommandGet.Username)
	require.Equal(t, historyCommand.UserIp, historyCommandGet.UserIp)
	require.Equal(t, historyCommand.Command, historyCommandGet.Command)
	require.Equal(t, historyCommand.NeName, historyCommandGet.NeName)
	require.Equal(t, historyCommand.Result, historyCommandGet.Result)
	//require.Equal(t, historyCommand.ExecutedTime, historyCommandGet.ExecutedTime)
	require.Equal(t, historyCommand.Mode, historyCommandGet.Mode)

	// Delete History
	err = historyService.DeleteHistoryById(historyCommand.Id)
}
