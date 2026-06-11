package auth

import (
	"MyBlogs/internal/service"
	"MyBlogs/pkg/logger"
	"MyBlogs/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CaptchaController 验证码控制器
type CaptchaController struct {
	CaptchaService service.CaptchaService
}

// NewCaptchaController 创建验证码控制器
func NewCaptchaController(cs service.CaptchaService) *CaptchaController {
	return &CaptchaController{CaptchaService: cs}
}

// CaptchaResponse 验证码响应
type CaptchaResponse struct {
	CaptchaID    string `json:"captcha_id"`
	CaptchaImage string `json:"captcha_image"`
}

// HandlerGetCaptcha 获取验证码
func (cc *CaptchaController) HandlerGetCaptcha(c *gin.Context) {
	captchaID, captchaImage, err := cc.CaptchaService.GenerateCaptcha()
	if err != nil {
		logger.Error("生成验证码失败", zap.Error(err))
		response.InternalServerError(c)
		return
	}

	response.Success(c, CaptchaResponse{
		CaptchaID:    captchaID,
		CaptchaImage: captchaImage,
	})
}

// RefreshCaptchaRequest 刷新验证码请求
type RefreshCaptchaRequest struct {
	CaptchaID string `json:"captcha_id" binding:"required"`
}

// HandlerRefreshCaptcha 刷新验证码
func (cc *CaptchaController) HandlerRefreshCaptcha(c *gin.Context) {
	var req RefreshCaptchaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	// 删除旧验证码
	if err := cc.CaptchaService.DeleteCaptcha(req.CaptchaID); err != nil {
		logger.Warn("删除旧验证码失败", zap.Error(err))
	}

	// 生成新验证码
	captchaID, captchaImage, err := cc.CaptchaService.GenerateCaptcha()
	if err != nil {
		logger.Error("生成验证码失败", zap.Error(err))
		response.InternalServerError(c)
		return
	}

	response.Success(c, CaptchaResponse{
		CaptchaID:    captchaID,
		CaptchaImage: captchaImage,
	})
}
