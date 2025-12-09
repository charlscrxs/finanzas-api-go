package routes

import (
	"github.com/charlscrxs/finanzas-api/controllers"
	"github.com/charlscrxs/finanzas-api/middleware"
	"github.com/gin-gonic/gin"
)

func MovimientosRoutes(r *gin.Engine) {

	auth := r.Group("movimientos")

	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/", controllers.GetMovimientos)
		auth.GET("/:id", controllers.GetMovimientosPorId)

		auth.POST("/", controllers.PostMovimientos)
		auth.PUT("/:id", controllers.PutMovimientos)

		auth.DELETE("/:id", controllers.DeleteMovimientos)
	}

}
