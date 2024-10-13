package unit

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_error"
	"go-cli-mgt/pkg/models/models_response"
	"go-cli-mgt/pkg/server"
	"go-cli-mgt/pkg/service/network_elements"
	"go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/bcrypt"
	"go-cli-mgt/pkg/service/utils/random"
	"io"
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
	userReqBytes, err := json.Marshal(userReq)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/auth/login", bytes.NewBuffer(userReqBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpApp.Test(req)
	require.NoError(t, err)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusUnauthorized)

	responseBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	var loginResp models_response.RespSuccess
	err = json.Unmarshal(responseBody, &loginResp)
	require.NoError(t, err)
	require.Equal(t, loginResp.Status, true)
	require.Equal(t, loginResp.Code, http.StatusOK)
	require.NotEmpty(t, loginResp.Message)
	tokenStr := "Basic " + loginResp.Message

	// Create user
	userTest := models_api.User{
		Username: random.StringRandom(10),
		Password: random.StringRandom(20),
		Email:    random.StringRandom(10),
	}
	userReqTest, err := json.Marshal(userTest)
	require.NoError(t, err)
	req = httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/user/profile", bytes.NewBuffer(userReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User from db
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

	// Deactivate user
	req = httptest.NewRequest(http.MethodDelete, "/mgt-svc/v1/user/profile", bytes.NewBuffer(userReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)
	resp, err = httpApp.Test(req)
	require.NoError(t, err)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User from db to check
	userGetTest, err = user.GetProfileByUsername(userTest.Username)
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
	err = user.DeleteProfile(userTest.Username)
	require.NoError(t, err)
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
	userReqBytes, err := json.Marshal(userReq)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/auth/login", bytes.NewBuffer(userReqBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpApp.Test(req)
	require.NoError(t, err)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusUnauthorized)

	responseBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	var loginResp models_response.RespSuccess
	err = json.Unmarshal(responseBody, &loginResp)
	require.NoError(t, err)
	require.Equal(t, loginResp.Status, true)
	require.Equal(t, loginResp.Code, http.StatusOK)
	require.NotEmpty(t, loginResp.Message)
	tokenStr := "Basic " + loginResp.Message

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
	req = httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/network-element", bytes.NewBuffer(neReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err = httpApp.Test(req)
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
	userReqBytes, err := json.Marshal(userReq)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/auth/login", bytes.NewBuffer(userReqBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpApp.Test(req)
	require.NoError(t, err)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusUnauthorized)

	responseBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	var loginResp models_response.RespSuccess
	err = json.Unmarshal(responseBody, &loginResp)
	require.NoError(t, err)
	require.Equal(t, loginResp.Status, true)
	require.Equal(t, loginResp.Code, http.StatusOK)
	require.NotEmpty(t, loginResp.Message)
	tokenStr := "Basic " + loginResp.Message

	// Create User
	userTest := models_api.User{
		Username: random.StringRandom(10),
		Password: random.StringRandom(20),
		Email:    random.StringRandom(10),
	}
	userReqTest, err := json.Marshal(userTest)
	require.NoError(t, err)
	req = httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/user/profile", bytes.NewBuffer(userReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err = httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

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
	req = httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/network-element", bytes.NewBuffer(neReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err = httpApp.Test(req)
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
	err = user.DeleteProfile(userTest.Username)
	require.NoError(t, err)
}
