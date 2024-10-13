package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/models/models_error"
	"go-cli-mgt/pkg/service/utils/response"
	"go-cli-mgt/pkg/service/utils/token"
	"strings"
)

const (
	UsernameContextKey = "username"
	RoleContextKey     = "role"
)

// BasicAuth for Authenticate middleware
func BasicAuth(c *fiber.Ctx) error {
	// Basic auth implementation

	authorizeHeaders := c.GetReqHeaders()["Authorization"]
	if authorizeHeaders == nil {
		logger.Logger.Error("Request don't have authorize header")
		response.Unauthorized(c)
		return models_error.MissingAuthHeader
	}
	authHeader := c.GetReqHeaders()["Authorization"][0]
	if authHeader == "" {
		logger.Logger.Error("Request don't have authorize header")
		response.Unauthorized(c)
		return models_error.MissingAuthHeader
	}

	authHeaderParts := strings.SplitN(authHeader, " ", 2)
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Basic" {
		logger.Logger.Error("Authorize header of request invalid")
		response.Unauthorized(c)
		return models_error.InvalidAuthHeader
	}

	tokenStr := authHeaderParts[1]
	username, roles, err := token.ParseToken(tokenStr)
	if err != nil {
		if errors.Is(err, models_error.InvalidToken) {
			logger.Logger.Error("Token invalid")
			response.Unauthorized(c)
			return models_error.InvalidToken
		}

		logger.Logger.Error("Cannot verify token: ", err)
		response.InternalError(c, "Cannot verify token")
		return err
	}

	c.Set(UsernameContextKey, username)
	c.Set(RoleContextKey, roles)

	logger.Logger.Info("User retrieved from token: ", username)

	// Check if the user has an "admin" role
	checkRole := false
	roleList := strings.Split(roles, " ")
	for _, role := range roleList {
		if strings.EqualFold(role, "admin") {
			checkRole = true
			break
		}
	}

	// If the user does not have the required role, return a forbidden status
	if checkRole == false {
		logger.Logger.Error("User does not have the required role")
		response.Forbidden(c, "no authority")
		return nil
	}

	return c.Next()
}
