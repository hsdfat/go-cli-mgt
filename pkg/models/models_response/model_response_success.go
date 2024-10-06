package models_response

type RespSuccess struct {
	Status   bool        `json:"status"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	DataResp interface{} `json:"dataResp"`
}
