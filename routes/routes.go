package routes

import (
	"mock-server/config"
	"mock-server/controllers"
	"mock-server/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(r *gin.Engine, cfg *config.Config) {

	api := r.Group("/api/v1")

	healthController := controllers.NewHealthController(cfg)
	api.GET("/health", healthController.HealthCheck)

	// Group for personal routes
	personal := api.Group("/personal")
	{
		personal.GET("/identification/:customerIdentification", WithTraceability(controllers.GetPersonalIdentification))
	}

	// Group for account routes
	accounts := api.Group("/accounts")
	{
		accounts.GET("/customer/:customerIdentification", WithTraceability(controllers.GetCustomerAccounts))
		accounts.GET("/balance/:productIdentification", WithTraceability(controllers.GetAccountBalance))
	}

	// Group for payment routes
	payments := api.Group("/payments")
	{
		payments.POST("", WithTraceability(controllers.CreatePayment))
	}
}

// Used for header traceability - add in paths where required
func WithTraceability(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.TraceabilityMiddleware()(c)
		if !c.IsAborted() {
			handler(c)
		}
	}
}
