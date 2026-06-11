package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	AuthController    *AuthController
	CaptchaController *CaptchaController
}

func NewAuthRouter(ac *AuthController, cc *CaptchaController) *AuthRouter {
	return &AuthRouter{
		AuthController:    ac,
		CaptchaController: cc,
	}
}

func (ar *AuthRouter) Register(front *gin.RouterGroup) {
	authGroup := front.Group("/auth")
	{
		authGroup.POST("/login", ar.AuthController.HandlerLogin)
		authGroup.GET("/github", ar.AuthController.HandlerGitHubLogin)
		authGroup.GET("/github/callback", ar.AuthController.HandlerGitHubCallback)

		// 验证码路由
		authGroup.GET("/captcha", ar.CaptchaController.HandlerGetCaptcha)
		authGroup.POST("/captcha/refresh", ar.CaptchaController.HandlerRefreshCaptcha)
	}
}
