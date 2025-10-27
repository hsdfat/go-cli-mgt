package mysql

import models_api "go-cli-mgt/pkg/models/api"

func (c *Client) SaveHistory(history *models_api.History) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetHistoryById(u uint64) (*models_api.History, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DeleteHistoryById(u uint64) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetHistoryListByMode(s string) ([]models_api.History, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetRecordHistoryByCommand(s string) (*models_api.History, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetHistoryCommandByModeLimit(s string, i int) ([]models_api.History, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetHistorySavingLog(neSiteName string) ([]models_api.History, error) {
	//TODO implement me
	panic("implement me")
}
