package controllers

import (
	"encoding/json"

	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/services"
	"github.com/Massad/gin-boilerplate/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct{}

var userService = new(services.UserService)
var userForm = new(forms.UserForm)

// getUserID ...
func getUserID(c *gin.Context) (userID string) {
	//MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.MustGet("userID").(string)
}

// Login ...
func (ctrl UserController) Login(c *gin.Context) {
	var loginForm forms.LoginForm

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Login(validationErr)
		// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Message: message})
		return
	}

	user, token, err := userService.Login(loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Message: "Invalid login details"})
		return
	}

	temp, _ := json.Marshal(&user)
	var result map[string]interface{}
	json.Unmarshal(temp, &result)
	result["token"] = token
	// c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
	c.JSON(http.StatusOK, utils.Response{StatusCode: http.StatusOK, Message: "Successfully logged in", Data: result})

}

// Register ...
func (ctrl UserController) Register(c *gin.Context) {
	var registerForm forms.RegisterForm

	if validationErr := c.ShouldBindJSON(&registerForm); validationErr != nil {
		message := userForm.Register(validationErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Message: message})
		return
	}

	user, err := userService.Register(registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	temp, _ := json.Marshal(&user)
	var result map[string]interface{}
	json.Unmarshal(temp, &result)

	c.JSON(http.StatusOK, utils.Response{StatusCode: http.StatusOK, Message: "Register new account successfully", Data: result})
}

// Logout ...
func (ctrl UserController) Logout(c *gin.Context) {
	_, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Message: "User not logged in"})
		return
	}
	// for redis
	// deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
	// if delErr != nil || deleted == 0 { //if any goes wrong
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
	// 	return
	// }
	c.JSON(http.StatusOK, utils.Response{StatusCode: http.StatusOK, Message: "Successfully logged out"})
}

func (ctrl UserController) RegisterCampaign(c *gin.Context) {
	var registerCampaignForm forms.RegisterCampaignForm

	if validationErr := c.ShouldBindJSON(&registerCampaignForm); validationErr != nil {
		message := userForm.Register(validationErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Message: message})
		return
	}

	userVoucher, err := userService.RegisterCampaign(registerCampaignForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}


	temp, _ := json.Marshal(&userVoucher)
	var result map[string]interface{}
	json.Unmarshal(temp, &result)

	c.JSON(http.StatusOK, utils.Response{StatusCode: http.StatusOK, Message: "Register new account successfully", Data: result})
}
