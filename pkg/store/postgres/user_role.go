package postgres

import (
	"errors"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"

	"gorm.io/gorm"
)

// UserRoleAdd adds a user-role relationship
func (c *PgClient) UserRoleAdd(userRoleDB *models_db.UserRole) error {
	// Create user_role record using GORM
	result := c.Db.Create(userRoleDB)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to add user-role relationship: %v", result.Error)
		return result.Error
	}

	logger.Logger.Infof("User-Role relationship added successfully: UserID %d, RoleID %d", userRoleDB.UserID, userRoleDB.RoleID)
	return nil
}

// UserRoleGet retrieves a user-role relationship
func (c *PgClient) UserRoleGet(userId, roleId uint) (*models_db.UserRole, error) {
	var userRoleDB models_db.UserRole

	result := c.Db.Where("user_id = ? AND role_id = ?", userId, roleId).First(&userRoleDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.Logger.Warnf("User-Role relationship not found: UserID %d, RoleID %d", userId, roleId)
			return nil, errors.New("user do not have this role")
		}
		logger.Logger.Errorf("Failed to get user-role relationship: %v", result.Error)
		return nil, result.Error
	}

	return &userRoleDB, nil
}

// UserRoleDelete deletes a user-role relationship
func (c *PgClient) UserRoleDelete(userId, roleId uint) {
	result := c.Db.Where("user_id = ? AND role_id = ?", userId, roleId).Delete(&models_db.UserRole{})
	if result.Error != nil {
		logger.Logger.Errorf("Failed to delete user-role relationship: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		logger.Logger.Warnf("No user-role relationship found: UserID %d, RoleID %d", userId, roleId)
		return
	}

	logger.Logger.Infof("User-Role relationship deleted successfully: UserID %d, RoleID %d", userId, roleId)
}
