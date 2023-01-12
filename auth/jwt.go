package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("jwtsecretkey")

// Создание модели JWT-утверждений(?)
type JWTClaim struct {
	Username string `json:"username"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// Генерирование JWT-токена
func GenerateJWT(email string, username string) (tokenStr string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour) // Определение дедлайна для токена
	// Создание экземпляра модели JWTClaim и заполнение полей
	claims := &JWTClaim{
		Email: email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Установка ранее созданного дедлайна
		},
	}
	// Генерирование нового токена на основе метода SigningMethodES256 с использованием данных модели claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Подписание токена с использованим секретного ключа
	tokenStr, err = token.SignedString(jwtKey)
	return
}

// Проверка/валидация токена
func ValidateToken(signedToken string) (err error) {
	// Парсинг переданного токена
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
		)
	if err != nil {
		return
	}

	// Извлекаем утверждения
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	// Проверям истек ли срок действия токена или нет
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}