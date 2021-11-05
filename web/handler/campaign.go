package handler

import (
	"net/http"
	"nproject/campaign"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) Index(c *gin.Context) {
	campaigns, err := h.campaignService.GetCampaigns(0)
	if err != nil {
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Cannot get campaigns"
		linkErr := "users"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}

	c.HTML(http.StatusOK, "campaign_index.html", gin.H{"campaigns": campaigns})
}
