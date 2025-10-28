package server

import (
	"github.com/hsdfat/go-cli-mgt/pkg/server/middleware"

	"github.com/gofiber/fiber/v2"
)

// NewFiber New echo router
func NewFiber() *fiber.App {
	app := fiber.New()
	router := app.Group("/mgt-svc/v1")

	authHandler := NewAuthHandlerHandler()
	neHandler := NewNetworkElementHandlerHandler()
	historyHandler := NewHistoryHandler()
	roleHandler := NewRoleHandler()
	userHandler := NewUserHandler()

	authRouter := router.Group("/auth")
	{
		authRouter.Post("/login", authHandler.LoginHandler)

		r := authRouter.Post("/change-password", authHandler.ChangePasswordHandler)
		r.Use(middleware.BasicAuth)
	}

	userRouter := router.Group("/user")
	{
		userRouter.Use(middleware.BasicAuth)
		// profile
		userRouter.Post("/profile", userHandler.ProfileCreateHandler)
		userRouter.Delete("/profile", userHandler.ProfileDeactivateHandler)
		// change password
		// userRouter.Post("change-password", userHandler.ChangePasswordHandler)
		// user's permissions
		userRouter.Post("/role", userHandler.RoleAddHandler)
		userRouter.Delete("/role", userHandler.RoleDeleteHandler)
		userRouter.Get("/role", userHandler.PermissionGetHandler)
		// user's network elements
		userRouter.Post("/network-element", userHandler.NetworkElementAddHandler)
		userRouter.Delete("/network-element", userHandler.NetworkElementDeleteHandler)
		userRouter.Get("/network-elements", userHandler.NetworkElementsListHandler)
		userRouter.Post("/network-elements/delete", userHandler.NetworkElementsListDeleteHandler)
	}
	userListRouter := router.Group("/users")
	{
		userListRouter.Use(middleware.BasicAuth)

		userListRouter.Get("/role", userHandler.ListUsersPermissionHandler)
		userListRouter.Get("/network-element", userHandler.ListUsersNetworkElementHandler)
		userListRouter.Get("/profile", userHandler.ListUsersProfileHandler)
	}
	permissionRouter := router.Group("/role")
	{
		permissionRouter.Use(middleware.BasicAuth)
		router.Get("/role", roleHandler.ListRoleHandler)
		router.Post("/role", roleHandler.CreateOrUpdateHandler)
		router.Delete("/role", roleHandler.DeleteHandler)
	}
	networkElementRouter := router.Group("/network-element")
	{
		networkElementRouter.Use(middleware.BasicAuth)
		router.Get("/network-element", neHandler.ListNetworkElementHandler)
		router.Post("/network-element", neHandler.CreateOrUpdateHandler)
		router.Delete("/network-element", neHandler.DeleteHandler)
	}
	historyRouter := router.Group("/history")
	{
		historyRouter.Use(middleware.BasicAuth)
		router.Post("/history", historyHandler.SaveHistoryHandler)
		router.Get("/history", historyHandler.GetHistoryHandler)
	}
	return app
}
