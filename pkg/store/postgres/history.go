package postgres

import (
	"errors"
	"time"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"

	"gorm.io/gorm"
)

// SaveHistory saves operation history to database
func (c *PgClient) SaveHistory(historyDB *models_db.OperationHistory) error {
	// Create history record using GORM
	result := c.Db.Create(historyDB)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to save history: %v", result.Error)
		return result.Error
	}

	logger.Logger.Infof("History saved successfully with ID: %d", historyDB.ID)
	return nil
}

// GetHistoryById retrieves a history record by ID
func (c *PgClient) GetHistoryById(id uint64) (*models_db.OperationHistory, error) {
	var historyDB models_db.OperationHistory

	result := c.Db.Where("id = ?", id).First(&historyDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.Logger.Warnf("History not found with ID: %d", id)
			return nil, errors.New("history command not found")
		}
		logger.Logger.Errorf("Failed to get history by ID: %v", result.Error)
		return nil, result.Error
	}

	return &historyDB, nil
}

// DeleteHistoryById deletes a history record by ID
func (c *PgClient) DeleteHistoryById(id uint64) error {
	result := c.Db.Where("id = ?", id).Delete(&models_db.OperationHistory{})
	if result.Error != nil {
		logger.Logger.Errorf("Failed to delete history: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.Logger.Warnf("No history found with ID: %d", id)
		return errors.New("history not found")
	}

	logger.Logger.Infof("History deleted successfully: ID %d", id)
	return nil
}

// GetHistoryListByMode retrieves all history records filtered by mode
func (c *PgClient) GetHistoryListByMode(mode string) ([]models_db.OperationHistory, error) {
	var historiesDB []models_db.OperationHistory

	result := c.Db.Where("mode = ?", mode).Find(&historiesDB)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to get history list by mode: %v", result.Error)
		return nil, result.Error
	}

	logger.Logger.Infof("Retrieved %d history records for mode: %s", len(historiesDB), mode)
	return historiesDB, nil
}

// GetRecordHistoryByCommand retrieves a history record by command (for testing)
// This function only use for test check if existed record in database
func (c *PgClient) GetRecordHistoryByCommand(command string) (*models_db.OperationHistory, error) {
	var historyDB models_db.OperationHistory

	result := c.Db.Where("command = ?", command).First(&historyDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.Logger.Warnf("History not found with command: %s", command)
			return nil, errors.New("history command not found")
		}
		logger.Logger.Errorf("Failed to get history by command: %v", result.Error)
		return nil, result.Error
	}

	return &historyDB, nil
}

// GetHistoryCommandByModeLimit retrieves history records filtered by mode with limit
func (c *PgClient) GetHistoryCommandByModeLimit(mode string, limit int) ([]models_db.OperationHistory, error) {
	var historiesDB []models_db.OperationHistory

	result := c.Db.
		Where("mode = ?", mode).
		Order("id DESC").
		Limit(limit).
		Find(&historiesDB)

	if result.Error != nil {
		logger.Logger.Errorf("Failed to get history by mode with limit: %v", result.Error)
		return nil, result.Error
	}

	logger.Logger.Infof("Retrieved %d history records for mode: %s (limit: %d)", len(historiesDB), mode, limit)
	return historiesDB, nil
}

// GetHistorySavingLog retrieves history records for a specific network element from last 24 hours
func (c *PgClient) GetHistorySavingLog(neSiteName string) ([]models_db.OperationHistory, error) {
	var historiesDB []models_db.OperationHistory

	// Calculate time 24 hours ago
	oneDayAgo := time.Now().Add(-24 * time.Hour)

	result := c.Db.
		Where("ne_name = ? AND executed_time >= ?", neSiteName, oneDayAgo).
		Find(&historiesDB)

	if result.Error != nil {
		logger.Logger.Errorf("Failed to get saving log history: %v", result.Error)
		return nil, result.Error
	}

	logger.Logger.Infof("Retrieved %d saving log history records for NE: %s", len(historiesDB), neSiteName)
	return historiesDB, nil
}
