package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/charlstg09/finanzas-api/migrations"
	"github.com/charlstg09/finanzas-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hola mundo")

	migrations.ConectarBaseDeDatos()

	r := gin.Default()

	routes.CategoriasRoutes(r)
	routes.MovimientosRoutes(r)
	routes.UsuariosRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mensaje": "servidor corriendo"})
	})

	P := ":8080"

	fmt.Println("servidor corriendo en el puerto:", P)
	if err := r.Run(P); err != nil {
		log.Fatal("error de conexion", err)
	}
}
