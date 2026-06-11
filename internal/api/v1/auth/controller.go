package auth

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/service"
	"MyBlogs/pkg/config"
	bizerrors "MyBlogs/pkg/errors"
	"MyBlogs/pkg/logger"
	"MyBlogs/pkg/response"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	AuthService    service.AuthService
	CaptchaService service.CaptchaService
	Config         *config.Config
}

func NewAuthController(s service.AuthService, cs service.CaptchaService, cfg *config.Config) *AuthController {
	return &AuthController{
		AuthService:    s,
		CaptchaService: cs,
		Config:         cfg,
	}
}

// HandlerLogin 管理员登录
func (ac *AuthController) HandlerLogin(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	// 验证验证码
	if !ac.CaptchaService.VerifyCaptcha(req.CaptchaID, req.CaptchaCode) {
		// 删除验证码
		ac.CaptchaService.DeleteCaptcha(req.CaptchaID)
		response.BizError(c, bizerrors.New(bizerrors.CodeCaptchaError, bizerrors.GetMessage(bizerrors.CodeCaptchaError)))
		return
	}

	// 验证码验证通过后删除
	ac.CaptchaService.DeleteCaptcha(req.CaptchaID)

	resp, err := ac.AuthService.Login(req.Username, req.Password)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("登录失败", zap.String("username", req.Username), zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("登录失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, resp)
}

// HandlerGitHubLogin GitHub OAuth跳转
func (ac *AuthController) HandlerGitHubLogin(c *gin.Context) {
	clientID := ac.Config.GitHub.ClientID
	redirectURL := ac.Config.GitHub.RedirectURL

	url := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=user:email",
		clientID,
		redirectURL,
	)

	c.Redirect(302, url)
}

// HandlerGitHubCallback GitHub OAuth回调
func (ac *AuthController) HandlerGitHubCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		response.BadRequest(c, "缺少code参数")
		return
	}

	resp, err := ac.AuthService.GitHubLogin(code)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("GitHub登录失败", zap.Error(err))
			c.Redirect(302, "/login?error="+err.Error())
			return
		}
		logger.Error("GitHub登录失败", zap.Error(err))
		c.Redirect(302, "/login?error=登录失败，请稍后重试")
		return
	}

	frontendURL := "http://localhost:3000"
	redirectURL := fmt.Sprintf("%s/login?token=%s", frontendURL, resp.Token)
	c.Redirect(302, redirectURL)
}
