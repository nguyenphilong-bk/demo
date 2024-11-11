package services

import (
	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/models"
)

// UserService ...
type UserService struct{}

var userModel = new(models.UserModel)

// Login ...
func (service UserService) Login(form forms.LoginForm) (models.User, models.Token, error) {
	user, token, err := userModel.Login(form)
	if err != nil {
		return user, token, err
	}

	return user, token, nil
}

// Register ...
func (service UserService) Register(form forms.RegisterForm) (models.User, error) {
	user, err := userModel.Register(form)
	if err != nil {
		return user, err
	}
	
	return user, nil
}

// Logout ...
// func (service UserService) Logout(c *gin.Context) {
// 	_, err := authModel.ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Message: "User not logged in"})
// 		return
// 	}
// 	// for redis
// 	// deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
// 	// if delErr != nil || deleted == 0 { //if any goes wrong
// 	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
// 	// 	return
// 	// }
// 	c.JSON(http.StatusOK, utils.Response{StatusCode: http.StatusOK, Message: "Successfully logged out"})
// }

func (service UserService) RegisterCampaign(form forms.RegisterCampaignForm) {
	var registerForm forms.RegisterForm
	registerForm.Email = form.Email
	registerForm.Name = form.Name
	registerForm.Password = form.Password

	// Create new user
	// user, err := UserService.Register(registerForm)
	// if err != nil {
	// 	return
	// }

	// Create voucher for this user
	

	// temp, _ := json.Marshal(&user)
	// var result map[string]interface{}
	// json.Unmarshal(temp, &result)

}
