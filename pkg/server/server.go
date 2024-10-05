package server

import (
	"github.com/gofiber/fiber/v2"
	"go-cli-mgt/pkg/config"
	"go-cli-mgt/pkg/handler/auth"
	"go-cli-mgt/pkg/handler/history"
	"go-cli-mgt/pkg/handler/network_element"
	"go-cli-mgt/pkg/handler/permission"
	"go-cli-mgt/pkg/handler/user"
	"go-cli-mgt/pkg/logger"
	"go-cli-mgt/pkg/middleware"
)

// New echo router

func NewFiber() *fiber.App {
	app := fiber.New()
	router := app.Group("/mgt-svc/v1")

	authRouter := router.Group("/auth")
	{
		authRouter.Post("/login", auth.LoginHandler)

		r := authRouter.Post("/change-password", auth.ChangePasswordHandler)
		r.Use(middleware.BasicAuth)
	}

	userRouter := router.Group("/user")
	{
		userRouter.Use(middleware.BasicAuth)
		// profile
		userRouter.Post("/profile", user.ProfileCreateHandler)
		userRouter.Delete("/profile", user.ProfileDeactivateHandler)
		// change password
		userRouter.Post("change-password", user.ChangePasswordHandler)
		// user's permissions
		userRouter.Post("/permission", user.PermissionAddHandler)
		userRouter.Delete("/permission", user.PermissionDeleteHandler)
		userRouter.Get("/permission", user.PermissionGetHandler)
		// user's network elements
		userRouter.Post("/network-element", user.NetworkElementAddHandler)
		userRouter.Delete("/network-element", user.NetworkElementDeleteHandler)
		userRouter.Get("/network-elements", user.NetworkElementsListHandler)
		userRouter.Post("/network-elements/delete", user.NetworkElementsListDeleteHandler)
	}
	userListRouter := router.Group("/users")
	{
		userListRouter.Use(middleware.BasicAuth)

		userListRouter.Get("/permission", user.ListUsersPermissionHandler)
		userListRouter.Get("/network-element", user.ListUsersNetworkElementHandler)
		userListRouter.Get("/profile", user.ListUsersProfileHandler)
	}
	permissionRouter := router.Group("/permission")
	{
		permissionRouter.Use(middleware.BasicAuth)
		router.Get("/permission", permission.ListPermissionHandler)
		router.Post("/permission", permission.CreateOrUpdateHandler)
		router.Delete("/permission", permission.DeleteHandler)
	}
	networkElementRouter := router.Group("/network-element")
	{
		networkElementRouter.Use(middleware.BasicAuth)
		router.Get("/network-element", network_element.ListNetworkElementHandler)
		router.Post("/network-element", network_element.CreateOrUpdateHandler)
		router.Delete("/network-element", network_element.DeleteHandler)
	}
	historyRouter := router.Group("/history")
	{
		historyRouter.Use(middleware.BasicAuth)
		router.Post("/hitory", history.SaveHistoryHandler)
		router.Get("/hitory", history.GetHistoryHandler)
	}
	return app
}

func ListenAndServe(app *fiber.App) {
	cfg := config.GetServerConfig()

	err := app.Listen(cfg.Host + ":" + cfg.Port)
	if err != nil {
		logger.Logger.Fatalf("Can't listen server: %v", err)
	}
}
