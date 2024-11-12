package controllers

import (
	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

//CampaignController ...
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

//One ...
// func (ctrl VoucherController) One(c *gin.Context) {
// 	// userID := getUserID(c)
// 	id := c.Param("id")

// 	data, err := campaignModel.One(id)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Campaign not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": data})
// }

//Update ...
// func (ctrl VoucherController) Update(c *gin.Context) {
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
// func (ctrl VoucherController) Delete(c *gin.Context) {
// 	userID := getUserID(c)

// 	id := c.Param("id")

// 	err := campaignModel.Delete(userID, id)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Campaign could not be deleted"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted"})

// }
