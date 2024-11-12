package services

import (
	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/models"
)

//CampaignService ...
type CampaignService struct{}

var campaignModel = new(models.CampaignModel)

//Create ...
func (service CampaignService) Create(userID string, form forms.CreateCampaignForm) (string, error){
	id, err := campaignModel.Create(userID, form)
	if err != nil {
		return "", err
	}

	return id, nil
}

//All ...
func (service CampaignService) All() ([]models.Campaign, error) {
	results, err := campaignModel.All()
	if err != nil {
		return results, err
	}

	return results, nil
}

//One ...
func (service CampaignService) One(id string) (models.Campaign, error) {
	data, err := campaignModel.One(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

//Delete ...
func (service CampaignService) Delete(userID, id string) error{
	err := campaignModel.Delete(userID, id)
	if err != nil {
		return err
	}

	return nil
}
