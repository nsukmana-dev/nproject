package handler

import (
	"net/http"
	"nproject/transaction"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) Index(c *gin.Context) {
	transactions, err := h.transactionService.GetAllTransaction()
	if err != nil {
		kodeErr := strconv.Itoa(http.StatusInternalServerError)
		nameErr := "Cannot get transactions"
		linkErr := "transactions"
		errorStatus := ErrorData(kodeErr, nameErr, linkErr)
		c.HTML(http.StatusInternalServerError, "error.html", errorStatus)
		return
	}

	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"transactions": transactions})
}
