package utils

import (
	"encoding/json"
	"mailman/src/config"
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
		"iss":  "SLIIT FOSS",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(expiry).Unix(),
		"data": user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Env.JWTSecret))
	if err != nil {
		panic(fiber.NewError(fiber.StatusInternalServerError, "Error generating jwt token"))
	}
	return t
}

func ValidateUserJWTToken(token string) *models.User {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env.JWTSecret), nil
	})
	if err != nil || !parsedToken.Valid {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid token"))
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid token"))
	}
	user := models.User{}
	jsonString, _ := json.Marshal(claims["data"])
	json.Unmarshal(jsonString, &user)
	return &user
}
