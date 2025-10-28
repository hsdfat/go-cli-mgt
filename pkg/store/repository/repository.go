package repository

import (
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"
)

type IDatabaseStore interface {
	Ping() error

	IUserRepository
	IRoleRepository
	IHistoryRepository
	INetworkElementRepository
	IUserNeRepository
	IUserRoleRepository
	ILoginRepository
	IMmeCommandRepository
}

type IUserRepository interface {
	CreateUser(userDB *models_db.User) error
	DeleteUser(username string) error
	UpdateUser(userUpdate *models_db.User) error
	UpdatePasswordUser(userDB *models_db.User)
	GetUserByID(id uint) (*models_db.User, error)
	GetUserByUsername(username string) (*models_db.User, error)
	ListUsers() ([]models_db.User, error)
}

type IRoleRepository interface {
	GetRoleByUserId(userId uint) ([]models_db.Role, error)
	GetRoleByName(roleName string) (*models_db.Role, error)
	GetListRole() ([]models_db.Role, error)
	CreateRole(roleDB *models_db.Role) error
	DeleteRole(roleDB *models_db.Role) error
	UpdateRole(roleDB *models_db.Role) error
}

type INetworkElementRepository interface {
	CreateNetworkElement(neDB *models_db.NetworkElement) error
	DeleteNetworkElementByName(neName string, namespace string) error
	GetNetworkElementByName(neName string, namespace string) (*models_db.NetworkElement, error)
	GetListNetworkElement() ([]models_db.NetworkElement, error)
	GetNetworkElementByUserName(userName string) ([]models_db.NetworkElement, error)
}

type IHistoryRepository interface {
	SaveHistory(historyDB *models_db.OperationHistory) error
	GetHistoryById(id uint64) (*models_db.OperationHistory, error)
	DeleteHistoryById(id uint64) error
	GetHistoryListByMode(mode string) ([]models_db.OperationHistory, error)
	GetRecordHistoryByCommand(command string) (*models_db.OperationHistory, error)
	GetHistoryCommandByModeLimit(mode string, limit int) ([]models_db.OperationHistory, error)
	GetHistorySavingLog(neSiteName string) ([]models_db.OperationHistory, error)
}

type IUserNeRepository interface {
	UserNeAdd(userNeDB *models_db.UserNe) error
	UserNeDelete(userId, neId uint) error
	UserNeGet(userId, neId uint) (*models_db.UserNe, error)
}

type IUserRoleRepository interface {
	UserRoleAdd(userRoleDB *models_db.UserRole) error
	UserRoleGet(userId, roleId uint) (*models_db.UserRole, error)
	UserRoleDelete(userId, roleId uint)
}

type ILoginRepository interface {
}

type IMmeCommandRepository interface {
}
