package unit_test

import (
	"errors"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/bcrypt"
	"go-cli-mgt/pkg/service/utils/random"
	"go-cli-mgt/pkg/store/postgres"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAndDeleteProfile(t *testing.T) {
	// Create userTest
	userTest := models_api.User{
		Username: random.StringRandom(10),
		Password: random.StringRandom(20),
		Email:    random.StringRandom(10),
	}

	// Create User
	err := user.CreateProfile(userTest)
	require.NoError(t, err)

	// Try to get user
	userGetTest, err := user.GetProfileByUsername(userTest.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userGetTest)
	require.Equal(t, userTest.Username, userGetTest.Username)
	if bcrypt.Matches(userTest.Username+userTest.Password, userGetTest.Password) {
		require.Error(t, errors.New("password in correct"))
	}
	require.NotZero(t, userGetTest.Id)
	require.NotZero(t, userGetTest.CreatedDate)

	// Delete User
	err = user.DeleteProfile(userTest.Username)
	require.NoError(t, err)

	// Try to get user again, must be return an error
	userGetTest, err = user.GetProfileByUsername(userTest.Username)
	if errors.Is(err, postgres.ErrNotFoundUser) == false {
		require.Error(t, errors.New("delete user un success"))
	}
}

func TestGetUser(t *testing.T) {
	userTest := models_api.User{
		Username: "userTest1",
		Password: "userTest1",
		Email:    "userTest1",
	}

	userGetTest, err := user.GetProfileByUsername(userTest.Username)
	require.NoError(t, err)
	require.NotZero(t, userGetTest.Id)
	require.NotEmpty(t, userGetTest)
	require.Equal(t, userTest.Username, userGetTest.Username)
	require.Equal(t, userTest.Email, userGetTest.Email)
	require.Equal(t, userGetTest.Active, true)
	if bcrypt.Matches(userTest.Username+userTest.Password, userGetTest.Password) {
		require.Error(t, errors.New("password in correct"))
	}
}

func TestDisableAndEnableUser(t *testing.T) {
	userTest := models_api.User{
		Username: "userTest1",
		Password: "userTest1",
		Email:    "userTest1",
	}

	err := user.DisableProfile(userTest.Username)
	require.NoError(t, err)
	userGetTest, err := user.GetProfileByUsername(userTest.Username)
	require.NoError(t, err)
	require.NotZero(t, userGetTest.Id)
	require.NotEmpty(t, userGetTest)
	require.Equal(t, userTest.Username, userGetTest.Username)
	require.Equal(t, userTest.Email, userGetTest.Email)
	require.Equal(t, userGetTest.Active, false)
	if bcrypt.Matches(userTest.Username+userTest.Password, userGetTest.Password) {
		require.Error(t, errors.New("password in correct"))
	}

	userGetTest.Active = true
	err = user.UpdateProfile(userGetTest)
	require.NoError(t, err)

	userGetTest2, err := user.GetProfileByUsername(userTest.Username)
	require.NoError(t, err)
	require.NotZero(t, userGetTest2.Id)
	require.NotEmpty(t, userGetTest2)
	require.Equal(t, userTest.Username, userGetTest2.Username)
	require.Equal(t, userTest.Email, userGetTest2.Email)
	require.Equal(t, userGetTest2.Active, true)
	if bcrypt.Matches(userTest.Username+userTest.Password, userGetTest2.Password) {
		require.Error(t, errors.New("password in correct"))
	}
}
