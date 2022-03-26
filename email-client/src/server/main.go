package main

import (
	"server/configs"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	configs.ConnectDB()
	router = gin.Default()
	initializeRoutes()
	router.Run()
}