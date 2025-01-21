package http

import (
	"MiniWallet/internal/models"
	"MiniWallet/internal/usecase"
	"MiniWallet/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionHandler(router *gin.RouterGroup, t usecase.TransactionUsecase) {
	handler := &CartHandler{transactionUsecase: t}

	router.POST("/v1/wallet/deposits", handler.Deposit)
	router.POST("/v1/wallet/withdrawals", handler.Withdraw)
	router.GET("/v1/wallet/transactions", handler.ViewAllTransaction)

}

// create deposit transaction
func (h *CartHandler) Deposit(c *gin.Context) {
	userID, err := utils.ExtractUserIDFromToken(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var deposit models.TransactionRequest
	if err := c.ShouldBindJSON(&deposit); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	trans, err := h.transactionUsecase.Deposit(userID, deposit.ReferenceId, deposit.Amount)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, map[string]interface{}{
		"deposit": trans,
	}, "Deposit created successfully")
}

// create withdraw transaction
func (h *CartHandler) Withdraw(c *gin.Context) {
	userID, err := utils.ExtractUserIDFromToken(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var withdraw models.TransactionRequest
	if err := c.ShouldBindJSON(&withdraw); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	trans, err := h.transactionUsecase.Withdraw(userID, withdraw.ReferenceId, withdraw.Amount)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, map[string]interface{}{
		"withdrawal": trans,
	}, "Withdraw created successfully")
}

// view all transaction
func (h *CartHandler) ViewAllTransaction(c *gin.Context) {
	userID, err := utils.ExtractUserIDFromToken(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	transactions, err := h.transactionUsecase.ViewAllTransaction(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, transactions, "Transaction retrieved successfully")
}
