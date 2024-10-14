package unit

import (
	"errors"
	"github.com/stretchr/testify/require"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_error"
	"go-cli-mgt/pkg/service/role"
	"go-cli-mgt/pkg/service/utils/random"
	"testing"
)

func TestCreateAndUpdateAndDeleteRole(t *testing.T) {
	// Create role test
	roleTest := models_api.Role{
		RoleName:    random.StringRandom(20),
		Priority:    random.StringRandom(20),
		Description: random.StringRandom(30),
	}

	// Create role
	err := role.CreateRole(&roleTest)
	require.NoError(t, err)

	// Get Role
	roleGetTest, err := role.GetRoleByName(roleTest.RoleName)
	require.NoError(t, err)
	require.NotEmpty(t, roleGetTest)
	require.NotZero(t, roleGetTest.RoleId)
	require.Equal(t, roleTest.RoleName, roleGetTest.RoleName)
	require.Equal(t, roleTest.Priority, roleGetTest.Priority)
	require.Equal(t, roleTest.Description, roleGetTest.Description)

	// Update Role
	roleTest.Description = random.StringRandom(20)
	role.UpdateRole(&roleTest)

	// Get Role
	roleGetTest, err = role.GetRoleByName(roleTest.RoleName)
	require.NoError(t, err)
	require.NotEmpty(t, roleGetTest)
	require.NotZero(t, roleGetTest.RoleId)
	require.Equal(t, roleTest.RoleName, roleGetTest.RoleName)
	require.Equal(t, roleTest.Priority, roleGetTest.Priority)
	require.Equal(t, roleTest.Description, roleGetTest.Description)

	// Delete Role
	err = role.DeleteRole(&roleTest)
	require.NoError(t, err)

	// Get role again to test
	roleGetTest, err = role.GetRoleByName(roleTest.RoleName)
	if errors.Is(err, models_error.ErrNotFoundRole) == false {
		require.Error(t, errors.New("delete user un success"))
	}
}

func TestGetRole(t *testing.T) {
	roleName := "admin"

	roleGetTest, err := role.GetRoleByName(roleName)
	require.NoError(t, err)
	require.NotEmpty(t, roleGetTest)
	require.NotZero(t, roleGetTest.RoleId)
	require.Equal(t, roleName, roleGetTest.RoleName)
}
