package svc

import (
	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
)

type Role struct {
	RoleId      uint
	RoleName    string
	Description string
}

type RoleService struct {
	repo repository.IDatabaseStore
}

func NewRoleService() *RoleService {
	return &RoleService{
		repo: repository.GetSingleton(),
	}
}

func (s *RoleService) toDb(roleSVC *Role) *models_db.Role {
	return &models_db.Role{
		ID:          roleSVC.RoleId,
		RoleName:    roleSVC.RoleName,
		Description: roleSVC.Description,
	}
}

func (s *RoleService) fromDb(roleDB *models_db.Role) *Role {
	return &Role{
		RoleId:      roleDB.ID,
		RoleName:    roleDB.RoleName,
		Description: roleDB.Description,
	}
}

func (s *RoleService) CreateRole(role *Role) error {
	logger.Logger.Info("Create role: ", role.RoleName)
	roleDb := s.toDb(role)
	err := s.repo.CreateRole(roleDb)
	if err != nil {
		logger.Logger.Error("Cannot create role, err: ", err)
		return err
	}
	logger.Logger.Infof("Create role %s success ", role.RoleName)
	return nil
}

func (s *RoleService) DeleteRole(role *Role) error {
	logger.Logger.Info("Delete role: ", role.RoleName)
	roleDb := s.toDb(role)
	err := s.repo.DeleteRole(roleDb)
	if err != nil {
		logger.Logger.Error("Cannot delete role, err: ", err)
		return err
	}
	logger.Logger.Infof("Delete role %s success ", role.RoleName)
	return nil
}

func (s *RoleService) GetRoleByName(roleName string) (*Role, error) {
	logger.Logger.Info("Get role from database with role name ", roleName)
	roleDb, err := s.repo.GetRoleByName(roleName)
	if err != nil {
		logger.Logger.Error("Cannot get role by role name from database, err: ", err)
		return nil, err
	}
	logger.Logger.Info("Get role success from database with role name ", roleName)
	return s.fromDb(roleDb), nil
}

func (s *RoleService) GetListRole() ([]*Role, error) {
	rolesDb, err := s.repo.GetListRole()
	if err != nil {
		logger.Logger.Error("Cannot get role list, err: ", err)
		return nil, err
	}

	var rolesSVC []*Role
	for _, roleDb := range rolesDb {
		rolesSVC = append(rolesSVC, s.fromDb(&roleDb))
	}

	return rolesSVC, nil
}

func (s *RoleService) UpdateRole(role *Role) error {
	logger.Logger.Info("Update role from database with role name ", role.RoleName)
	logger.Logger.Debug("Description: ", role.Description)
	roleDb := s.toDb(role)
	err := s.repo.UpdateRole(roleDb)
	if err != nil {
		logger.Logger.Error("Cannot update role, err: ", err)
		return err
	}
	return nil
}
