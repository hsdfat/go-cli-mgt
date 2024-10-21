package models_api

type HistoryRequest struct {
	Mode  string `json:"mode"`
	Limit int    `json:"limit"`
}
