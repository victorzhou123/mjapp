package services

import (
	"context"
	"errors"
	"time"

	"mjbackend/database"
	"mjbackend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MemoService struct{}

func NewMemoService() *MemoService {
	return &MemoService{}
}

// 创建备忘录
func (s *MemoService) CreateMemo(userID primitive.ObjectID, req *models.CreateMemoRequest) (*models.Memo, error) {
	collection := database.GetCollection("memos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	memo := &models.Memo{
		ID:        primitive.NewObjectID(),
		UserID:    userID,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := collection.InsertOne(ctx, memo)
	if err != nil {
		return nil, err
	}

	return memo, nil
}

// 获取备忘录列表
func (s *MemoService) GetMemoList(userID primitive.ObjectID, page, limit int, keyword string) (*models.MemoListResponse, error) {
	collection := database.GetCollection("memos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 构建查询条件
	filter := bson.M{"user_id": userID}
	if keyword != "" {
		filter["$or"] = []bson.M{
			{"title": bson.M{"$regex": keyword, "$options": "i"}},
			{"content": bson.M{"$regex": keyword, "$options": "i"}},
		}
	}

	// 计算总数
	total, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	// 计算跳过的文档数
	skip := (page - 1) * limit

	// 查询选项
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"created_at", -1}}) // 按创建时间倒序
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	// 查询数据
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var memos []models.Memo
	if err = cursor.All(ctx, &memos); err != nil {
		return nil, err
	}

	return &models.MemoListResponse{
		List:  memos,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

// 获取备忘录详情
func (s *MemoService) GetMemoByID(userID, memoID primitive.ObjectID) (*models.Memo, error) {
	collection := database.GetCollection("memos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var memo models.Memo
	filter := bson.M{
		"_id":     memoID,
		"user_id": userID,
	}

	err := collection.FindOne(ctx, filter).Decode(&memo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("备忘录不存在")
		}
		return nil, err
	}

	return &memo, nil
}

// 更新备忘录
func (s *MemoService) UpdateMemo(userID, memoID primitive.ObjectID, req *models.UpdateMemoRequest) (*models.Memo, error) {
	collection := database.GetCollection("memos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 检查备忘录是否存在且属于当前用户
	filter := bson.M{
		"_id":     memoID,
		"user_id": userID,
	}

	update := bson.M{
		"$set": bson.M{
			"title":      req.Title,
			"content":    req.Content,
			"updated_at": time.Now(),
		},
	}

	result := collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("备忘录不存在")
		}
		return nil, result.Err()
	}

	var memo models.Memo
	if err := result.Decode(&memo); err != nil {
		return nil, err
	}

	return &memo, nil
}

// 删除备忘录
func (s *MemoService) DeleteMemo(userID, memoID primitive.ObjectID) error {
	collection := database.GetCollection("memos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"_id":     memoID,
		"user_id": userID,
	}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("备忘录不存在")
	}

	return nil
}