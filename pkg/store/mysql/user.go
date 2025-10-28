package mysql

import models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"

func (c *Client) CreateUser(user *models_api.User) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetUserByID(u uint) (*models_api.User, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) ListUsers() ([]models_api.User, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetUserByUsername(s string) (*models_api.User, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DeleteUser(s string) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UpdateUser(user *models_api.User) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UpdatePasswordUser(user *models_api.User) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UserNeAdd(ne *models_api.UserNe) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UserNeDelete(u uint, u2 uint) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UserNeGet(u uint, u2 uint) (*models_api.UserNe, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UserRoleAdd(role *models_api.UserRole) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UserRoleGet(userId, roleId uint) (*models_api.UserRole, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UserRoleDelete(userId, roleId uint) {
	//TODO implement me
	panic("implement me")
}
