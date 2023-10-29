package v1

import (
	"mailman/src/modules/users/api/v1/dto"
	"mailman/src/modules/users/api/v1/models"
	"mailman/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
)

func createUser(c *fiber.Ctx, payload dto.CreateUserReq) *dto.CreateUserRes {
	log.Info("Creating user within system - ", payload.Email)
	verificationCode := uuid.New().String()
	generatedPassword, _ := password.Generate(8, 2, 1, false, false)
	insertedID := repository.Create(models.User{
		Email:            payload.Email,
		Name:             payload.Name,
		VerificationCode: &verificationCode,
		Password:         utils.HashStr(generatedPassword),
	}.WithDefaults())
	return &dto.CreateUserRes{
		ID:       insertedID,
		Password: generatedPassword,
	}
}
