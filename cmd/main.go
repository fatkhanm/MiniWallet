package main

import (
	"log"

	deliveryhttp "MiniWallet/internal/delivery/http"
	"MiniWallet/internal/repository"
	"MiniWallet/internal/usecase"
	"MiniWallet/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := utils.ConnectDB() // Assume a helper function to connect to DB

	userRepo := repository.NewUserRepository(db)
	walletRepo := repository.NewWalletRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	walletUsecase := usecase.NewWalletUsecase(walletRepo)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo, walletRepo)

	// Public routes
	public := r.Group("/api")
	{
		deliveryhttp.NewUserHandler(public, userUsecase, walletUsecase)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(utils.AuthMiddleware())
	{
		deliveryhttp.NewTransactionHandler(public, transactionUsecase)
		deliveryhttp.NewWalletHandler(public, walletUsecase)
	}

	log.Println("Server running at :8080")
	log.Fatal(r.Run(":8080"))
}
