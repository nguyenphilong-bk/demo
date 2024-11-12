package controllers

import (
	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

//VoucherController ...
type VoucherController struct{}

var voucherService = new(services.VoucherService)

//Create ...
func (ctrl VoucherController) Create(c *gin.Context) {
	// userID := getUserID(c)

	var form forms.CreateCampaignForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := campaignForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	// id, err := voucherModel.Create(userID, form)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Campaign could not be created"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Campaign created", "id": id})
}

//All ...
func (ctrl VoucherController) All(c *gin.Context) {
	userID := getUserID(c)
	results, err := voucherService.All(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Could not get campaigns"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}