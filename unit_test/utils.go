package unit

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"go-cli-mgt/pkg/models/models_api"
	"go-cli-mgt/pkg/models/models_response"
	"go-cli-mgt/pkg/service/user"
	"go-cli-mgt/pkg/service/utils/bcrypt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Login(t *testing.T, userReq models_api.RequestUser, app *fiber.App) string {
	userReqBytes, err := json.Marshal(userReq)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/auth/login", bytes.NewBuffer(userReqBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
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
	return tokenStr
}

func CreateUser(t *testing.T, userCreate models_api.User, tokenStr string, httpApp *fiber.App) {
	userReqTest, err := json.Marshal(userCreate)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/mgt-svc/v1/user/profile", bytes.NewBuffer(userReqTest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenStr)

	resp, err := httpApp.Test(req)
	require.NotEqual(t, resp.StatusCode, http.StatusForbidden)
	require.NotEqual(t, resp.StatusCode, http.StatusInternalServerError)
	require.NotEqual(t, resp.StatusCode, http.StatusBadRequest)

	// Get User from db
	userGetTest, err := user.GetProfileByUsername(userCreate.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userGetTest)
	require.Equal(t, userCreate.Username, userGetTest.Username)
	require.Equal(t, true, userGetTest.Active)
	if bcrypt.Matches(userCreate.Username+userCreate.Password, userGetTest.Password) {
		require.Error(t, errors.New("password in correct"))
	}
	require.NotZero(t, userGetTest.Id)
	require.NotZero(t, userGetTest.CreatedDate)
}

func DeleteUser(t *testing.T, username string) {
	err := user.DeleteProfile(username)
	require.NoError(t, err)
}
