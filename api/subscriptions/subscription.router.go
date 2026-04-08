package subscriptions

import (
	"github.com/devlopersabbir/juan_don82-server/arch/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup) {
	subs := v1.Group("/v1/subscriptions")
	{
		subs.GET("/plans", GetAllPlans)

		protected := subs.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/purchase", PurchaseSubscription)
			// Admin routes could be here or in admin module
			protected.POST("/plans", AddPlan)
		}
	}
}
