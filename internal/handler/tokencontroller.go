package handler

import (
	"github.com/gin-gonic/gin"
	"login2/auth"
	"login2/pkg/models"
	"login2/pkg/models/mysql"
	"net/http"
)

// Создание структуры для запроса токена
type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Функция-обработчик для генерирование токена
func GenerateToken(c *gin.Context) {
	var req TokenRequest
	var user models.User

	// Парсим запрос и записываем данные в переменную req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Проверяем email на наличие в БД
	stmt := "SELECT password FROM users WHERE email = ?"
	row := mysql.DB.QueryRow(stmt, req.Email)
	if err := row.Scan(&user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Проверяем пароль на соответствие в БД
	if err := user.CheckPassword(req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		c.Abort()
		return
	}

	// Генерируем JWT-токен, если все нормально
	token, err := auth.GenerateJWT(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Отправляем ответ, если все норм
	c.JSON(http.StatusOK, gin.H{"token": token})
}
