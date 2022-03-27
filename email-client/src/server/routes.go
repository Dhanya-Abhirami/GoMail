package main

import (
	"server/handlers"
	"server/middleware"
)

func initializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)
	mailRoutes := router.Group("/mail")
	{
		mailRoutes.GET("/inbox", middleware.EnsureLoggedIn(), handlers.ShowInbox)
		mailRoutes.GET("/outbox", middleware.EnsureLoggedIn(), handlers.ShowOutbox)
		mailRoutes.POST("/send", middleware.EnsureLoggedIn(), handlers.PerformSend)
	}

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn(), handlers.ShowRegisterPage)
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn(), handlers.RegisterNewUser)
		mailRoutes.POST("/block", middleware.EnsureLoggedIn(), handlers.PerformBlock)
		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), handlers.PerformLogout)
	}
}
