package routes

import (
	"github.com/charlstg09/finanzas-api/controllers"
	"github.com/gin-gonic/gin"
)

func UsuariosRoutes(r *gin.Engine) {

	r.GET("/usuarios", controllers.GetUsuarios)
	r.GET("/usuarios/:id", controllers.GetUsuariosPorId)
	r.POST("/usuarios", controllers.PostUsuarios)
	r.PUT("/usuarios/:id", controllers.PutUsuarios)
	r.DELETE("/usuarios/:id", controllers.DeleteUsuarios)

}
