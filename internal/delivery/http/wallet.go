package http

import (
	"MiniWallet/internal/usecase"
	"MiniWallet/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	walletUsecase usecase.WalletUsecase
}

// NewProductHandler initializes the product handler and routes
func NewWalletHandler(router *gin.RouterGroup, walletUsecase usecase.WalletUsecase) {
	handler := &WalletHandler{walletUsecase: walletUsecase}

	router.POST("/v1/wallet", handler.EnabledWallet)
	router.GET("/v1/wallet", handler.ViewBalance)
	router.PATCH("/v1/wallet", handler.DisableWallet)
}

// Enable Wallet
func (h *WalletHandler) EnabledWallet(c *gin.Context) {
	userID, err := utils.ExtractUserIDFromToken(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	_, err = h.walletUsecase.UpdateEnableStatus(userID, true)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Wallet enabled successfully")
}

// view wallet balance
func (h *WalletHandler) ViewBalance(c *gin.Context) {
	userID, err := utils.ExtractUserIDFromToken(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	balance, err := h.walletUsecase.GetWalletAmount(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, map[string]float64{
		"balance": balance.Balance,
	}, "Wallet balance retrieved successfully")
}

// disable wallet
func (h *WalletHandler) DisableWallet(c *gin.Context) {
	userID, err := utils.ExtractUserIDFromToken(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	wallet, err := h.walletUsecase.UpdateEnableStatus(userID, false)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, wallet, "Wallet disabled successfully")
}
