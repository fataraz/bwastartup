package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// get query string lalu dijadikan int
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaings, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaings", http.StatusOK, "success", campaign.FormatCampaigns(campaings))
	c.JSON(http.StatusOK, response)
}
