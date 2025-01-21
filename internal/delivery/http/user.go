package http

import (
	"MiniWallet/internal/models"
	"MiniWallet/internal/usecase"
	"MiniWallet/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase   usecase.UserUsecase
	WalletUsecase usecase.WalletUsecase
}

func NewUserHandler(r *gin.RouterGroup, u usecase.UserUsecase, w usecase.WalletUsecase) {
	handler := &UserHandler{u, w}
	r.POST("/v1/init", handler.Init)
}

func (h *UserHandler) Init(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := h.UserUsecase.Init(&user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.WalletUsecase.CreateWallet(user.CustomerXid)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, map[string]string{
		"token": token,
	}, "User registered successfully")
}
