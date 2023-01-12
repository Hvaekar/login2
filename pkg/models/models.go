package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	pwdInBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // Генереция шифра
	if err != nil {
		return err
	}
	user.Password = string(pwdInBytes) // Перевод в строку
	return nil
}

// Проверка пароля
func (user *User) CheckPassword(password string) error {
	// Сравнение шифра с предостявляемым паролем
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}