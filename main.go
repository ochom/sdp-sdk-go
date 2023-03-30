package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ochom/sdk-go/cache"
)

func main() {
	token := cache.NewToken()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/health", handleHealth())

	api := router.Group("/api")
	api.Use(authMiddleware())
	{
		api.POST("/subscription/activate", handleSubscriptions(token))
		api.POST("/subscription/deactivate", handleUnSubscriptions(token))
		api.POST("/sms/send-bulk", handleBulkSms(token))
		api.POST("/sms/send-premium", handlePremiumSms(token))
	}

	router.Run(":8080")
}
