package unit

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_response"
	"go-cli-mgt/pkg/server"
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
