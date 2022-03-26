package handlers
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"server/models"
	"server/utils"
)

type UserInput struct {
	Username  string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if models.IsUserValid(input.Username, input.Password) {
		setLoginToken(c)
		c.JSON(http.StatusOK, gin.H{
			"message":   "Successful Login",
		}) 
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Login Failed",
			"error": "Invalid Credentials",
		}) 
	}
}

func RegisterNewUser(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if _, err := models.RegisterNewUser(input.Username, input.Password); err == nil {
		setLoginToken(c)
		c.JSON(http.StatusOK, gin.H{
			"message":   "Registration Successful",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Failed Registration",
			"error": err.Error(),
		})
	}
	
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

