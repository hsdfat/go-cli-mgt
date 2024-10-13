package unit

import (
	"errors"
	"github.com/stretchr/testify/require"
	"go-cli-mgt/pkg/models/models_error"
	userService "go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/random"
	"testing"
)

func TestGetUserNe(t *testing.T) {

	userId, neId := uint(1), uint(1)
	userNeGetTest, err := userService.NeUserGet(userId, neId)
	require.NoError(t, err)
	require.NotEmpty(t, userNeGetTest)
	require.Equal(t, userId, userNeGetTest.UserId)
	require.Equal(t, neId, userNeGetTest.NeId)
}

func TestCreateAndDeleteUserNe(t *testing.T) {

	// Create User Ne
	userId, neId := uint(random.IntRandom(1, 10)), uint(random.IntRandom(1, 10))
	err := userService.NeUserAdd(userId, neId)
	require.NoError(t, err)

	// Get User Ne
	userNeGetTest, err := userService.NeUserGet(userId, neId)
	require.NoError(t, err)
	require.NotEmpty(t, userNeGetTest)
	require.Equal(t, userId, userNeGetTest.UserId)
	require.Equal(t, neId, userNeGetTest.NeId)

	// Delete User Ne
	err = userService.NeUserDelete(userId, neId)
	require.NoError(t, err)

	userNeGetTest, err = userService.NeUserGet(userId, neId)
	if !errors.Is(err, models_error.ErrNotFoundUserNe) {
		require.Error(t, errors.New("delete user ne un success"))
	}
}
