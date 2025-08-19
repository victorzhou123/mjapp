package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CurrencyBalance 用户算力余额模型
type CurrencyBalance struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID         primitive.ObjectID `bson:"user_id" json:"-"`
	Balance        int                `bson:"balance" json:"balance"`
	LastUpdateTime time.Time          `bson:"last_update_time" json:"lastUpdateTime"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updatedAt"`
}

// CurrencyTransaction 算力交易记录模型
type CurrencyTransaction struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID        primitive.ObjectID `bson:"user_id" json:"-"`
	Type          string             `bson:"type" json:"type"` // "deduct" 或 "recharge"
	Amount        int                `bson:"amount" json:"amount"`
	Reason        string             `bson:"reason" json:"reason"`
	MemoID        *primitive.ObjectID `bson:"memo_id,omitempty" json:"memoId,omitempty"`
	TransactionID string             `bson:"transaction_id" json:"transactionId"`
	Source        string             `bson:"source,omitempty" json:"source,omitempty"`
	CreatedAt     time.Time          `bson:"created_at" json:"createdAt"`
}

// DeductRequest 扣减算力请求模型
type DeductRequest struct {
	Amount int                `json:"amount" binding:"required,min=1"`
	Reason string             `json:"reason" binding:"required"`
	MemoID *primitive.ObjectID `json:"memoId,omitempty"`
}

// RechargeRequest 充值算力请求模型
type RechargeRequest struct {
	Amount        int    `json:"amount" binding:"required,min=1"`
	TransactionID string `json:"transactionId" binding:"required"`
	Source        string `json:"source"`
}

// BalanceResponse 余额查询响应模型
type BalanceResponse struct {
	Balance        int       `json:"balance"`
	LastUpdateTime time.Time `json:"lastUpdateTime"`
}

// DeductResponse 扣减算力响应模型
type DeductResponse struct {
	RemainingBalance int    `json:"remainingBalance"`
	DeductedAmount   int    `json:"deductedAmount"`
	TransactionID    string `json:"transactionId"`
}

// RechargeResponse 充值算力响应模型
type RechargeResponse struct {
	NewBalance      int    `json:"newBalance"`
	RechargedAmount int    `json:"rechargedAmount"`
	TransactionID   string `json:"transactionId"`
}

// InsufficientBalanceError 余额不足错误响应模型
type InsufficientBalanceError struct {
	CurrentBalance  int `json:"currentBalance"`
	RequiredAmount  int `json:"requiredAmount"`
}