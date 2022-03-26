package handlers
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Home Page",
	}) 
}

func ShowInbox(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Home Page",
	}) 
}

func ShowOutbox(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Home Page",
	}) 
}

func PerformSend(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Home Page",
	}) 
}

func PerformBlock(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Home Page",
	}) 
}