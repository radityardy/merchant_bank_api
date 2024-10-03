package main

import (
	_ "merchant_bank_api/docs" // This is required for Swagger documentation

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Corrected import path
	ginSwagger "github.com/swaggo/gin-swagger"

	"merchant_bank_api/handlers"
)

// @title Merchant Bank API
// @version 1.0
// @description API for merchant and bank transactions.

// @contact.name API Support
// @contact.email support@merchantbankapi.com

// @host localhost:8080
// @BasePath /

func main() {
	router := gin.Default()

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API Routes
	router.POST("/login", gin.WrapF(handlers.LoginHandler))
	router.POST("/logout", gin.WrapF(handlers.LogoutHandler))
	router.POST("/payment", gin.WrapF(handlers.PaymentHandler))

	router.Run(":8080")
}
