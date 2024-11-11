package models

import (
	"time"

	"github.com/Massad/gin-boilerplate/db"
	"github.com/google/uuid"
)

// Campaign ...
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

// CampaignModel ...
type VoucherModel struct{}

// Create ...
// func (m VoucherModel) Create(userID string, form forms.CreateCampaignForm) (campaignID uuid.UUID, err error) {
// 	err = db.GetDB().QueryRow("INSERT INTO public.campaigns(created_by, name, discount_rate, voucher_limit, start_date, end_date) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", userID, form.Name, form.DiscountRate, form.VoucherLimit, form.StartDate, form.EndDate).Scan(&campaignID)
// 	return campaignID, err
// }

// One ...
// func (m VoucherModel) One(id string) (campaign Campaign, err error) {
// 	err = db.GetDB().SelectOne(&campaign, "SELECT id, name, discount_rate, voucher_limit, start_date, end_date, created_by FROM campaigns where id = $1", id)
// 	return campaign, err
// }

// All ...
func (m VoucherModel) All(userID string) (campaigns []Voucher, err error) {
	_, err = db.GetDB().Select(&campaigns, "SELECT id, name, discount_rate, voucher_limit, start_date, end_date, created_by FROM campaigns WHERE user_id=$1", userID)
	return campaigns, err
}

// Update ...
// func (m VoucherModel) Update(userID string, id int64, form forms.CreateCampaignForm) (err error) {
// 	//METHOD 1
// 	//Check the campaign by ID using this way
// 	// _, err = m.One(userID, id)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	operation, err := db.GetDB().Exec("UPDATE public.campaign SET name=$2, discount_rate=$3, voucher_limit=$4, start_date=$5, end_date=$6 WHERE id=$1", id, form.Name, form.DiscountRate, form.VoucherLimit, form.StartDate, form.EndDate)
// 	if err != nil {
// 		return err
// 	}

// 	success, _ := operation.RowsAffected()
// 	if success == 0 {
// 		return errors.New("updated 0 records")
// 	}

// 	return err
// }

// Delete ...
// func (m VoucherModel) Delete(userID, id string) (err error) {
// 	operation, err := db.GetDB().Exec("UPDATE campaigns SET deleted_by=$2 WHERE id=$1", id, userID)
// 	if err != nil {
// 		return err
// 	}

// 	success, _ := operation.RowsAffected()
// 	if success == 0 {
// 		return errors.New("no records were deleted")
// 	}

// 	return err
// }
