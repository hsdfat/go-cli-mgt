package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"
)

// HistoryRequest represents the request for getting history list
type HistoryRequest struct {
	Mode  string `json:"mode"`
	Limit int    `json:"limit"`
}

// HistoryAPI represents the API model for history
type HistoryAPI struct {
	Id           uint64    `json:"id"`
	Username     string    `json:"username,omitempty"`
	UserIp       string    `json:"user-ip,omitempty"`
	Command      string    `json:"command,omitempty"`
	NeName       string    `json:"ne-name,omitempty"`
	Result       bool      `json:"result,omitempty"`
	ExecutedTime time.Time `json:"executed-time,omitempty"`
	Mode         string    `json:"mode"`
}

// HistoryHandler struct contains history service
type HistoryHandler struct {
	historyService *svc.HistoryService
}

// NewHandler initializes handler with history service
func NewHistoryHandler() *HistoryHandler {
	return &HistoryHandler{
		historyService: svc.NewHistoryService(),
	}
}

// apiToSvcHistory converts API model to service model
func (h *HistoryHandler) apiToSvcHistory(historyAPI *HistoryAPI) *svc.History {
	return &svc.History{
		Id:           historyAPI.Id,
		Username:     historyAPI.Username,
		UserIp:       historyAPI.UserIp,
		Command:      historyAPI.Command,
		NeName:       historyAPI.NeName,
		Result:       historyAPI.Result,
		ExecutedTime: historyAPI.ExecutedTime,
		Mode:         historyAPI.Mode,
	}
}

// GetHistoryHandler handles getting history list by mode and limit
func (h *HistoryHandler) GetHistoryHandler(c *fiber.Ctx) error {
	var historyReq HistoryRequest
	err := c.BodyParser(&historyReq)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")

	// Create history command for tracking
	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Get history command mode " + historyReq.Mode,
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler get history list")

	// Get history list from service
	histories, err := h.historyService.GetListHistoryByMode(historyReq.Mode, historyReq.Limit)
	if err != nil {
		logger.Logger.Error("Cannot get history list, err: ", err)
		response.BadRequest(c, "cannot get history list")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, histories)
	logger.Logger.Info("Get history list success")
	return nil
}

// SaveHistoryHandler handles saving history command
func (h *HistoryHandler) SaveHistoryHandler(c *fiber.Ctx) error {
	var historyAPI HistoryAPI
	err := c.BodyParser(&historyAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	logger.Logger.Info("Saving history command: ", historyAPI.Command)

	// Convert API model to service model
	historySVC := h.apiToSvcHistory(&historyAPI)

	// Save history based on result
	if historySVC.Result {
		h.historyService.SaveHistoryCommandSuccess(historySVC)
	} else {
		h.historyService.SaveHistoryCommandFailure(historySVC)
	}

	return nil
}
