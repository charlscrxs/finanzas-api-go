package routes

import (
	"github.com/charlscrxs/finanzas-api/controllers"
	"github.com/charlscrxs/finanzas-api/middleware"
	"github.com/gin-gonic/gin"
)

func CategoriasRoutes(r *gin.Engine) {

	auth := r.Group("categorias")

	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/", controllers.GetCategorias)
		auth.GET("/:id", controllers.GetCategoriasPorId)
		auth.POST("/", controllers.PostCategorias)
		auth.PUT("/:id", controllers.PutCategorias)
		auth.DELETE("/:id", controllers.DeleteCategorias)

	}

}
