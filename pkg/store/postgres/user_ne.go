package postgres

import (
	"errors"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"

	"gorm.io/gorm"
)

// UserNeAdd adds a user-network element relationship
func (c *PgClient) UserNeAdd(userNeDB *models_db.UserNe) error {
	// Create user_ne record using GORM
	result := c.Db.Create(userNeDB)
	if result.Error != nil {
		return result.Error
	}

	logger.Logger.Infof("User-NE relationship added successfully: UserID %d, NeID %d", userNeDB.UserID, userNeDB.NeID)
	return nil
}

// UserNeDelete deletes a user-network element relationship
func (c *PgClient) UserNeDelete(userId, neId uint) error {
	result := c.Db.Where("user_id = ? AND ne_id = ?", userId, neId).Delete(&models_db.UserNe{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user-ne relationship not found")
	}

	logger.Logger.Infof("User-NE relationship deleted successfully: UserID %d, NeID %d", userId, neId)
	return nil
}

// UserNeGet retrieves a user-network element relationship
func (c *PgClient) UserNeGet(userId, neId uint) (*models_db.UserNe, error) {
	var userNeDB models_db.UserNe

	result := c.Db.Where("user_id = ? AND ne_id = ?", userId, neId).First(&userNeDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.Logger.Warnf("User-NE relationship not found: UserID %d, NeID %d", userId, neId)
			return nil, errors.New("user do not have permission with ne")
		}
		logger.Logger.Errorf("Failed to get user-ne relationship: %v", result.Error)
		return nil, result.Error
	}

	return &userNeDB, nil
}
