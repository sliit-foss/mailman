package dto

type LoginReq struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
