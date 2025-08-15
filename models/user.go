package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username" binding:"required,min=3,max=20"`
	Password  string             `bson:"password" json:"-"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username        string `json:"username" binding:"required,min=3,max=20"`
	Password        string `json:"password" binding:"required,min=6,max=20"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}



type UserResponse struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}