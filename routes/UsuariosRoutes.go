package routes

import (
	"github.com/charlscrxs/finanzas-api/controllers"
	"github.com/charlscrxs/finanzas-api/middleware"
	"github.com/gin-gonic/gin"
)

func UsuariosRoutes(r *gin.Engine) {

	r.POST("/usuarios", controllers.PostUsuarios)
	r.POST("/login", controllers.Login)

	auth := r.Group("usuarios")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/", controllers.GetUsuarios)
		auth.GET("/:id", controllers.GetUsuariosPorId)
		auth.PUT("/:id", controllers.PutUsuarios)
		auth.DELETE("/:id", controllers.DeleteUsuarios)
	}

}
