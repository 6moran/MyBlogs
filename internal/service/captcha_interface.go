package service

// CaptchaService 验证码服务接口
type CaptchaService interface {
	// GenerateCaptcha 生成验证码，返回验证码ID和base64图片
	GenerateCaptcha() (captchaID, captchaImage string, err error)
	// VerifyCaptcha 验证验证码是否正确
	VerifyCaptcha(captchaID, captchaCode string) bool
	// DeleteCaptcha 删除验证码
	DeleteCaptcha(captchaID string) error
}
