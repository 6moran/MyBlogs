package service

import (
	"MyBlogs/internal/model/dto/response"
	"MyBlogs/internal/model/entity"
	"MyBlogs/internal/repository"
	"MyBlogs/pkg/config"
	bizerrors "MyBlogs/pkg/errors"
	"MyBlogs/pkg/jwt"
	"MyBlogs/pkg/oauth"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	UserRepo     repository.UserRepo
	GitHubClient *oauth.GitHubClient
}

func NewAuthService(userRepo repository.UserRepo, cfg *config.Config) AuthService {
	return &authService{
		UserRepo:     userRepo,
		GitHubClient: oauth.NewGitHubClient(cfg.GitHub.ClientID, cfg.GitHub.ClientSecret),
	}
}

func (s *authService) Login(username, password string) (*response.LoginResponse, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerrors.New(bizerrors.CodeUserNotFound, bizerrors.GetMessage(bizerrors.CodeUserNotFound))
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, bizerrors.New(bizerrors.CodeUsernameOrPasswordErr, bizerrors.GetMessage(bizerrors.CodeUsernameOrPasswordErr))
	}

	token, err := jwt.GenerateToken(
		user.ID,
		user.Username,
		user.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("生成token失败: %w", err)
	}

	return &response.LoginResponse{
		Token: token,
		User: response.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Email:     user.Email,
			Status:    user.Status,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}

func (s *authService) GitHubLogin(code string) (*response.LoginResponse, error) {
	accessToken, err := s.GitHubClient.GetAccessToken(code)
	if err != nil {
		return nil, fmt.Errorf("获取GitHub access_token失败: %w", err)
	}

	githubUser, err := s.GitHubClient.GetUser(accessToken)
	if err != nil {
		return nil, fmt.Errorf("获取GitHub用户信息失败: %w", err)
	}

	user, err := s.UserRepo.FindByGitHubID(githubUser.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			nickname := githubUser.Name
			if nickname == "" {
				nickname = githubUser.Login
			}

			user = &entity.User{
				Username: githubUser.Login,
				Nickname: nickname,
				Avatar:   githubUser.AvatarURL,
				Email:    githubUser.Email,
				GitHubID: githubUser.ID,
				Role:     0,
			}

			if err := s.UserRepo.Create(user); err != nil {
				return nil, fmt.Errorf("创建用户失败: %w", err)
			}
		} else {
			return nil, fmt.Errorf("查询GitHub用户失败: %w", err)
		}
	}

	token, err := jwt.GenerateToken(
		user.ID,
		user.Username,
		user.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("生成token失败: %w", err)
	}

	return &response.LoginResponse{
		Token: token,
		User: response.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Email:     user.Email,
			Status:    user.Status,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}
