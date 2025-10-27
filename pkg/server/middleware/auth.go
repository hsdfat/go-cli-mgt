package middleware

import (
	"errors"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/utils/response"
	"go-cli-mgt/pkg/utils/token"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	UsernameContextKey = "username"
	RoleContextKey     = "role"
)

var (
	MissingAuthHeader = errors.New("missing authorize header")
	InvalidAuthHeader = errors.New("authorize header of request invalid")
	InvalidToken      = errors.New("invalid token")
)

// BasicAuth for Authenticate middleware
func BasicAuth(c *fiber.Ctx) error {
	// Basic auth implementation

	authorizeHeaders := c.GetReqHeaders()["Authorization"]
	if authorizeHeaders == nil {
		logger.Logger.Error("Request don't have authorize header")
		response.Unauthorized(c)
		return MissingAuthHeader
	}
	authHeader := c.GetReqHeaders()["Authorization"][0]
	if authHeader == "" {
		logger.Logger.Error("Request don't have authorize header")
		response.Unauthorized(c)
		return MissingAuthHeader
	}

	authHeaderParts := strings.SplitN(authHeader, " ", 2)
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Basic" {
		logger.Logger.Error("Authorize header of request invalid")
		response.Unauthorized(c)
		return InvalidAuthHeader
	}

	tokenStr := authHeaderParts[1]
	username, roles, err := token.ParseToken(tokenStr)
	if err != nil {
		if errors.Is(err, InvalidToken) {
			logger.Logger.Error("Token invalid")
			response.Unauthorized(c)
			return InvalidToken
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
