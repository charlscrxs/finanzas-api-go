package routes

import (
	"github.com/charlstg09/finanzas-api/controllers"
	"github.com/gin-gonic/gin"
)

func MovimientosRoutes(r *gin.Engine) {
	r.GET("/movimientos", controllers.GetMovimientos)
	r.GET("/movimientos/:id", controllers.GetMovimientosPorId)

	r.POST("/movimientos", controllers.PostMovimientos)
	r.PUT("/movimientos/:id", controllers.PutMovimientos)

	r.DELETE("/movimientos/:id", controllers.DeleteMovimientos)

}
