package svc

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
)

type UserNe struct {
	Id     uint
	UserId uint
	NeId   uint
}

type UserNeService struct {
	repo repository.IDatabaseStore
}

func NewUserNeService() *UserNeService {
	return &UserNeService{
		repo: repository.GetSingleton(),
	}
}

func (s *UserNeService) toDb(userNeSVC *UserNe) *models_db.UserNe {
	return &models_db.UserNe{
		ID:     userNeSVC.Id,
		UserID: userNeSVC.UserId,
		NeID:   userNeSVC.NeId,
	}
}

func (s *UserNeService) fromDb(userNeDB *models_db.UserNe) *UserNe {
	return &UserNe{
		Id:     userNeDB.ID,
		UserId: userNeDB.UserID,
		NeId:   userNeDB.NeID,
	}
}

func (s *UserNeService) NeUserAdd(userId, neId uint) error {
	userNeDb := &models_db.UserNe{
		UserID: userId,
		NeID:   neId,
	}
	err := s.repo.UserNeAdd(userNeDb)
	if err != nil {
		logger.Logger.Error("Cannot add user ne into database: ", err)
		return err
	}
	return nil
}

func (s *UserNeService) NeUserDelete(userId, neId uint) error {
	err := s.repo.UserNeDelete(userId, neId)
	if err != nil {
		logger.Logger.Error("Cannot delete user ne from database: ", err)
		return err
	}
	return nil
}

func (s *UserNeService) NeUserGet(userId, neId uint) (*UserNe, error) {
	logger.Logger.Infof("get user ne with userId %d and neId %d", userId, neId)
	userNeDb, err := s.repo.UserNeGet(userId, neId)
	if err != nil {
		logger.Logger.Error("Cannot get user ne from database: ", err)
		return nil, err
	}
	logger.Logger.Infof("get user ne with userId %d and neId %d success", userId, neId)
	return s.fromDb(userNeDb), nil
}

func (s *UserNeService) GetListNeByUser(username string) ([]*NetworkElement, error) {
	logger.Logger.Info("get list ne of user: ", username)
	nesDb, err := s.repo.GetNetworkElementByUserName(username)
	if err != nil {
		logger.Logger.Error("Cannot get ne from database: ", err)
		return nil, err
	}
	logger.Logger.Infof("get list ne of user %s success", username)

	neService := NewNetworkElementService()
	var nesSVC []*NetworkElement
	for _, neDb := range nesDb {
		nesSVC = append(nesSVC, neService.fromDb(&neDb))
	}

	return nesSVC, nil
}
