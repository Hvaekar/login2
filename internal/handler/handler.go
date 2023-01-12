package handler

import (
	"github.com/gin-gonic/gin"
	"login2/middlewares"
)

type Handler struct {

}

// Инициализая маршрутизатора и обработка путей
func InitRoutes() *gin.Engine {
	router := gin.Default() // Создания маршрутизатора

	//router.GET("/")
	auth := router.Group("/auth")
	{
		auth.POST("/register", register)
		auth.POST("/token", GenerateToken)
		sec := auth.Group("/secured").Use(middlewares.Auth())
		{
			sec.GET("/example", SecureExample)
		}
	}

	return router
}