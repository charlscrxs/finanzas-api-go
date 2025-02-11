package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hola mundo")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mensaje": "servidor corriendo"})
	})

	P := ":8080"

	fmt.Println("servidor corriendo en el puerto:", P)
	if err := r.Run(P); err != nil {
		log.Fatal("error de conexion", err)
	}
}
