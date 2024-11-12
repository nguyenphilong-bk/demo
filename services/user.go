package services

import (
	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/models"
	"github.com/Massad/gin-boilerplate/utils"
)

// UserService ...
type UserService struct{}

var userModel = new(models.UserModel)
var campaignService = new(CampaignService)
var voucherService = new(VoucherService)

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

func (service UserService) RegisterCampaign(form forms.RegisterCampaignForm) (userVoucher models.UserVoucher, err error) {
	var registerForm forms.RegisterForm
	registerForm.Email = form.Email
	registerForm.Name = form.Name
	registerForm.Password = form.Password

	// Get campaign
	campaign, err := campaignService.One(form.CampaignID)
	if err != nil {
		return userVoucher, err
	}

	// Create new user
	user, err := service.Register(registerForm)
	if err != nil {
		return userVoucher, err
	}

	userVoucher.User = user

	numberOfVouchers, err := voucherService.CountByCampaign(campaign.ID.String())
	if err != nil {
		return userVoucher, err
	}

	if numberOfVouchers < campaign.VoucherLimit {
		code, err := utils.RandomStringCrypto(utils.CODE_LENGTH)
		if err != nil {
			return userVoucher, err
		}
		voucherForm := forms.CreateVoucherForm{
			CampaignID:     campaign.ID.String(),
			UserID:         user.ID.String(),
			Code:           code,
			DiscountRate:   campaign.DiscountRate,
			ExpirationDate: campaign.EndDate,
			Status:         utils.STATUS_ISSUED,
		}
		voucher, err := voucherService.Create(voucherForm)
		if err != nil {
			return userVoucher, err
		}

		userVoucher.Code = voucher.Code
		userVoucher.VoucherID = voucher.ID.String()
		return userVoucher, nil
	}

	return userVoucher, nil
}
