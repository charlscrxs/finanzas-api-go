package controllers

import (
	"net/http"

	"github.com/charlstg09/finanzas-api/migrations"
	"github.com/charlstg09/finanzas-api/models"
	"github.com/gin-gonic/gin"
)

func GetMovimientos(c *gin.Context) {
	var Movimientos []models.Movimiento
	if err := migrations.DB.Find(&Movimientos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "no se pudieron traer los movimientos"})
		return
	}
	c.JSON(http.StatusOK, Movimientos)
}

func GetMovimientosPorId(c *gin.Context) {
	id := c.Param("id")
	var Movimientos models.Movimiento

	if err := migrations.DB.First(&Movimientos, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No se encontro el movimiento"})
		return
	}
	c.JSON(http.StatusOK, Movimientos)

}

func PostMovimientos(c *gin.Context) {
	var Movimiento models.Movimiento

	if err := c.ShouldBindJSON(&Movimiento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Datos no Validos"})
		return
	}

	if err := migrations.DB.Create(&Movimiento).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo crear El movimiento"})
		return
	}

	c.JSON(http.StatusOK, Movimiento)

}
func PutMovimientos(c *gin.Context) {
	id := c.Param("id")
	var Movimiento models.Movimiento

	// Buscar si el movimiento existe
	if err := migrations.DB.First(&Movimiento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No se encontró el movimiento"})
		return
	}

	// Intentar hacer el binding de los datos enviados en el request
	if err := c.ShouldBindJSON(&Movimiento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Datos no válidos"})
		return
	}

	// Guardar los cambios
	if err := migrations.DB.Save(&Movimiento).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo actualizar el movimiento"})
		return
	}

	c.JSON(http.StatusOK, Movimiento)

}

func DeleteMovimientos(c *gin.Context) {
	id := c.Param("id")

	var Movimientos models.Movimiento

	if err := migrations.DB.First(&Movimientos, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "NO se encontro el movimiento con ese Id"})
		return
	}

	if err := migrations.DB.Delete(&Movimientos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No se pudo eliminar el movimiento"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Se elimino correctamente el movimiento"})

}
