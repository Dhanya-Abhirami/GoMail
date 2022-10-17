package main

import (
	"server/configs"
	"server/messaging"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	configs.ConnectDB()
	router = gin.Default()
	initializeRoutes()
	go messaging.ConsumeEmail()
	router.Run()
}