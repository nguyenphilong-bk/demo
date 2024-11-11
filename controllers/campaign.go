package controllers

import (
	"encoding/json"

	"github.com/Massad/gin-boilerplate/forms"
	"github.com/Massad/gin-boilerplate/services"
	"github.com/Massad/gin-boilerplate/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

//CampaignController ...
type CampaignController struct{}

var campaignService = new(services.CampaignService)
var campaignForm = new(forms.CampaignForm)

//Create ...
func (ctrl CampaignController) Create(c *gin.Context) {
	userID := getUserID(c)

	var form forms.CreateCampaignForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := campaignForm.Create(validationErr)
		// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": message})
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Message: message})
		return
	}

	id, err := campaignService.Create(userID, form)
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Campaign could not be created"})
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Message: "Campaign could not be created"})
		return
	}

	result := map[string]interface{}{}
	result["id"] = id
	
	// c.JSON(http.StatusOK, gin.H{"message": "Campaign created", "id": id})
	c.JSON(http.StatusOK, utils.Response{StatusCode: http.StatusOK, Message: "Campaign created successfully", Data: result})
}

//All ...
func (ctrl CampaignController) All(c *gin.Context) {
	// userID := getUserID(c)
	results, err := campaignService.All()
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Could not get campaigns"})
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.RetrieveResponse{StatusCode: http.StatusBadRequest, Message: "Could not get campaigns"})

		return
	}

	data := make([]interface{}, len(results))
	for i, v := range results {
		data[i] = v
	}

	c.JSON(http.StatusOK, utils.RetrieveResponse{StatusCode: http.StatusOK, Data: data})
}

//One ...
func (ctrl CampaignController) One(c *gin.Context) {
	// userID := getUserID(c)
	id := c.Param("id")

	data, err := campaignService.One(id)
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Campaign not found"})
		c.AbortWithStatusJSON(http.StatusNotFound, utils.Response{StatusCode: http.StatusBadRequest, Message: "Campaign not found"})

		return
	}
	
	temp, _ := json.Marshal(&data)
	var result map[string]interface{}
	json.Unmarshal(temp, &result)

	// c.JSON(http.StatusOK, gin.H{"data": data})
	c.JSON(http.StatusOK, utils.Response{StatusCode: http.StatusOK, Data: result})
}

//Update ...
// func (ctrl CampaignController) Update(c *gin.Context) {
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

// 	err = campaignService.Update(userID, getID, form)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Campaign could not be updated"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Campaign updated"})
// }

//Delete ...
func (ctrl CampaignController) Delete(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	err := campaignService.Delete(userID, id)
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Campaign could not be deleted"})
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Message: "Campaign could not be deleted"})

		return
	}

	c.JSON(http.StatusOK, utils.Response{StatusCode: http.StatusOK, Message: "Campaign deleted successfully"})
}
