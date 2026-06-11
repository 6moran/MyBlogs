package captcha

import (
	"github.com/mojocn/base64Captcha"
)

// Captcha 验证码管理器
type Captcha struct {
	store base64Captcha.Store
}

// NewCaptcha 创建验证码管理器
func NewCaptcha() *Captcha {
	return &Captcha{
		store: base64Captcha.DefaultMemStore,
	}
}

// Generate 生成验证码，返回(id, base64图片, answer)
func (c *Captcha) Generate() (string, string, string, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, c.store)
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return "", "", "", err
	}
	return id, b64s, answer, nil
}

// Verify 验证验证码
func (c *Captcha) Verify(id, answer string, clear bool) bool {
	return c.store.Verify(id, answer, clear)
}
