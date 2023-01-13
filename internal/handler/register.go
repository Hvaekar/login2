package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"login2/pkg/models"
	"login2/pkg/mysql"
	"net/http"
)

// Функция-обработчик для регистрации пользователя
func register(c *gin.Context) {
	var user models.User

	// Принимаем JSON-формат и записываем его с переменную user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возврат ошибки и информации о ней d JSON-формате
		c.Abort()                                                  // Прерывание контекста
		return
	}

	// Шифруем пароль для последующего добавления в БД
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возврат ошибки и информации о ней в JSON
		c.Abort()                                                           // Прерывание контекста
		return
	}

	// Добавление юзера в БД
	stmt := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	res, err := mysql.DB.Exec(stmt, user.Username, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	// Вытаскиваем id и записываем в поле объекта
	id, err := res.LastInsertId()
	user.Id = id
	// Выводим результат в виде JSON
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.Id,
		"username": user.Username,
		"email":    user.Email,
	})
}
