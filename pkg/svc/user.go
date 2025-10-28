package svc

import (
	"time"

	"github.com/hsdfat/go-cli-mgt/pkg/logger"
	models_db "github.com/hsdfat/go-cli-mgt/pkg/models/db"
	"github.com/hsdfat/go-cli-mgt/pkg/store/repository"
	"github.com/hsdfat/go-cli-mgt/pkg/utils/bcrypt"
)

type User struct {
	Id           uint
	Username     string
	Password     string
	Active       bool
	Email        string
	CreatedDate  uint64
	DisableDate  uint64
	DeActivateBy string
}

type UserService struct {
	repo repository.IDatabaseStore
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.GetSingleton(),
	}
}

// Convert từ SVC model sang DB model
func (s *UserService) toDb(userSVC *User) *models_db.User {
	return &models_db.User{
		ID:              userSVC.Id,
		Username:        userSVC.Username,
		Password:        userSVC.Password,
		Active:          userSVC.Active,
		Email:           userSVC.Email,
		CreatedDateUnix: userSVC.CreatedDate,
		DisableDateUnix: userSVC.DisableDate,
		DeactivateBy:    userSVC.DeActivateBy,
	}
}

// Convert từ DB model sang SVC model
func (s *UserService) fromDb(userDB *models_db.User) *User {
	return &User{
		Id:           userDB.ID,
		Username:     userDB.Username,
		Password:     userDB.Password,
		Active:       userDB.Active,
		Email:        userDB.Email,
		CreatedDate:  uint64(userDB.CreatedDateUnix),
		DisableDate:  uint64(userDB.DisableDateUnix),
		DeActivateBy: userDB.DeactivateBy,
	}
}

func (s *UserService) CreateProfile(user *User) error {
	userDb, err := s.repo.GetUserByUsername(user.Username)
	if err != nil && err.Error() != "user not found" {
		logger.Logger.Error("Cannot get user by username from db, username: ", user.Username, " err: ", err)
		return err
	}

	if userDb != nil && !userDb.Active {
		logger.Logger.Info("user disable, reactivating...")
		userDb.Active = true
		userDb.Password = bcrypt.Encode(user.Username + user.Password)
		err = s.repo.UpdateUser(userDb)
		if err != nil {
			logger.Logger.Error("Cannot update user to database: ", err)
			return err
		}
		return nil
	}

	userDbNew := s.toDb(user)
	userDbNew.Active = true
	userDbNew.CreatedDateUnix = uint64(time.Now().Unix())
	userDbNew.DisableDateUnix = 1
	userDbNew.Password = bcrypt.Encode(user.Username + user.Password)

	err = s.repo.CreateUser(userDbNew)
	if err != nil {
		logger.Logger.Error("Cannot create user with username: ", user.Username)
		return err
	}

	return nil
}

func (s *UserService) DeleteProfile(username string) error {
	_, err := s.repo.GetUserByUsername(username)
	if err != nil {
		logger.Logger.Error("Cannot get user by username from database: ", err)
		return err
	}
	err = s.repo.DeleteUser(username)
	if err != nil {
		logger.Logger.Error("Cannot delete user by username from database: ", err)
		return err
	}
	logger.Logger.Info("Delete user complete, username: ", username)
	return nil
}

func (s *UserService) DisableProfile(username string, userDeactivate string) error {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		logger.Logger.Error("Cannot get user by username from database: ", err)
		return err
	}

	if !user.Active {
		return nil
	}

	user.Active = false
	user.DisableDateUnix = uint64(time.Now().Unix())
	user.DeactivateBy = userDeactivate
	err = s.repo.UpdateUser(user)
	if err != nil {
		logger.Logger.Error("Cannot update user to database: ", err)
		return err
	}
	return nil
}

func (s *UserService) GetProfileByUsername(username string) (*User, error) {
	logger.Logger.Info("Get profile username: ", username)
	userDb, err := s.repo.GetUserByUsername(username)
	if err != nil {
		logger.Logger.Error("Cannot get user by username from database: ", err)
		return nil, err
	}
	logger.Logger.Infof("Get profile success with username %s and id %d ", username, userDb.ID)
	return s.fromDb(userDb), nil
}

func (s *UserService) GetListProfile() ([]*User, error) {
	usersDb, err := s.repo.ListUsers()
	if err != nil {
		logger.Logger.Error("Cannot get list user from database: ", err)
		return nil, err
	}

	var usersSVC []*User
	for _, userDb := range usersDb {
		usersSVC = append(usersSVC, s.fromDb(&userDb))
	}

	return usersSVC, nil
}

func (s *UserService) UpdateProfile(user *User) error {
	userDb := s.toDb(user)
	err := s.repo.UpdateUser(userDb)
	if err != nil {
		logger.Logger.Error("Cannot update user to database: ", err)
		return err
	}
	return nil
}

func (s *UserService) UpdatePassword(user *User) error {
	logger.Logger.Info("update password for user: ", user.Username)
	userDb := s.toDb(user)
	s.repo.UpdatePasswordUser(userDb)
	return nil
}
