package middlewares

import (
	"github.com/gin-gonic/gin"
	"login2/auth"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Берем значение по ключу из заголовка HTTP
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Запрос не содержит токен"})
			c.Abort()
			return
		}
		// Проверям токен из заголовка
		if err := auth.ValidateToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
