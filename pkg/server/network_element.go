package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	"github.com/hsdfat/go-cli-mgt/pkg/svc"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/response"
)

type NeData struct {
	NeId             uint   `json:"ne-id,omitempty"`
	Name             string `json:"name,omitempty"`
	Type             string `json:"type,omitempty"`
	MasterIpConfig   string `json:"master-ip-config,omitempty"`
	MasterPortConfig string `json:"master-port-config,omitempty"`
	SlaveIpConfig    string `json:"slave-ip-config,omitempty"`
	SlavePortConfig  string `json:"slave-port-config,omitempty"`
	IpCommand        string `json:"ip-command,omitempty"`
	PortCommand      string `json:"port-command,omitempty"`
	Description      string `json:"description,omitempty"`
	Namespace        string `json:"namespace,omitempty"`
}

type NetworkElementHandler struct {
	neService      *svc.NetworkElementService
	historyService *svc.HistoryService
}

func NewNetworkElementHandlerHandler() *NetworkElementHandler {
	return &NetworkElementHandler{
		neService:      svc.NewNetworkElementService(),
		historyService: svc.NewHistoryService(),
	}
}

func (h *NetworkElementHandler) apiToSvcNe(neAPI *NeData) *svc.NetworkElement {
	return &svc.NetworkElement{
		NeId:             neAPI.NeId,
		Name:             neAPI.Name,
		Type:             neAPI.Type,
		MasterIpConfig:   neAPI.MasterIpConfig,
		MasterPortConfig: neAPI.MasterPortConfig,
		SlaveIpConfig:    neAPI.SlaveIpConfig,
		SlavePortConfig:  neAPI.SlavePortConfig,
		IpCommand:        neAPI.IpCommand,
		PortCommand:      neAPI.PortCommand,
		Description:      neAPI.Description,
		Namespace:        neAPI.Namespace,
	}
}

func (h *NetworkElementHandler) ListNetworkElementHandler(c *fiber.Ctx) error {
	username := c.Get("username")

	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Get list ne",
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Info("Handler get ne list")
	neList, err := h.neService.GetListNetworkElement()
	if err != nil {
		logger.Logger.Error("Error get list ne: ", err)
		response.InternalError(c, "Error get list ne")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	response.Write(c, neList)
	return nil
}

func (h *NetworkElementHandler) CreateOrUpdateHandler(c *fiber.Ctx) error {
	var neAPI NeData
	err := c.BodyParser(&neAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")

	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Create network element " + neAPI.Name + " namespace " + neAPI.Namespace,
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Infof("Handler create ne %s namespace %s", neAPI.Name, neAPI.Namespace)

	neSVC := h.apiToSvcNe(&neAPI)
	err = h.neService.CreateNetworkElement(neSVC)
	if err != nil {
		logger.Logger.Error("Error create ne: ", err)
		response.InternalError(c, "Error create ne")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	logger.Logger.Infof("Create ne %s namespace %s success", neAPI.Name, neAPI.Namespace)
	return nil
}

func (h *NetworkElementHandler) DeleteHandler(c *fiber.Ctx) error {
	var neAPI NeData
	err := c.BodyParser(&neAPI)
	if err != nil {
		logger.Logger.Error("Error parsing JSON request body: ", err)
		response.InternalError(c, "Error parsing JSON request body")
		return err
	}

	username := c.Get("username")

	historyCommand := &svc.History{
		Username: username,
		UserIp:   c.IP(),
		Command:  "Delete network element " + neAPI.Name + " namespace " + neAPI.Namespace,
		NeName:   "",
		Mode:     "cli-config",
	}

	logger.Logger.Infof("Handler delete ne %s namespace %s", neAPI.Name, neAPI.Namespace)

	err = h.neService.DeleteNetworkElement(neAPI.Name, neAPI.Namespace)
	if err != nil {
		logger.Logger.Error("Error delete ne: ", err)
		response.InternalError(c, "Error delete ne")
		h.historyService.SaveHistoryCommandFailure(historyCommand)
		return err
	}

	h.historyService.SaveHistoryCommandSuccess(historyCommand)
	logger.Logger.Infof("Delete ne %s namespace %s success", neAPI.Name, neAPI.Namespace)
	return nil
}
