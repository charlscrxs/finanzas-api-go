package controllers

import (
	"net/http"

	"github.com/charlstg09/finanzas-api/migrations"
	"github.com/charlstg09/finanzas-api/models"
	"github.com/gin-gonic/gin"
)

func GetUsuarios(c *gin.Context) {
	var Usuarios []models.Usuarios
	if err := migrations.DB.Find(&Usuarios).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "error al hacer la peticion al servidor"})
		return
	}

	c.JSON(http.StatusOK, Usuarios)

}

func GetUsuariosPorId(c *gin.Context) {
	id := c.Param("id")
	var Usuarios models.Usuarios
	if err := migrations.DB.First(&Usuarios, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No Se encontro el usuario con ese id"})
		return
	}
	c.JSON(http.StatusOK, Usuarios)

}

func PostUsuarios(c *gin.Context) {
	var Usuario models.Usuarios

	if err := c.ShouldBindJSON(&Usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Datos no Validos"})
		return
	}

	if err := migrations.DB.Create(&Usuario).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No se pudo crear el usuario"})
		return
	}

	c.JSON(http.StatusOK, Usuario)
}

func PutUsuarios(c *gin.Context) {

	id := c.Param("id")
	var Usuario models.Usuarios

	if err := migrations.DB.First(&Usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No se encontro el usuario"})
		return

	}

	if err := c.ShouldBindJSON(&Usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Datos no validos"})
		return
	}

	if err := migrations.DB.Save(&Usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo crear el usuario"})

		return
	}

	c.JSON(http.StatusOK, Usuario)

}

func DeleteUsuarios(c *gin.Context) {
	id := c.Param("id")
	var Usuarios models.Usuarios

	if err := migrations.DB.First(&Usuarios, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No se encontro al usuario"})
		return
	}

	if err := migrations.DB.Delete(&Usuarios).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo eliminar al usuario"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Se elimino el usuario con exito"})

}
