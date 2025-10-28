package postgres

import (
	"errors"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"

	"gorm.io/gorm"
)

// CreateNetworkElement creates a new network element in database
func (c *PgClient) CreateNetworkElement(neDB *models_db.NetworkElement) error {
	// Create network element using GORM
	result := c.Db.Create(neDB)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to create network element: %v", result.Error)
		return result.Error
	}

	logger.Logger.Infof("Network element created successfully with ID: %d", neDB.ID)
	return nil
}

// DeleteNetworkElementByName deletes a network element by name and namespace
func (c *PgClient) DeleteNetworkElementByName(neName string, namespace string) error {
	result := c.Db.Where("name = ? AND namespace = ?", neName, namespace).Delete(&models_db.NetworkElement{})
	if result.Error != nil {
		logger.Logger.Errorf("Failed to delete network element: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		logger.Logger.Warnf("No network element found with name: %s, namespace: %s", neName, namespace)
		return errors.New("network element not found")
	}

	logger.Logger.Infof("Network element deleted successfully: %s (namespace: %s)", neName, namespace)
	return nil
}

// GetNetworkElementByName retrieves a network element by name and namespace
func (c *PgClient) GetNetworkElementByName(neName string, namespace string) (*models_db.NetworkElement, error) {
	var neDB models_db.NetworkElement

	result := c.Db.Where("name = ? AND namespace = ?", neName, namespace).First(&neDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.Logger.Warnf("Network element not found with name: %s, namespace: %s", neName, namespace)
			return nil, errors.New("network element not found")
		}
		logger.Logger.Errorf("Failed to get network element: %v", result.Error)
		return nil, result.Error
	}

	return &neDB, nil
}

// GetListNetworkElement retrieves all network elements
func (c *PgClient) GetListNetworkElement() ([]models_db.NetworkElement, error) {
	var nesDB []models_db.NetworkElement

	result := c.Db.Find(&nesDB)
	if result.Error != nil {
		logger.Logger.Errorf("Failed to list network elements: %v", result.Error)
		return nil, result.Error
	}

	logger.Logger.Infof("Retrieved %d network elements", len(nesDB))
	return nesDB, nil
}

// GetNetworkElementByUserName retrieves network elements accessible by a specific user
func (c *PgClient) GetNetworkElementByUserName(userName string) ([]models_db.NetworkElement, error) {
	var nesDB []models_db.NetworkElement

	// Use JOIN to get network elements associated with the user
	result := c.Db.
		Table("network_element ne").
		Select("ne.id, ne.name, ne.type, ne.namespace, ne.master_ip_config, ne.master_port_config, ne.slave_ip_config, ne.slave_port_config, ne.base_url, ne.ip_command, ne.port_command").
		Joins("JOIN user_ne un ON ne.id = un.ne_id").
		Joins("JOIN \"user\" u ON un.user_id = u.id").
		Where("u.username = ?", userName).
		Find(&nesDB)

	if result.Error != nil {
		logger.Logger.Errorf("Failed to get network elements for user %s: %v", userName, result.Error)
		return nil, result.Error
	}

	logger.Logger.Infof("Retrieved %d network elements for user %s", len(nesDB), userName)
	return nesDB, nil
}
