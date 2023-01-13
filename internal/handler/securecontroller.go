package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Функция-обработчик, работающая только для авторизированных
func SecureExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hi there"})
}
