package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// UserForm ...
type UserForm struct{}

// LoginForm ...
type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=50"`
}

// RegisterForm ...
type RegisterForm struct {
	Name     string `form:"name" json:"name" binding:"required,min=3,max=20,fullName"` //fullName rule is in validator.go
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=50"`
}

// RegisterCampaignForm ...
type RegisterCampaignForm struct {
	Name       string `form:"name" json:"name" binding:"required,min=3,max=20,fullName"` //fullName rule is in validator.go
	Email      string `form:"email" json:"email" binding:"required,email"`
	Password   string `form:"password" json:"password" binding:"required,min=3,max=50"`
	CampaignID string `form:"campaign_id" json:"campaign_id" binding:"required"`
}

// Name ...
func (f UserForm) Name(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your name"
		}
		return errMsg[0]
	case "min", "max":
		return "Your name should be between 3 to 20 characters"
	case "fullName":
		return "Name should not include any special characters or numbers"
	default:
		return "Something went wrong, please try again later"
	}
}

// Email ...
func (f UserForm) Email(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your email"
		}
		return errMsg[0]
	case "min", "max", "email":
		return "Please enter a valid email"
	default:
		return "Something went wrong, please try again later"
	}
}

// Password ...
func (f UserForm) Password(tag string) (message string) {
	switch tag {
	case "required":
		return "Please enter your password"
	case "min", "max":
		return "Your password should be between 3 and 50 characters"
	case "eqfield":
		return "Your passwords does not match"
	default:
		return "Something went wrong, please try again later"
	}
}

// CampaignID ...
func (f UserForm) CampaignID(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the campaign id"
		}
		return errMsg[0]
	default:
		return "CampaignID: Something went wrong, please try again later"
	}
}

// Signin ...
func (f UserForm) Login(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}
			if err.Field() == "Password" {
				return f.Password(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// Register ...
func (f UserForm) Register(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return f.Name(err.Tag())
			}

			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}

			if err.Field() == "Password" {
				return f.Password(err.Tag())
			}

		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// RegisterCampaign ...
func (f UserForm) RegisterCampaign(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return f.Name(err.Tag())
			}

			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}

			if err.Field() == "Password" {
				return f.Password(err.Tag())
			}

			if err.Field() == "CampaignID" {
				return f.CampaignID(err.Tag())
			}

		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
