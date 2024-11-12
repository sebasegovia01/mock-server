package routes

import (
	"mock-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAccountRoutes(r *gin.Engine) {
	accounts := r.Group("/accounts")
	{
		accounts.GET("/customer/:customerIdentification", controllers.GetCustomerAccounts)
		accounts.GET("/balance/:productIdentification", controllers.GetAccountBalance)
	}
}
