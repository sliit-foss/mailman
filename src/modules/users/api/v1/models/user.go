package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole string

const (
	Admin UserRole = "admin"
	Guest UserRole = "guest"
)

type User struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"name,omitempty"`
	Email            string             `json:"email" bson:"email,omitempty"`
	Password         string             `json:"password" bson:"password,omitempty"`
	Organizations    []string           `json:"organizations" bson:"organizations,omitempty"`
	Verified         bool               `json:"verified" bson:"verified,omitempty"`
	VerificationCode string             `json:"verification_code" bson:"verification_code,omitempty"`
	Role             UserRole           `json:"role" bson:"role,omitempty"`
	CreatedAt        string             `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt        string             `json:"updated_at" bson:"updated_at,omitempty"`
}

func (u User) WithDefaults() User {
	if u.Role == "" {
		u.Role = Guest
	}
	if u.Organizations == nil {
		u.Organizations = []string{}
	}
	if u.Verified == false {
		u.Verified = false
	}
	u.CreatedAt = time.Now().Format(time.RFC3339)
	u.UpdatedAt = time.Now().Format(time.RFC3339)
	return u
}
