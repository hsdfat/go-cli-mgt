package svc

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
)

type NetworkElement struct {
	NeId             uint
	Name             string
	Type             string
	MasterIpConfig   string
	MasterPortConfig string
	SlaveIpConfig    string
	SlavePortConfig  string
	IpCommand        string
	PortCommand      string
	Description      string
	Namespace        string
}

type NetworkElementService struct {
	repo repository.IDatabaseStore
}

func NewNetworkElementService() *NetworkElementService {
	return &NetworkElementService{
		repo: repository.GetSingleton(),
	}
}

func (s *NetworkElementService) toDb(neSVC *NetworkElement) *models_db.NetworkElement {
	return &models_db.NetworkElement{
		ID:               neSVC.NeId,
		Name:             neSVC.Name,
		Type:             neSVC.Type,
		Namespace:        neSVC.Namespace,
		MasterIpConfig:   neSVC.MasterIpConfig,
		MasterPortConfig: neSVC.MasterPortConfig,
		SlaveIpConfig:    neSVC.SlaveIpConfig,
		SlavePortConfig:  neSVC.SlavePortConfig,
		IpCommand:        neSVC.IpCommand,
		PortCommand:      neSVC.PortCommand,
	}
}

func (s *NetworkElementService) fromDb(neDB *models_db.NetworkElement) *NetworkElement {
	return &NetworkElement{
		NeId:             neDB.ID,
		Name:             neDB.Name,
		Type:             neDB.Type,
		Namespace:        neDB.Namespace,
		MasterIpConfig:   neDB.MasterIpConfig,
		MasterPortConfig: neDB.MasterPortConfig,
		SlaveIpConfig:    neDB.SlaveIpConfig,
		SlavePortConfig:  neDB.SlavePortConfig,
		IpCommand:        neDB.IpCommand,
		PortCommand:      neDB.PortCommand,
	}
}

func (s *NetworkElementService) CreateNetworkElement(ne *NetworkElement) error {
	neDb := s.toDb(ne)
	err := s.repo.CreateNetworkElement(neDb)
	if err != nil {
		logger.Logger.Error("Cannot create network element, err: ", err)
		return err
	}
	return nil
}

func (s *NetworkElementService) DeleteNetworkElement(neName string, namespace string) error {
	err := s.repo.DeleteNetworkElementByName(neName, namespace)
	if err != nil {
		logger.Logger.Error("Cannot delete network element, err: ", err)
		return err
	}
	return nil
}

func (s *NetworkElementService) GetNetworkElement(neName, namespace string) (*NetworkElement, error) {
	logger.Logger.Infof("Get network element info with ne %s namespace %s", neName, namespace)
	neDb, err := s.repo.GetNetworkElementByName(neName, namespace)
	if err != nil {
		logger.Logger.Error("Cannot get ne, err: ", err)
		return nil, err
	}
	logger.Logger.Infof("Get network element info with ne %s namespace %s success with id %d", neName, namespace, neDb.ID)
	return s.fromDb(neDb), nil
}

func (s *NetworkElementService) GetListNetworkElement() ([]*NetworkElement, error) {
	nesDb, err := s.repo.GetListNetworkElement()
	if err != nil {
		logger.Logger.Error("Cannot get ne list, err: ", err)
		return nil, err
	}

	var nesSVC []*NetworkElement
	for _, neDb := range nesDb {
		nesSVC = append(nesSVC, s.fromDb(&neDb))
	}

	return nesSVC, nil
}
