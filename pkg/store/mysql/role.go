package mysql

import (
	models_api "go-cli-mgt/pkg/models/api"
	models_db "go-cli-mgt/pkg/models/db"
)

func (c *Client) GetRoleByUserId(u uint) ([]models_db.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetRoleByName(s string) (*models_api.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetListRole() ([]models_api.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) CreateRole(role *models_api.Role) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DeleteRole(role *models_api.Role) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UpdateRole(role *models_api.Role) error {
	//TODO implement me
	panic("implement me")
}
