package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateUserReq struct {
	Name          string `validate:"required"`
	Email         string `validate:"required,email"`
	Role          string
	Organizations []string
}

type CreateUserRes struct {
	ID       primitive.ObjectID `json:"_id"`
	Password string             `json:"password"`
}
