package handlers
import (
	"context"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"server/configs"
	"server/models"
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
	filter := bson.D{{"receiver" , "dhanya"}}
	opts := options.Find()
	opts.SetSort(bson.D{{"sentAt", -1}})
	var inbox []models.Email
	cursor, err := emailCollection.Find(ctx,filter,opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Unable to process emails",
			"error": err.Error(),
		})
	}
	if err = cursor.All(ctx, &inbox); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Unable to process emails",
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   inbox,
	}) 
}

func ShowOutbox(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{"sender" , "dhanya"}}
	opts := options.Find()
	opts.SetSort(bson.D{{"sentAt", -1}})
	var outbox []models.Email
	cursor, err := emailCollection.Find(ctx,filter,opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Unable to process emails",
			"error": err.Error(),
		})
	}
	if err = cursor.All(ctx, &outbox); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Unable to process emails",
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   outbox,
	}) 
}

func PerformSend(c *gin.Context) {
	var input models.Email
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	newEmail := models.Email{
		Sender:     input.Sender,
		Receiver: input.Receiver,
		Subject: input.Subject,
		Body: input.Body,
	}
	result, err := emailCollection.InsertOne(ctx, newEmail)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Failed Email Send",
			"error": err.Error(),
		})
	} 
	c.JSON(http.StatusOK, gin.H{
		"message":   "Email Sent",
		"Id": result.InsertedID,
	})
}

