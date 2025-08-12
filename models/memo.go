package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Memo struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	Title      string             `bson:"title" json:"title" binding:"required"`
	Content    string             `bson:"content" json:"content"`
	CreatedAt  time.Time          `bson:"created_at" json:"createTime"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updateTime"`
}

type CreateMemoRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type UpdateMemoRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type MemoListResponse struct {
	List  []Memo `json:"list"`
	Total int64  `json:"total"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}