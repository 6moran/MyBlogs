package api

import (
	"MyBlogs/internal/api/v1/article"
	"MyBlogs/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Dependencies 依赖注入容器
type Dependencies struct {
	ArticleService service.ArticleService
	UserService    service.UserService
	TagService     service.TagService
}

// NewRouter 创建并配置路由，返回 gin.Engine
func NewRouter(deps Dependencies) *gin.Engine {
	engine := gin.Default()

	engine.StaticFS("/uploads/images", http.Dir("./uploads/images"))

	////配置CORS跨域
	//engine.Use()

	apiGroup := engine.Group("/api")
	{
		frontGroup := apiGroup.Group("/front")
		adminGroup := apiGroup.Group("/admin")

		//注册前台和后台api路由
		articleController := article.NewArticleController(deps.ArticleService)
		articleRouter := article.NewArticleRouter(articleController)
		articleRouter.Register(frontGroup, adminGroup)
	}

	return engine
}
