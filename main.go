package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ochom/sdk-go/cache"
	"github.com/ochom/sdk-go/dto"
	"github.com/ochom/sdk-go/services"
)

func handleHealth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// set allowed headers
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-api-key")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		xcsrf := ctx.GetHeader("x-api-key")
		if xcsrf == "" {
			ctx.AbortWithStatus(401)
			return
		}

		xtoken := os.Getenv("ACCESS_TOKEN")
		if xtoken != xcsrf {
			ctx.AbortWithStatus(401)
			return
		}

		ctx.Next()
	}
}

func authenticate(token *cache.Token, username, password string) (string, error) {
	tokenVal := token.GetVal()
	if tokenVal == "" {
		token.SetGenerating()
		newTokenVal, err := services.Authenticate(username, password)
		if err != nil {
			return "", err
		}
		token.SetVal(newTokenVal)
	}

	return token.GetVal(), nil
}

func handleSubscriptions(token *cache.Token) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.SubscriptionRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		accessToken, err := authenticate(token, req.Username, req.Password)
		if err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		res, err := services.Subscribe(accessToken, &req)
		if err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		ctx.JSON(200, res)
	}
}

func handleUnSubscriptions(token *cache.Token) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.SubscriptionRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		accessToken, err := authenticate(token, req.Username, req.Password)
		if err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		res, err := services.UnSubscribe(accessToken, &req)
		if err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		ctx.JSON(200, res)
	}
}

func handlePremiumSms(token *cache.Token) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.PremiumSmsRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		accessToken, err := authenticate(token, req.Username, req.Password)
		if err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		res, err := services.SendPremiumSms(accessToken, &req)
		if err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		ctx.JSON(200, res)
	}
}

func handleBulkSms(token *cache.Token) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.BulkSmsRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		accessToken, err := authenticate(token, req.Username, req.Password)
		if err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		res, err := services.SendBulkSms(accessToken, &req)
		if err != nil {
			ctx.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		ctx.JSON(200, res)

	}
}

func main() {
	token := cache.NewToken()

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
