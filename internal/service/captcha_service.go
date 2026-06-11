package service

import (
	"MyBlogs/pkg/databases"
	"context"
	"fmt"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
)

// redisStore 自定义Redis存储实现，用于验证码
type redisStore struct {
	client     *redis.Client
	prefix     string
	expiration time.Duration
}

// NewRedisStore 创建Redis存储实例
func NewRedisStore(client *redis.Client, prefix string, expiration time.Duration) *redisStore {
	return &redisStore{
		client:     client,
		prefix:     prefix,
		expiration: expiration,
	}
}

// Set 设置验证码
func (s *redisStore) Set(id string, value string) error {
	ctx := context.Background()
	key := s.prefix + id
	err := s.client.Set(ctx, key, value, s.expiration).Err()
	if err != nil {
		return fmt.Errorf("设置验证码失败: %w", err)
	}
	return nil
}

// Get 获取验证码
func (s *redisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	key := s.prefix + id
	value, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	if clear {
		s.client.Del(ctx, key)
	}
	return value
}

// Verify 验证验证码
func (s *redisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}

type captchaService struct {
	store *redisStore
}

// NewCaptchaService 创建验证码服务
func NewCaptchaService() CaptchaService {
	// 配置Redis存储
	store := NewRedisStore(databases.RedisClient, "captcha:", 5*time.Minute)

	return &captchaService{
		store: store,
	}
}

// GenerateCaptcha 生成验证码
func (s *captchaService) GenerateCaptcha() (captchaID, captchaImage string, err error) {
	// 配置验证码驱动
	driver := base64Captcha.NewDriverDigit(
		80,  // 高度
		240, // 宽度
		5,   // 验证码长度（5位数字）
		0.7, // 最大倾斜角度
		80,  // 背景圆点数量
	)

	// 创建验证码实例
	c := base64Captcha.NewCaptcha(driver, s.store)

	// 生成验证码
	captchaID, captchaImage, _, err = c.Generate()
	if err != nil {
		return "", "", fmt.Errorf("生成验证码失败: %w", err)
	}

	return captchaID, captchaImage, nil
}

// VerifyCaptcha 验证验证码
func (s *captchaService) VerifyCaptcha(captchaID, captchaCode string) bool {
	c := base64Captcha.NewCaptcha(nil, s.store)
	return c.Verify(captchaID, captchaCode, true)
}

// DeleteCaptcha 删除验证码
func (s *captchaService) DeleteCaptcha(captchaID string) error {
	ctx := context.Background()
	key := "captcha:" + captchaID
	err := s.store.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("删除验证码失败: %w", err)
	}
	return nil
}
