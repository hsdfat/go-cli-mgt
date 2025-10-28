package svc

import (
	"errors"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/bcrypt"
)

// Auth không cần model riêng vì không có struct phức tạp

// ===== SERVICE =====
type AuthService struct {
	repo repository.IDatabaseStore
}

func NewAuthService() *AuthService {
	return &AuthService{
		repo: repository.GetSingleton(),
	}
}

func (s *AuthService) Login(username, password string) (bool, error, uint) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		logger.Logger.Error("Cannot get user by username from db, username: ", username)
		return false, err, 0
	}

	if !user.Active {
		logger.Logger.Info("username: ", username, " does not active")
		return false, errors.New("user not found"), 0
	}

	if !bcrypt.Matches(username+password, user.Password) {
		logger.Logger.Info("Login fail, wrong username or password")
		return false, nil, 0
	}

	return true, nil, user.ID
}

func (s *AuthService) GetRole(userId uint) (string, error) {
	roleList, err := s.repo.GetRoleByUserId(userId)
	if err != nil {
		logger.Logger.Error("Cannot get role by user id: ", err)
		return "", err
	}
	roleStr := ""
	for _, role := range roleList {
		roleStr = roleStr + role.RoleName + " "
	}
	return roleStr, nil
}
