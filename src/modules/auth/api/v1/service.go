package v1

import (
	"fmt"
	"mailman/src/modules/auth/api/v1/dto"
	user "mailman/src/modules/users/api/v1"
	"mailman/src/modules/users/api/v1/models"
	"mailman/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func registerUser(c *fiber.Ctx, payload dto.RegisterReq) *dto.LoginRes {
	insertedID := user.Repository().Create(models.User{
		Email:    payload.Email,
		Name:     payload.Name,
		Verified: true,
		Password: utils.HashStr(payload.Password),
	}.WithDefaults())
	log.Info(fmt.Sprintf("Account creation successful. Email - %s Generating tokens....", payload.Email), insertedID)
	user := user.Repository().FindByID(insertedID)
	accessToken := utils.GenerateUserJWTToken(user, false)
	refreshToken := utils.GenerateUserJWTToken(user, true)
	return &dto.LoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}
}
