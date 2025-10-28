package svc

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
)

// ===== MODEL =====
type UserRole struct {
	Id     uint
	UserId uint
	RoleId uint
}

// ===== SERVICE =====
type UserRoleService struct {
	repo repository.IDatabaseStore
}

func NewUserRoleService() *UserRoleService {
	return &UserRoleService{
		repo: repository.GetSingleton(),
	}
}

func (s *UserRoleService) toDb(userRoleSVC *UserRole) *models_db.UserRole {
	return &models_db.UserRole{
		ID:     userRoleSVC.Id,
		UserID: userRoleSVC.UserId,
		RoleID: userRoleSVC.RoleId,
	}
}

func (s *UserRoleService) fromDb(userRoleDB *models_db.UserRole) *UserRole {
	return &UserRole{
		Id:     userRoleDB.ID,
		UserId: userRoleDB.UserID,
		RoleId: userRoleDB.RoleID,
	}
}

func (s *UserRoleService) RoleUserAdd(userId, roleId uint) error {
	userRoleDb := &models_db.UserRole{
		UserID: userId,
		RoleID: roleId,
	}
	err := s.repo.UserRoleAdd(userRoleDb)
	if err != nil {
		logger.Logger.Error("Cannot add user role into database: ", err)
		return err
	}
	return nil
}

func (s *UserRoleService) RoleUserDelete(userId, roleId uint) {
	logger.Logger.Infof("Delete user id %d with role id %d", userId, roleId)
	s.repo.UserRoleDelete(userId, roleId)
}

func (s *UserRoleService) RoleUserGet(userId, roleId uint) (*UserRole, error) {
	logger.Logger.Infof("get user role with userId %d and roleId %d", userId, roleId)
	userRoleDb, err := s.repo.UserRoleGet(userId, roleId)
	if err != nil {
		logger.Logger.Error("Cannot get user role from database: ", err)
		return nil, err
	}
	logger.Logger.Infof("get user role with userId %d and role %d success", userId, roleId)
	return s.fromDb(userRoleDb), nil
}
