package postgres

import (
	"errors"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"

	"gorm.io/gorm"
)

// CreateUser creates a new user in database
func (c *PgClient) CreateUser(userDB *models_db.User) error {
	// Create user using GORM
	result := c.Db.Create(userDB)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to create user: %v", result.Error)
		return result.Error
	}

	logger.Logger.Infof("User created successfully with ID: %d", userDB.ID)
	return nil
}

// DeleteUser deletes a user by username
func (c *PgClient) DeleteUser(username string) error {
	result := c.Db.Where("username = ?", username).Delete(&models_db.User{})
	if result.Error != nil {
		logger.Logger.Errorf("Failed to delete user: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.Logger.Warnf("No user found with username: %s", username)
		return errors.New("user not found")
	}

	logger.Logger.Infof("User deleted successfully: %s", username)
	return nil
}

// UpdateUser updates user information
func (c *PgClient) UpdateUser(userUpdate *models_db.User) error {
	result := c.Db.Model(&models_db.User{}).Where("id = ?", userUpdate.ID).Updates(userUpdate)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to update user: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.Logger.Warnf("No user found with ID: %d", userUpdate.ID)
		return errors.New("user not found")
	}

	logger.Logger.Infof("User updated successfully: ID %d", userUpdate.ID)
	return nil
}

// UpdatePasswordUser updates only the password of a user
func (c *PgClient) UpdatePasswordUser(userDB *models_db.User) {
	result := c.Db.Model(&models_db.User{}).Where("id = ?", userDB.ID).Update("password", userDB.Password)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to update password for user ID %d: %v", userDB.ID, result.Error)
		return
	}

	logger.Logger.Infof("Password updated successfully for user ID: %d", userDB.ID)
}

// GetUserByID retrieves a user by ID
func (c *PgClient) GetUserByID(id uint) (*models_db.User, error) {
	var userDB models_db.User

	result := c.Db.Where("id = ?", id).First(&userDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.Logger.Warnf("User not found with ID: %d", id)
			return nil, errors.New("user not found")
		}
		logger.Logger.Errorf("Failed to get user by ID: %v", result.Error)
		return nil, result.Error
	}

	return &userDB, nil
}

// GetUserByUsername retrieves a user by username
func (c *PgClient) GetUserByUsername(username string) (*models_db.User, error) {
	var userDB models_db.User

	result := c.Db.Where("username = ?", username).First(&userDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.Logger.Warnf("User not found with username: %s", username)
			return nil, errors.New("user not found")
		}
		logger.Logger.Errorf("Failed to get user by username: %v", result.Error)
		return nil, result.Error
	}

	return &userDB, nil
}

// ListUsers retrieves all users
func (c *PgClient) ListUsers() ([]models_db.User, error) {
	var usersDB []models_db.User

	result := c.Db.Find(&usersDB)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to list users: %v", result.Error)
		return nil, result.Error
	}

	logger.Logger.Infof("Retrieved %d users", len(usersDB))
	return usersDB, nil
}
