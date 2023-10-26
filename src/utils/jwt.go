package utils

import (
	"mailman/src/modules/users/api/v1/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateUserJWTToken(user models.User, refresh bool) string {
	expiry := time.Hour * 1
	if refresh {
		expiry = time.Hour * 24
	}
	claims := jwt.MapClaims{
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(expiry).Unix(),
		"data": user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(fiber.NewError(fiber.StatusInternalServerError, "Error generating jwt token"))
	}
	return t
}
