package unit_test

import (
	"errors"
	"github.com/stretchr/testify/require"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/bcrypt"
	"go-cli-mgt/pkg/service/utils/random"
	"testing"
)

func TestCreateProfile(t *testing.T) {
	userTest := models_api.User{
		Username: random.StringRandom(10),
		Password: random.StringRandom(20),
	}

	err := user.CreateProfile(userTest)
	require.NoError(t, err)

	userGetTest, err := user.GetProfile(userTest.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userGetTest)
	require.Equal(t, userTest.Username, userGetTest.Username)
	if bcrypt.Matches(userTest.Username+userTest.Password, userGetTest.Password) {
		require.Error(t, errors.New("password in correct"))
	}
	require.NotZero(t, userGetTest.Id)
	require.NotZero(t, userGetTest.CreatedDate)
}
