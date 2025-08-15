package services

import (
	"context"
	"errors"
	"time"

	"mjbackend/database"
	"mjbackend/models"
	"mjbackend/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// 用户注册
func (s *UserService) Register(req *models.RegisterRequest) (*models.User, error) {
	// 验证确认密码
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("密码和确认密码不匹配")
	}

	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 检查用户名是否已存在
	var existingUser models.User
	err := collection.FindOne(ctx, bson.M{"username": req.Username}).Decode(&existingUser)
	if err == nil {
		return nil, errors.New("用户名已存在")
	} else if err != mongo.ErrNoDocuments {
		return nil, err
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &models.User{
		ID:        primitive.NewObjectID(),
		Username:  req.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// 用户登录
func (s *UserService) Login(req *models.LoginRequest) (*models.LoginResponse, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 查找用户（仅支持用户名登录）
	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": req.Username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}

	// 验证密码
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
		User: models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
		},
	}, nil
}



// 根据ID获取用户
func (s *UserService) GetUserByID(userID primitive.ObjectID) (*models.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}