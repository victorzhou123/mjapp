package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"mjbackend/database"
	"mjbackend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CurrencyService struct{}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{}
}

// GetBalance 获取用户算力余额
func (s *CurrencyService) GetBalance(userID primitive.ObjectID) (*models.BalanceResponse, error) {
	balanceCollection := database.GetCollection("currency_balances")
	
	var balance models.CurrencyBalance
	err := balanceCollection.FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&balance)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// 如果用户没有余额记录，创建一个初始余额为0的记录
			balance = models.CurrencyBalance{
				UserID:         userID,
				Balance:        0,
				LastUpdateTime: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}
			_, err = balanceCollection.InsertOne(context.Background(), balance)
			if err != nil {
				return nil, fmt.Errorf("创建用户余额记录失败: %v", err)
			}
		} else {
			return nil, fmt.Errorf("查询用户余额失败: %v", err)
		}
	}
	
	return &models.BalanceResponse{
		Balance:        balance.Balance,
		LastUpdateTime: balance.LastUpdateTime,
	}, nil
}

// DeductBalance 扣减算力
func (s *CurrencyService) DeductBalance(userID primitive.ObjectID, request *models.DeductRequest) (*models.DeductResponse, error) {
	balanceCollection := database.GetCollection("currency_balances")
	transactionCollection := database.GetCollection("currency_transactions")
	
	// 开始事务
	session, err := database.DB.Client().StartSession()
	if err != nil {
		return nil, fmt.Errorf("启动事务失败: %v", err)
	}
	defer session.EndSession(context.Background())
	
	var result *models.DeductResponse
	err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		// 查询当前余额
		var balance models.CurrencyBalance
		err := balanceCollection.FindOne(sc, bson.M{"user_id": userID}).Decode(&balance)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return errors.New("用户余额记录不存在")
			}
			return fmt.Errorf("查询用户余额失败: %v", err)
		}
		
		// 检查余额是否足够
		if balance.Balance < request.Amount {
			return fmt.Errorf("算力余额不足，当前余额: %d，需要: %d", balance.Balance, request.Amount)
		}
		
		// 扣减余额
		newBalance := balance.Balance - request.Amount
		update := bson.M{
			"$set": bson.M{
				"balance":           newBalance,
				"last_update_time": time.Now(),
				"updated_at":        time.Now(),
			},
		}
		
		_, err = balanceCollection.UpdateOne(sc, bson.M{"user_id": userID}, update)
		if err != nil {
			return fmt.Errorf("更新用户余额失败: %v", err)
		}
		
		// 创建交易记录
		transactionID := fmt.Sprintf("tx_%d", time.Now().UnixNano())
		transaction := models.CurrencyTransaction{
			UserID:        userID,
			Type:          "deduct",
			Amount:        request.Amount,
			Reason:        request.Reason,
			MemoID:        request.MemoID,
			TransactionID: transactionID,
			CreatedAt:     time.Now(),
		}
		
		_, err = transactionCollection.InsertOne(sc, transaction)
		if err != nil {
			return fmt.Errorf("创建交易记录失败: %v", err)
		}
		
		result = &models.DeductResponse{
			RemainingBalance: newBalance,
			DeductedAmount:   request.Amount,
			TransactionID:    transactionID,
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	return result, nil
}

// RechargeBalance 充值算力
func (s *CurrencyService) RechargeBalance(userID primitive.ObjectID, request *models.RechargeRequest) (*models.RechargeResponse, error) {
	balanceCollection := database.GetCollection("currency_balances")
	transactionCollection := database.GetCollection("currency_transactions")
	
	// 检查交易ID是否已存在（防止重复充值）
	existingTransaction := transactionCollection.FindOne(context.Background(), bson.M{
		"transaction_id": request.TransactionID,
		"type":           "recharge",
	})
	if existingTransaction.Err() == nil {
		return nil, errors.New("交易ID已存在，请勿重复充值")
	}
	
	// 开始事务
	session, err := database.DB.Client().StartSession()
	if err != nil {
		return nil, fmt.Errorf("启动事务失败: %v", err)
	}
	defer session.EndSession(context.Background())
	
	var result *models.RechargeResponse
	err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		// 查询或创建用户余额记录
		var balance models.CurrencyBalance
		err := balanceCollection.FindOne(sc, bson.M{"user_id": userID}).Decode(&balance)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// 创建新的余额记录
				balance = models.CurrencyBalance{
					UserID:         userID,
					Balance:        0,
					LastUpdateTime: time.Now(),
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}
				_, err = balanceCollection.InsertOne(sc, balance)
				if err != nil {
					return fmt.Errorf("创建用户余额记录失败: %v", err)
				}
			} else {
				return fmt.Errorf("查询用户余额失败: %v", err)
			}
		}
		
		// 增加余额
		newBalance := balance.Balance + request.Amount
		update := bson.M{
			"$set": bson.M{
				"balance":           newBalance,
				"last_update_time": time.Now(),
				"updated_at":        time.Now(),
			},
		}
		
		_, err = balanceCollection.UpdateOne(sc, bson.M{"user_id": userID}, update)
		if err != nil {
			return fmt.Errorf("更新用户余额失败: %v", err)
		}
		
		// 创建交易记录
		internalTransactionID := fmt.Sprintf("tx_%d", time.Now().UnixNano())
		source := request.Source
		if source == "" {
			source = "purchase"
		}
		
		transaction := models.CurrencyTransaction{
			UserID:        userID,
			Type:          "recharge",
			Amount:        request.Amount,
			Reason:        fmt.Sprintf("充值算力 - %s", source),
			TransactionID: request.TransactionID, // 外部交易ID
			Source:        source,
			CreatedAt:     time.Now(),
		}
		
		_, err = transactionCollection.InsertOne(sc, transaction)
		if err != nil {
			return fmt.Errorf("创建交易记录失败: %v", err)
		}
		
		result = &models.RechargeResponse{
			NewBalance:      newBalance,
			RechargedAmount: request.Amount,
			TransactionID:   internalTransactionID,
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	return result, nil
}