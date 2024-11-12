package routes

import (
	"mock-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(r *gin.Engine) {
	payments := r.Group("/payments")
	{
		payments.POST("", controllers.CreatePayment)
	}
}
