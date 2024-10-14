package unit

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_error"
	"go-cli-mgt/pkg/server"
	"go-cli-mgt/pkg/service/network_elements"
	"go-cli-mgt/pkg/service/role"
	"go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/bcrypt"
	"go-cli-mgt/pkg/service/utils/random"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test01 Testing login, create and disable user
func Test01(t *testing.T) {
	// Create server
	httpApp := server.Initialize("E:\\Data\\Go\\go-cli-mgt\\.env")

	// Login
	// WARNING: This user must have been in db before
	userReq := models_api.RequestUser{
		Username: "userTest1",
		Password: "userTest1",
	}
	tokenStr := Login(t, userReq, httpApp)

	// Create user
	userTest := models_api.User{
		Username: random.StringRandom(10),
		Password: random.StringRandom(20),
		Email:    random.StringRandom(10),
	}
	CreateUser(t, userTest, tokenStr, httpApp)

	// Deactivate user
	userReqTest, err := json.Marshal(userTest)
	req := httptest.NewRequest(http.MethodDelete, "/mgt-svc/v1/user/profile", bytes.NewBuffer(userReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)
	resp, err := httpApp.Test(req)
	require.NoError(t, err)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User from db to check
	userGetTest, err := user.GetProfileByUsername(userTest.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userGetTest)
	require.Equal(t, userTest.Username, userGetTest.Username)
	require.Equal(t, false, userGetTest.Active)
	if bcrypt.Matches(userTest.Username+userTest.Password, userGetTest.Password) {
		require.Error(t, errors.New("password in correct"))
	}
	require.NotZero(t, userGetTest.Id)
	require.NotZero(t, userGetTest.CreatedDate)
	require.NotZero(t, userGetTest.DisableDate)

	// Test complete, delete user test
	DeleteUser(t, userTest.Username)
}

// Test02 Testing login, Create ne and delete ne
func Test02(t *testing.T) {
	// Create server
	httpApp := server.Initialize("E:\\Data\\Go\\go-cli-mgt\\.env")

	// Login
	// WARNING: This user must have been in db before
	userReq := models_api.RequestUser{
		Username: "userTest1",
		Password: "userTest1",
	}
	tokenStr := Login(t, userReq, httpApp)

	// Create ne
	neTest := models_api.NeData{
		Name:             random.StringRandom(10),
		Url:              random.StringRandom(10),
		Type:             random.StringRandom(10),
		MasterIpConfig:   random.Ipv4Random(),
		MasterPortConfig: int32(random.IntRandom(0, 9999)),
		SlaveIpConfig:    random.Ipv4Random(),
		SlavePortConfig:  int32(random.IntRandom(0, 9999)),
		IpCommand:        random.Ipv4Random(),
		PortCommand:      int32(random.IntRandom(0, 9999)),
		Description:      random.StringRandom(10),
		Namespace:        random.StringRandom(10),
	}

	neReqTest, err := json.Marshal(neTest)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/network-element", bytes.NewBuffer(neReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err := httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusNotFound)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get Ne from db
	neGetTest, err := network_elements.GetNetworkElement(neTest.Name, neTest.Namespace)
	require.NoError(t, err)
	require.NotEmpty(t, neGetTest)
	require.Equal(t, neTest.Name, neGetTest.Name)
	require.Equal(t, neTest.Url, neGetTest.Url)
	require.Equal(t, neTest.Type, neGetTest.Type)
	require.Equal(t, neTest.MasterIpConfig, neGetTest.MasterIpConfig)
	require.Equal(t, neTest.MasterPortConfig, neGetTest.MasterPortConfig)
	require.Equal(t, neTest.SlaveIpConfig, neGetTest.SlaveIpConfig)
	require.Equal(t, neTest.SlavePortConfig, neGetTest.SlavePortConfig)
	require.Equal(t, neTest.IpCommand, neGetTest.IpCommand)
	require.Equal(t, neTest.PortCommand, neGetTest.PortCommand)
	require.Equal(t, neTest.Namespace, neGetTest.Namespace)

	// Delete ne
	req = httptest.NewRequest(http.MethodDelete, "/mgt-svc/v1/network-element", bytes.NewBuffer(neReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get Me from db to Check
	neGetTest, err = network_elements.GetNetworkElement(neTest.Name, neTest.Namespace)
	if errors.Is(err, models_error.ErrNotFoundNe) == false {
		require.Error(t, errors.New("delete ne un success"))
	}
}

// Test03 testing login, create user, create ne, add user ne, delete user ne, delete ne, delete user
func Test03(t *testing.T) {
	// Create server
	httpApp := server.Initialize("E:\\Data\\Go\\go-cli-mgt\\.env")

	// Login
	// WARNING: This user must have been in db before
	userReq := models_api.RequestUser{
		Username: "userTest1",
		Password: "userTest1",
	}
	tokenStr := Login(t, userReq, httpApp)

	// Create User
	userTest := models_api.User{
		Username: random.StringRandom(10),
		Password: random.StringRandom(20),
		Email:    random.StringRandom(10),
	}
	CreateUser(t, userTest, tokenStr, httpApp)

	// Create Ne
	neTest := models_api.NeData{
		Name:             random.StringRandom(10),
		Url:              random.StringRandom(10),
		Type:             random.StringRandom(10),
		MasterIpConfig:   random.Ipv4Random(),
		MasterPortConfig: int32(random.IntRandom(0, 9999)),
		SlaveIpConfig:    random.Ipv4Random(),
		SlavePortConfig:  int32(random.IntRandom(0, 9999)),
		IpCommand:        random.Ipv4Random(),
		PortCommand:      int32(random.IntRandom(0, 9999)),
		Description:      random.StringRandom(10),
		Namespace:        random.StringRandom(10),
	}

	neReqTest, err := json.Marshal(neTest)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/network-element", bytes.NewBuffer(neReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err := httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusNotFound)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User From DB
	userGetTest, err := user.GetProfileByUsername(userTest.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userGetTest)
	require.Equal(t, userTest.Username, userGetTest.Username)
	require.Equal(t, true, userGetTest.Active)
	if bcrypt.Matches(userTest.Username+userTest.Password, userGetTest.Password) {
		require.Error(t, errors.New("password in correct"))
	}
	require.NotZero(t, userGetTest.Id)
	require.NotZero(t, userGetTest.CreatedDate)

	// Get Ne From DB
	neGetTest, err := network_elements.GetNetworkElement(neTest.Name, neTest.Namespace)
	require.NoError(t, err)
	require.NotEmpty(t, neGetTest)
	require.Equal(t, neTest.Name, neGetTest.Name)
	require.Equal(t, neTest.Url, neGetTest.Url)
	require.Equal(t, neTest.Type, neGetTest.Type)
	require.Equal(t, neTest.MasterIpConfig, neGetTest.MasterIpConfig)
	require.Equal(t, neTest.MasterPortConfig, neGetTest.MasterPortConfig)
	require.Equal(t, neTest.SlaveIpConfig, neGetTest.SlaveIpConfig)
	require.Equal(t, neTest.SlavePortConfig, neGetTest.SlavePortConfig)
	require.Equal(t, neTest.IpCommand, neGetTest.IpCommand)
	require.Equal(t, neTest.PortCommand, neGetTest.PortCommand)
	require.Equal(t, neTest.Namespace, neGetTest.Namespace)

	// Add User Ne
	userNeTest := models_api.UserNe{
		UserId: userGetTest.Id,
		NeId:   neGetTest.NeId,
	}
	userNeReqTest, err := json.Marshal(userNeTest)
	require.NoError(t, err)
	req = httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/user/network-element", bytes.NewBuffer(userNeReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)
	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User Ne from DB to verify
	userNeGetTest, err := user.NeUserGet(userNeTest.UserId, userNeTest.NeId)
	require.NoError(t, err)
	require.NotEmpty(t, userNeGetTest)
	require.NotZero(t, userNeGetTest.Id)
	require.Equal(t, userNeTest.UserId, userNeGetTest.UserId)
	require.Equal(t, userNeTest.NeId, userNeGetTest.NeId)

	// Delete User Ne
	req = httptest.NewRequest(http.MethodDelete, "/mgt-svc/v1/user/network-element", bytes.NewBuffer(userNeReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)
	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User Ne from DB again to verify
	userNeGetTest, err = user.NeUserGet(userNeTest.UserId, userNeTest.NeId)
	if !errors.Is(err, models_error.ErrNotFoundUserNe) {
		require.Error(t, errors.New("delete user ne un success"))
	}

	// Delete Ne
	req = httptest.NewRequest(http.MethodDelete, "/mgt-svc/v1/network-element", bytes.NewBuffer(neReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Delete User
	DeleteUser(t, userTest.Username)
}

// Test04 testing login, create role, update role, delete role
func Test04(t *testing.T) {
	// Create server
	httpApp := server.Initialize("E:\\Data\\Go\\go-cli-mgt\\.env")

	// Login
	// WARNING: This user must have been in db before
	userReq := models_api.RequestUser{
		Username: "userTest1",
		Password: "userTest1",
	}
	tokenStr := Login(t, userReq, httpApp)

	// Create role
	roleTest := models_api.Role{
		RoleName:    random.StringRandom(10),
		Priority:    random.StringRandom(10),
		Description: random.StringRandom(30),
	}
	roleReqTest, err := json.Marshal(roleTest)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/role", bytes.NewBuffer(roleReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err := httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusNotFound)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get Role from db
	roleGetTest, err := role.GetRoleByName(roleTest.RoleName)
	require.NoError(t, err)
	require.NotEmpty(t, roleGetTest)
	require.NotEmpty(t, roleGetTest.RoleId)
	require.Equal(t, roleTest.RoleName, roleGetTest.RoleName)
	require.Equal(t, roleTest.Priority, roleGetTest.Priority)
	require.Equal(t, roleTest.Description, roleGetTest.Description)

	// Update role
	roleTest.Description = random.StringRandom(20)
	roleReqTest, err = json.Marshal(roleTest)
	require.NoError(t, err)
	req = httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/role", bytes.NewBuffer(roleReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusNotFound)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get Role from db
	roleGetTest, err = role.GetRoleByName(roleTest.RoleName)
	require.NoError(t, err)
	require.NotEmpty(t, roleGetTest)
	require.Equal(t, roleTest.RoleName, roleGetTest.RoleName)
	require.Equal(t, roleTest.Priority, roleGetTest.Priority)
	require.Equal(t, roleTest.Description, roleGetTest.Description)

	// Delete Role
	req = httptest.NewRequest(http.MethodDelete, "/mgt-svc/v1/role", bytes.NewBuffer(roleReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get Role from db again to verify
	roleGetTest, err = role.GetRoleByName(roleTest.RoleName)
	if errors.Is(err, models_error.ErrNotFoundRole) == false {
		require.Error(t, errors.New("delete role un success"))
	}
}

// Test05 testing login, create user, create role, add user role, delete user role, delete role, delete user
func Test05(t *testing.T) {
	// Create server
	httpApp := server.Initialize("E:\\Data\\Go\\go-cli-mgt\\.env")

	// Login
	// WARNING: This user must have been in db before
	userReq := models_api.RequestUser{
		Username: "userTest1",
		Password: "userTest1",
	}
	tokenStr := Login(t, userReq, httpApp)

	// Create User
	userTest := models_api.User{
		Username: random.StringRandom(10),
		Password: random.StringRandom(20),
		Email:    random.StringRandom(10),
	}
	CreateUser(t, userTest, tokenStr, httpApp)

	// Create role
	roleTest := models_api.Role{
		RoleName:    random.StringRandom(10),
		Priority:    random.StringRandom(10),
		Description: random.StringRandom(30),
	}
	roleReqTest, err := json.Marshal(roleTest)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/role", bytes.NewBuffer(roleReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err := httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusNotFound)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User From DB
	userGetTest, err := user.GetProfileByUsername(userTest.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userGetTest)
	require.Equal(t, userTest.Username, userGetTest.Username)
	require.Equal(t, true, userGetTest.Active)
	if bcrypt.Matches(userTest.Username+userTest.Password, userGetTest.Password) {
		require.Error(t, errors.New("password in correct"))
	}
	require.NotZero(t, userGetTest.Id)
	require.NotZero(t, userGetTest.CreatedDate)

	// Get Role from db
	roleGetTest, err := role.GetRoleByName(roleTest.RoleName)
	require.NoError(t, err)
	require.NotEmpty(t, roleGetTest)
	require.NotEmpty(t, roleGetTest.RoleId)
	require.Equal(t, roleTest.RoleName, roleGetTest.RoleName)
	require.Equal(t, roleTest.Priority, roleGetTest.Priority)
	require.Equal(t, roleTest.Description, roleGetTest.Description)

	// Add User Role
	userRoleTest := models_api.UserRole{
		UserId: userGetTest.Id,
		RoleId: roleGetTest.RoleId,
	}
	userRoleReqTest, err := json.Marshal(userRoleTest)
	require.NoError(t, err)
	req = httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/user/role", bytes.NewBuffer(userRoleReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)
	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User Role from db
	userRoleGetTest, err := user.RoleUserGet(userRoleTest.UserId, userRoleTest.RoleId)
	require.NoError(t, err)
	require.NotEmpty(t, userRoleGetTest)
	require.NotZero(t, userRoleGetTest.Id)
	require.Equal(t, userRoleTest.UserId, userRoleGetTest.UserId)
	require.Equal(t, userRoleTest.RoleId, userRoleGetTest.RoleId)

	// Delete User Role
	req = httptest.NewRequest(http.MethodDelete, "/mgt-svc/v1/user/role", bytes.NewBuffer(userRoleReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)
	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User Role from db to verify again
	userRoleGetTest, err = user.RoleUserGet(userRoleTest.UserId, userRoleTest.RoleId)
	if !errors.Is(err, models_error.ErrNotFoundUserRole) {
		require.Error(t, errors.New("delete user role un success"))
	}

	// Delete Role
	req = httptest.NewRequest(http.MethodDelete, "/mgt-svc/v1/role", bytes.NewBuffer(roleReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get Role from db again to verify
	roleGetTest, err = role.GetRoleByName(roleTest.RoleName)
	if errors.Is(err, models_error.ErrNotFoundRole) == false {
		require.Error(t, errors.New("delete role un success"))
	}

	// Delete User
	DeleteUser(t, userTest.Username)
}

// Test06 testing login, create user, change password user and delete user
func Test06(t *testing.T) {
	// Create server
	httpApp := server.Initialize("E:\\Data\\Go\\go-cli-mgt\\.env")

	// Login
	// WARNING: This user must have been in db before
	userReq := models_api.RequestUser{
		Username: "userTest1",
		Password: "userTest1",
	}
	tokenStr := Login(t, userReq, httpApp)

	// Create User
	userTest := models_api.User{
		Username: random.StringRandom(10),
		Password: random.StringRandom(20),
		Email:    random.StringRandom(10),
	}
	CreateUser(t, userTest, tokenStr, httpApp)

	// Get User From DB
	userGetTest, err := user.GetProfileByUsername(userTest.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userGetTest)
	require.Equal(t, userTest.Username, userGetTest.Username)
	require.Equal(t, true, userGetTest.Active)
	if bcrypt.Matches(userTest.Username+userTest.Password, userGetTest.Password) {
		require.Error(t, errors.New("password in correct"))
	}
	require.NotZero(t, userGetTest.Id)
	require.NotZero(t, userGetTest.CreatedDate)

	// Change password user
	newPassword := random.StringRandom(20)
	userChangePasswordTest := models_api.ChangePassWord{
		Username:    userTest.Username,
		OldPassword: userTest.Password,
		NewPassword: newPassword,
	}
	userChangePasswordReqTest, err := json.Marshal(userChangePasswordTest)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/auth/change-password", bytes.NewBuffer(userChangePasswordReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err := httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusNotFound)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User From Db again
	userGetTest, err = user.GetProfileByUsername(userTest.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userGetTest)
	require.Equal(t, userTest.Username, userGetTest.Username)
	if bcrypt.Matches(userTest.Username+newPassword, userGetTest.Password) {
		require.Error(t, errors.New("password incorrect"))
	}

	// Delete User
	DeleteUser(t, userTest.Username)
}
