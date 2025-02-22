package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		TokenString := ctx.GetHeader("Authorrization")
		if TokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Token Requerido"})
			ctx.Abort()
			return
		}

		TokenString = strings.TrimPrefix(TokenString, "Bearer ")

		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(TokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			ctx.Abort()
			return
		}

		ctx.Set("usuario_id", int(claims["usuario_id"].(float64)))
		ctx.Next()

	}
}
