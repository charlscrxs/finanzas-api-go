package routes

import (
	"github.com/charlstg09/finanzas-api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoriasRoutes(r *gin.Engine) {

	r.GET("/categorias", controllers.GetCategorias)
	r.GET("/categorias/:id", controllers.GetCategoriasPorId)

	r.POST("/categorias", controllers.PostCategorias)

	r.PUT("/categorias/:id", controllers.PutCategorias)

	r.DELETE("/categorias/:id", controllers.DeleteCategorias)

}
