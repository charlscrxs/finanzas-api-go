package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/charlscrxs/finanzas-api/migrations"
	"github.com/charlscrxs/finanzas-api/routes"

	"github.com/gin-gonic/gin"

	// Swagger
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"

	// Testify (para pruebas unitarias)
	_ "github.com/stretchr/testify/assert"
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
