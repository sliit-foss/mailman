package v1

import (
	"mailman/src/database"
	"mailman/src/modules/users/api/v1/dto"
	"mailman/src/modules/users/api/v1/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createUser(c *fiber.Ctx, payload dto.CreateUserReq) *dto.CreateUserRes {
	log.Info("Creating user within system", payload.Email)
	verificationCode := uuid.New().String()
	generatedPassword, _ := password.Generate(8, 2, 1, false, false)
	result, err := database.UseDefault().Collection("users").InsertOne(c.Context(), models.User{
		Email:            payload.Email,
		Name:             payload.Name,
		VerificationCode: verificationCode,
		Password:         generatedPassword,
	}.WithDefaults())
	if err != nil {
		panic(err)
	}
	return &dto.CreateUserRes{
		ID:       result.InsertedID.(primitive.ObjectID),
		Password: generatedPassword,
	}
}
