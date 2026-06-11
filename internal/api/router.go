package api

import (
	"MyBlogs/internal/api/v1/article"
	"MyBlogs/internal/api/v1/auth"
	"MyBlogs/internal/api/v1/comment"
	"MyBlogs/internal/api/v1/tag"
	"MyBlogs/internal/api/v1/user"
	"MyBlogs/internal/service"
	"MyBlogs/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dependencies 依赖注入容器
type Dependencies struct {
	ArticleService service.ArticleService
	AuthService    service.AuthService
	UserService    service.UserService
	TagService     service.TagService
	CommentService service.CommentService
	CaptchaService service.CaptchaService
	Config         *config.Config
}

// NewRouter 创建并配置路由，返回 gin.Engine
func NewRouter(deps Dependencies) *gin.Engine {
	engine := gin.Default()

	engine.StaticFS("/uploads/images", http.Dir("./uploads/images"))

	apiGroup := engine.Group("/api")
	{
		frontGroup := apiGroup.Group("/front")
		adminGroup := apiGroup.Group("/admin")

		// 注册文章路由
		articleController := article.NewArticleController(deps.ArticleService)
		articleRouter := article.NewArticleRouter(articleController)
		articleRouter.Register(frontGroup, adminGroup)

		// 注册认证路由
		authController := auth.NewAuthController(deps.AuthService, deps.CaptchaService, deps.Config)
		captchaController := auth.NewCaptchaController(deps.CaptchaService)
		authRouter := auth.NewAuthRouter(authController, captchaController)
		authRouter.Register(frontGroup)

		// 注册用户路由
		userController := user.NewUserController(deps.UserService)
		userRouter := user.NewUserRouter(userController)
		userRouter.Register(frontGroup, adminGroup)

		// 注册评论路由
		commentController := comment.NewCommentController(deps.CommentService)
		commentRouter := comment.NewCommentRouter(commentController)
		commentRouter.Register(frontGroup, adminGroup)

		// 注册标签路由
		tagController := tag.NewTagController(deps.TagService)
		tagRouter := tag.NewTagRouter(tagController)
		tagRouter.Register(frontGroup, adminGroup)

	}

	return engine
}
