package routes

import (
	"mock-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPersonalRoutes(r *gin.Engine) {
	personal := r.Group("/personal")
	{
		personal.GET("/identification/:customerIdentification", controllers.GetPersonalIdentification)
	}
}
