package controllers

import (
	"net/http"
	"time"

	"github.com/charlscrxs/finanzas-api/middleware"
	"github.com/charlscrxs/finanzas-api/migrations"
	"github.com/charlscrxs/finanzas-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func generarToken(usuarioID uint, duracion time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuario_id": usuarioID,
		"exp":        time.Now().Add(duracion).Unix(), // Expira en X tiempo
	})
	return token.SignedString(middleware.JwtKey)
}
func Login(c *gin.Context) {
	var creds models.Credenciales

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Datos no válidos"})
		return
	}

	var usr models.Usuarios
	if err := migrations.DB.Where("email = ?", creds.Email).First(&usr).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Credenciales incorrectas"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(usr.Contrasena), []byte(creds.Contrasena))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Credenciales incorrectas"})
		return
	}

	accessToken, err := generarToken(usr.ID, time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo generar el access token"})
		return
	}

	refreshToken, err := generarToken(usr.ID, time.Hour*24*7)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo generar el refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(ctx *gin.Context) {
	var requestBody struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Refresh token requerido"})
		return
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(requestBody.RefreshToken, &claims, func(t *jwt.Token) (interface{}, error) {
		return middleware.JwtKey, nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Refresh token inválido"})
		return
	}

	usuarioID, ok := claims["usuario_id"].(float64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Token inválido"})
		return
	}

	newAccessToken, err := generarToken(uint(usuarioID), time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo generar el nuevo token"})
		return
	}

	// Devolver el nuevo Access Token
	ctx.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}

func GetUsuarios(c *gin.Context) {
	var Usuarios []models.Usuarios
	if err := migrations.DB.Find(&Usuarios).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error al hacer la petición al servidor"})
		return
	}

	c.JSON(http.StatusOK, Usuarios)
}

func GetUsuariosPorId(c *gin.Context) {
	id := c.Param("id")
	var Usuarios models.Usuarios
	if err := migrations.DB.First(&Usuarios, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No se encontró el usuario con ese ID"})
		return
	}
	c.JSON(http.StatusOK, Usuarios)
}

func HashPasword(pass string) (string, error) {
	bytes, erro := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(bytes), erro
}

func PostUsuarios(c *gin.Context) {
	var Usuario models.Usuarios

	if err := c.ShouldBindJSON(&Usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Datos no válidos"})
		return
	}
	if Usuario.Contrasena == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Contraseña requerida"})

		return
	}

	hashPass, err := HashPasword(Usuario.Contrasena)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo cifrar la conteraseña"})
		return
	}

	Usuario.Contrasena = hashPass

	if err := migrations.DB.Create(&Usuario).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No se pudo crear el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     Usuario.ID,
		"nombre": Usuario.Nombre,
		"email":  Usuario.Email,
	})
}

func PutUsuarios(c *gin.Context) {
	id := c.Param("id")
	var Usuario models.Usuarios

	if err := migrations.DB.First(&Usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No se encontró el usuario"})
		return
	}

	if err := c.ShouldBindJSON(&Usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Datos no válidos"})
		return
	}

	if err := migrations.DB.Save(&Usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, Usuario)
}

func DeleteUsuarios(c *gin.Context) {
	id := c.Param("id")
	var Usuarios models.Usuarios

	if err := migrations.DB.First(&Usuarios, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No se encontró al usuario"})
		return
	}

	if err := migrations.DB.Delete(&Usuarios).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No se pudo eliminar al usuario"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Se eliminó el usuario con éxito"})
}
