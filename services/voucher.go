package services

import (
	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/models"
)

//CampaignController ...
type VoucherService struct{}

var voucherModel = new(models.VoucherModel)

// All ... 
// Find all the voucher belongs to this user
func (service VoucherService) All(userID string) ([]models.Voucher, error) {
	results, err := voucherModel.All(userID)
	if err != nil {
		return results, err
	}

	return results, nil
}

// Count the number of vouchers by campaign id
func (service VoucherService) CountByCampaign(campaignID string) (int, error) {
	counter, err := voucherModel.CountByCampaign(campaignID)
	if err != nil {
		return counter, err
	}

	return counter, nil
}

// Create a new voucher
func (service VoucherService) Create(form forms.CreateVoucherForm) (voucher models.Voucher, err error) {
	voucher, err = voucherModel.Create(form)
	if err != nil {
		return voucher, err
	} 

	return voucher, nil
}
