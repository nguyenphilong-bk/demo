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

//Update ...
// func (service CampaignService) Update(c *gin.Context) {
// 	userID := getUserID(c)

// 	id := c.Param("id")

// 	getID, err := strconv.ParseInt(id, 10, 64)
// 	if getID == 0 || err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
// 		return
// 	}

// 	var form forms.CreateCampaignForm

// 	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
// 		message := campaignForm.Create(validationErr)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
// 		return
// 	}

// 	err = campaignModel.Update(userID, getID, form)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Campaign could not be updated"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Campaign updated"})
// }

//Delete ...
func (service CampaignService) Delete(userID, id string) error{
	err := campaignModel.Delete(userID, id)
	if err != nil {
		return err
	}

	return nil
}
