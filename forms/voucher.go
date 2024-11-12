package forms

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
)

//VoucherForm ...
type VoucherForm struct{}

//CreateCampaignForm ...
type CreateVoucherForm struct {
	CampaignID   string `form:"campaign_id" json:"campaign_id" binding:"required"`
	UserID   string `form:"user_id" json:"user_id" binding:"required"`
	Code   string `form:"code" json:"code" binding:"required"`
	DiscountRate float64 `form:"discount_rate" json:"discount_rate" binding:"required,gte=1,lte=100"`
	ExpirationDate time.Time `form:"expiration_date" json:"expiration_date" binding:"required" time_format:"2003-01-02 15:04:05"`
	Status string `form:"status" json:"status"`
}

//CampaignID ...
func (f VoucherForm) CampaignID(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the campaign id"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

//UserID ...
func (f VoucherForm) UserID(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the user id"
		}
		return errMsg[0]
	default:
		return "Something went wrong, please try again later"
	}
}

//DiscountRate ...
func (f VoucherForm) DiscountRate(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the voucher discount rate"
		}
		return errMsg[0]
	case "gte", "lte":
		return "Discount rate should be between 1 and 100"
	default:
		return "Something went wrong, please try again later"
	}
}

//Code ...
func (f VoucherForm) Code(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the voucher code"
		}
		return errMsg[0]
	default:
		return "Voucher limit: Something went wrong, please try again later"
	}
}

func (f VoucherForm) ExpirationDate(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the campaign voucher limit"
		}
		return errMsg[0]
	default:
		return "Expiration Date: Something went wrong, please try again later"
	}
}

//Create ...
func (f VoucherForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "CampaignID" {
				return f.CampaignID(err.Tag())
			}

			if err.Field() == "UserID" {
				return f.UserID(err.Tag())
			}

			if err.Field() == "DiscountRate" {
				return f.DiscountRate(err.Tag())
			}

            if err.Field() == "ExpirationDate" {
				return f.ExpirationDate(err.Tag())
			}

            if err.Field() == "Code" {
				return f.Code(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Create voucher: Something went wrong, please try again later"
}