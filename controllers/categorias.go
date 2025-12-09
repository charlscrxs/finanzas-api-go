package controllers

import (
	"net/http"

	"github.com/charlscrxs/finanzas-api/migrations"
	"github.com/charlscrxs/finanzas-api/models"
	"github.com/gin-gonic/gin"
)

func GetCategorias(c *gin.Context) {
	var categorias []models.Categorias

	if err := migrations.DB.Find(&categorias).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error al consultar las categorias"})
		return
	}
	c.JSON(http.StatusOK, categorias)

}

func GetCategoriasPorId(c *gin.Context) {
	id := c.Param("id")
	var categorias models.Categorias

	if err := migrations.DB.First(&categorias, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No Se encontro la categoria con ese id"})
		return
	}
	c.JSON(http.StatusOK, categorias)

}

func PostCategorias(c *gin.Context) {
	var Categoria models.Categorias

	if err := c.ShouldBindJSON(&Categoria); err != nil {

		if err := c.ShouldBindJSON(&Categoria); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		return
	}
	if err := migrations.DB.Create(&Categoria).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo guardar La categoria"})

		return
	}

	c.JSON(http.StatusOK, Categoria)

}

func PutCategorias(c *gin.Context) {

	id := c.Param("id")

	var Categoria models.Categorias

	if err := migrations.DB.First(&Categoria, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NO se encontro la categoria"})
		return
	}

	if err := c.ShouldBindJSON(&Categoria); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Datos no validos"})
		return
	}

	if err := migrations.DB.Save(&Categoria).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "NO se pudo guardar la categoria"})
		return
	}

	c.JSON(http.StatusOK, Categoria)

}

func DeleteCategorias(c *gin.Context) {
	id := c.Param("id")

	var CT models.Categorias

	if err := migrations.DB.First(&CT, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se encontro la categoria"})
		return
	}

	if err := migrations.DB.Delete(&CT).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No se pudo eliminar la categoria"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Mensaje": "Se ha eliminado la Categoria con exito"})

}
