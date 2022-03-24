package main

import (
	"src/handlers"
	"src/middleware"
)

func initializeRoutes() {
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/inbox", middleware.EnsureLoggedIn(), handlers.ShowInbox)
		userRoutes.GET("/outbox", middleware.EnsureLoggedIn(), handlers.ShowOutbox)
		userRoutes.POST("/send", middleware.EnsureLoggedIn(), handlers.PerformSend)
		userRoutes.POST("/block", middleware.EnsureLoggedIn(), handlers.PerformBlock)
	}

	mailRoutes := router.Group("/mail")
	{
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn(), handlers.ShowRegisterPage)
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn(), handlers.RegisterNewUser)
		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), handlers.PerformLogout)
	}
}
