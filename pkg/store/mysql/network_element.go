package mysql

import models_api "github.com/hsdfat/go-cli-mgt/pkg/models/api"

func (c *Client) CreateNetworkElement(data *models_api.NeData) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DeleteNetworkElementByName(s string, s2 string) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetNetworkElementByName(s string, s2 string) (*models_api.NeData, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetListNetworkElement() ([]models_api.NeData, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetNetworkElementByUserName(s string) ([]models_api.NeData, error) {
	//TODO implement me
	panic("implement me")
}
