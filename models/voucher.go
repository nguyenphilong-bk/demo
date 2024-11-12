package models

import (
	"time"

	"github.com/Massad/gin-boilerplate/db"
	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/utils"
	"github.com/google/uuid"
)

// Voucher ...
type Voucher struct {
	ID             uuid.UUID `db:"id, primarykey" json:"id"`
	CampaignID     uuid.UUID `db:"campaign_id" json:"campaign_id"`
	UserID         uuid.UUID `db:"user_id" json:"user_id"`
	Code           string    `db:"code" json:"code"`
	DiscountRate   float64   `db:"discount_rate" json:"discount_rate"`
	Status         string    `db:"status" json:"status"`
	ExpirationDate time.Time `db:"expiration_date" json:"expiration_date"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	CreatedBy      uuid.UUID `db:"created_by" json:"created_by"`
	DeletedBy      uuid.UUID `db:"deleted_by" json:"deleted_by"`
}

// VoucherModel ...
type VoucherModel struct{}

// Create ...
func (m VoucherModel) Create(form forms.CreateVoucherForm) (voucher Voucher, err error) {
	err = db.GetDB().QueryRow("INSERT INTO vouchers(campaign_id, user_id, code, discount_rate, expiration_date, status) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", form.CampaignID, form.UserID, form.Code, form.DiscountRate, form.ExpirationDate, form.Status).Scan(&voucher.ID)
	if err != nil {
		return voucher, utils.NewServerError(utils.INTERNAL_SERVER_ERROR, err.Error(), "Internal server error")
	}

	voucher.Code = form.Code
	return voucher, err
}

// All ...
func (m VoucherModel) All(userID string) (vouchers []Voucher, err error) {
	_, err = db.GetDB().Select(&vouchers, "SELECT id, campaign_id, user_id,discount_rate, status, expiration_date, status FROM vouchers WHERE user_id=$1", userID)
	return vouchers, err
}

// Count voucher by campaign_id ...
func (m VoucherModel) CountByCampaign(campaignID string) (result int, err error) {
	err = db.GetDB().SelectOne(&result, "SELECT count(id) FROM vouchers WHERE campaign_id=$1", campaignID)
	return result, err
}

