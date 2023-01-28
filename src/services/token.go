package services

import (
	"packform-test/src/database"
	"packform-test/src/models"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidToken(t *jwt.Token, username string) bool {
	claims := t.Claims.(jwt.MapClaims)
	uid := claims["username"].(string)

	return uid == username
}

func ValidUser(username string, p string) bool {
	db := database.DB
	var user = &models.User{Username: username}
	db.First(&user)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

func ExpireToken(t *jwt.Token) bool {
	claims := t.Claims.(jwt.MapClaims)
	for k := range claims {
		delete(claims, k)
	}
	return true
}
