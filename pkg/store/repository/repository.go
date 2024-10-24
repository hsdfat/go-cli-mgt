package repository

import (
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_config"
	"go-cli-mgt/pkg/models/models_db"
)

type DatabaseStore interface {
	Init(cfg models_config.DatabaseConfig) error

	UserRepository
	LoginRepository
	HistoryRepository
	RoleRepository
	NetworkElementRepository
	MmeCommandRepository
}

type UserRepository interface {
	CreateUser(*models_api.User) error
	GetUserByID(uint) (*models_api.User, error)
	ListUsers() ([]models_api.User, error)
	GetUserByUsername(string) (*models_api.User, error)
	DeleteUser(string) error
	UpdateUser(*models_api.User) error
	UpdatePasswordUser(*models_api.User)

	UserNeAdd(*models_api.UserNe) error
	UserNeDelete(uint, uint) error
	UserNeGet(uint, uint) (*models_api.UserNe, error)

	UserRoleAdd(*models_api.UserRole) error
	UserRoleGet(userId, roleId uint) (*models_api.UserRole, error)
	UserRoleDelete(userId, roleId uint)
}

type RoleRepository interface {
	GetRoleByUserId(uint) ([]models_db.Role, error)
	GetRoleByName(string) (*models_api.Role, error)
	GetListRole() ([]models_api.Role, error)
	CreateRole(*models_api.Role) error
	DeleteRole(*models_api.Role) error
	UpdateRole(*models_api.Role) error
}

type NetworkElementRepository interface {
	CreateNetworkElement(*models_api.NeData) error
	DeleteNetworkElementByName(string, string) error
	GetNetworkElementByName(string, string) (*models_api.NeData, error)
	GetListNetworkElement() ([]models_api.NeData, error)
	GetNetworkElementByUserName(string) ([]models_api.NeData, error)
}

type HistoryRepository interface {
	SaveHistory(*models_api.History) error
	GetHistoryById(uint64) (*models_api.History, error)
	DeleteHistoryById(uint64) error
	GetHistoryListByMode(string) ([]models_api.History, error)
	GetRecordHistoryByCommand(string) (*models_api.History, error)
	GetHistoryCommandByModeLimit(string, int) ([]models_api.History, error)
	GetHistorySavingLog(neSiteName string) ([]models_api.History, error)
}

type LoginRepository interface {
}

type MmeCommandRepository interface {
}
