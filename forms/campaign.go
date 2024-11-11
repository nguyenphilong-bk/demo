package forms

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
)

//CampaignForm ...
type CampaignForm struct{}

//CreateCampaignForm ...
type CreateCampaignForm struct {
	Name   string `form:"name" json:"name" binding:"required,min=3,max=100"`
	DiscountRate float64 `form:"discount_rate" json:"discount_rate" binding:"required,gte=1,lte=100"`
	VoucherLimit int `form:"voucher_limit" json:"voucher_limit" binding:"required,gte=1"`
	StartDate time.Time `form:"start_date" json:"start_date" binding:"required,ltefield=EndDate" time_format:"2003-01-02 15:04:05"`
	EndDate time.Time `form:"end_date" json:"end_date" binding:"required" time_format:"2003-01-02 15:04:05"`
}

//Name ...
func (f CampaignForm) Name(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the campaign name"
		}
		return errMsg[0]
	case "min", "max":
		return "Name should be between 3 to 100 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

//DiscountRate ...
func (f CampaignForm) DiscountRate(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the campaign discount rate"
		}
		return errMsg[0]
	case "gte", "lte":
		return "Discount rate should be between 1 and 100"
	default:
		return "Something went wrong, please try again later"
	}
}

//DiscountRate ...
func (f CampaignForm) VoucherLimit(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the campaign voucher limit"
		}
		return errMsg[0]
	default:
		return "Voucher limit: Something went wrong, please try again later"
	}
}

func (f CampaignForm) StartDate(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the campaign voucher limit"
		}
		return errMsg[0]
    case "ltefield":
        return "start_date must be less than or equal to end_date"
	default:
		return "Voucher limit: Something went wrong, please try again later"
	}
}

func (f CampaignForm) EndDate(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the campaign voucher limit"
		}
		return errMsg[0]
	default:
		return "Voucher limit: Something went wrong, please try again later"
	}
}

//Create ...
func (f CampaignForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return f.Name(err.Tag())
			}

			if err.Field() == "DiscountRate" {
				return f.DiscountRate(err.Tag())
			}

            if err.Field() == "VoucherLimit" {
				return f.VoucherLimit(err.Tag())
			}

            if err.Field() == "StartDate" {
				return f.StartDate(err.Tag())
			}

            if err.Field() == "EndDate" {
				return f.EndDate(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Create Campaign: Something went wrong, please try again later"
}

//Update ...
func (f CampaignForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return f.Name(err.Tag())
			}

			if err.Field() == "DiscountRate" {
				return f.DiscountRate(err.Tag())
			}

            if err.Field() == "VoucherLimit" {
				return f.DiscountRate(err.Tag())
			}

            if err.Field() == "StartDate" {
				return f.DiscountRate(err.Tag())
			}

            if err.Field() == "EndDate" {
				return f.DiscountRate(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
