package handler

import (
	"net/http"
	"nproject/helper"
	"nproject/transaction"
	"nproject/user"

	"github.com/gin-gonic/gin"
)

type transactionHendler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHendler {
	return &transactionHendler{service}
}

func (h *transactionHendler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
