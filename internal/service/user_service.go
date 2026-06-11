package service

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/model/dto/response"
	"MyBlogs/internal/model/entity"
	"MyBlogs/internal/repository"
	bizerrors "MyBlogs/pkg/errors"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	UserRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		UserRepo: userRepo,
	}
}

func (s *userService) GetUserInfo(userID int) (*response.UserResponse, error) {
	user, err := s.UserRepo.FindByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerrors.New(bizerrors.CodeUserNotFound, bizerrors.GetMessage(bizerrors.CodeUserNotFound))
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	return &response.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Status:    user.Status,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *userService) CreateUser(req request.CreateUserRequest) error {
	existing, _ := s.UserRepo.FindByUsername(req.Username)
	if existing != nil {
		return bizerrors.New(bizerrors.CodeUsernameAlreadyExists, "用户名已存在")
	}

	user := &entity.User{
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
		Role:     req.Role,
	}

	if req.Password != "" {
		hashedPassword, err := HashPassword(req.Password)
		if err != nil {
			return fmt.Errorf("密码加密失败: %w", err)
		}
		user.Password = hashedPassword
	}

	return s.UserRepo.Create(user)
}

func (s *userService) UpdateUser(id int, req request.UpdateUserRequest) error {
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bizerrors.New(bizerrors.CodeUserNotFound, "用户不存在")
		}
		return fmt.Errorf("查询用户失败: %w", err)
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Status != nil {
		user.Status = *req.Status
	}
	if req.Role != nil {
		user.Role = *req.Role
	}
	if req.Password != "" {
		hashedPassword, err := HashPassword(req.Password)
		if err != nil {
			return fmt.Errorf("密码加密失败: %w", err)
		}
		user.Password = hashedPassword
	}

	return s.UserRepo.Update(user)
}

func (s *userService) DeleteUser(id int) error {
	_, err := s.UserRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bizerrors.New(bizerrors.CodeUserNotFound, "用户不存在")
		}
		return fmt.Errorf("查询用户失败: %w", err)
	}

	return s.UserRepo.Delete(id)
}

func (s *userService) GetUserByID(id int) (*response.UserResponse, error) {
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerrors.New(bizerrors.CodeUserNotFound, "用户不存在")
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	return &response.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Status:    user.Status,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *userService) GetUserList(req request.QueryUserRequest) ([]response.UserResponse, int64, error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 || req.Size > 100 {
		req.Size = 10
	}

	users, total, err := s.UserRepo.FindWithCondition(req)
	if err != nil {
		return nil, 0, fmt.Errorf("查询用户列表失败: %w", err)
	}

	var resp []response.UserResponse
	for _, user := range users {
		resp = append(resp, response.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Email:     user.Email,
			Status:    user.Status,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		})
	}

	return resp, total, nil
}

// HashPassword 密码加密
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// InitAdmin 初始化管理员账号
func InitAdmin(userRepo repository.UserRepo, username, password string) error {
	_, err := userRepo.FindByUsername(username)
	if err == nil {
		return nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	return userRepo.Create(&entity.User{
		Username: username,
		Password: hashedPassword,
		Nickname: "管理员",
		Role:     1,
	})
}
