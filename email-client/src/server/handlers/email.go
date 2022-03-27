package handlers
import (
	"context"
	"log"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"server/configs"
)

var emailCollection *mongo.Collection = configs.GetCollection(configs.DB, "email")

func ShowIndexPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Home Page",
	}) 
}

func ShowInbox(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	cursor, err := emailCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   cursor,
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

