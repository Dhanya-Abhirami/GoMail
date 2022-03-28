package handlers
import (
	"context"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/models"
	"server/utils"
	"server/configs"
)
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func ShowLoginPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Please login",
	}) 
}

func ShowRegisterPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Please Register",
	}) 
}

func PerformLogin(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	filter := bson.M{"name" : input.Name}

	var user models.User
	opts := options.FindOne()
	opts.SetProjection(bson.M{"password": 1})
	err := userCollection.FindOne(ctx, filter,opts).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":   "Login Failed",
			"error": "No User Exists",
		}) 
	} else{
		if user.Password == input.Password {
			setLoginToken(c)
			c.JSON(http.StatusOK, gin.H{
				"message":   "Successful Login",
			}) 
		} else{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":   "Login Failed",
				"error": "Passwords Don't Match",
			}) 
		}
	}
}

func RegisterNewUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	filter := bson.M{"name" : input.Name}
	var user models.User
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		newUser := models.User{
			Name:     input.Name,
			Password: input.Password,
		}
		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message":   "Failed Registration",
				"error": err.Error(),
			})
		} else{
			c.JSON(http.StatusOK, gin.H{
				"message":   "Registration Successful",
				"Id": result.InsertedID,
			})
		}
	} else{
		c.JSON(http.StatusOK, gin.H{
			"message":   "User Exists",
		})
	}
	
}

func PerformBlock(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Home Page",
	}) 
}

func PerformLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Set("is_logged_in", false) 
	c.JSON(http.StatusOK, gin.H{
		"message":   "Logout Successful",
	})
}

func setLoginToken(c *gin.Context) {
	token := utils.GenerateSessionToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)
}

