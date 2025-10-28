package postgres

import (
	"errors"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"

	"gorm.io/gorm"
)

// GetRoleByUserId retrieves all roles assigned to a user
func (c *PgClient) GetRoleByUserId(userId uint) ([]models_db.Role, error) {
	var rolesDB []models_db.Role

	// Use JOIN to get roles associated with the user
	result := c.Db.
		Table("role r").
		Select("r.id, r.role_name, r.description").
		Joins("JOIN user_role ur ON r.id = ur.role_id").
		Where("ur.user_id = ?", userId).
		Find(&rolesDB)

	if result.Error != nil {
		logger.Logger.Errorf("Failed to get roles for user ID %d: %v", userId, result.Error)
		return nil, result.Error
	}

	logger.Logger.Infof("Retrieved %d roles for user ID %d", len(rolesDB), userId)
	return rolesDB, nil
}

// GetRoleByName retrieves a role by its name
func (c *PgClient) GetRoleByName(roleName string) (*models_db.Role, error) {
	var roleDB models_db.Role

	result := c.Db.Where("role_name = ?", roleName).First(&roleDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.Logger.Warnf("Role not found with name: %s", roleName)
			return nil, errors.New("role not found")
		}
		logger.Logger.Errorf("Failed to get role by name: %v", result.Error)
		return nil, result.Error
	}

	return &roleDB, nil
}

// GetListRole retrieves all roles
func (c *PgClient) GetListRole() ([]models_db.Role, error) {
	var rolesDB []models_db.Role

	result := c.Db.Find(&rolesDB)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to list roles: %v", result.Error)
		return nil, result.Error
	}

	logger.Logger.Infof("Retrieved %d roles", len(rolesDB))
	return rolesDB, nil
}

// CreateRole creates a new role in database
func (c *PgClient) CreateRole(roleDB *models_db.Role) error {
	// Create role using GORM
	result := c.Db.Create(roleDB)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to create role: %v", result.Error)
		return result.Error
	}

	logger.Logger.Infof("Role created successfully with ID: %d", roleDB.ID)
	return nil
}

// DeleteRole deletes a role from database
func (c *PgClient) DeleteRole(roleDB *models_db.Role) error {
	result := c.Db.Where("role_name = ?", roleDB.RoleName).Delete(&models_db.Role{})
	if result.Error != nil {
		logger.Logger.Errorf("Failed to delete role: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.Logger.Warnf("No role found with name: %s", roleDB.RoleName)
		return errors.New("role not found")
	}

	logger.Logger.Infof("Role deleted successfully: %s", roleDB.RoleName)
	return nil
}

// UpdateRole updates role information
func (c *PgClient) UpdateRole(roleDB *models_db.Role) error {
	updates := map[string]interface{}{
		"description": roleDB.Description,
	}

	result := c.Db.Model(&models_db.Role{}).Where("role_name = ?", roleDB.RoleName).Updates(updates)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to update role: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.Logger.Warnf("No role found with name: %s", roleDB.RoleName)
		return errors.New("role not found")
	}

	logger.Logger.Infof("Role updated successfully: %s", roleDB.RoleName)
	return nil
}
