package article

import (
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
	ArticleController *Controller
}

func NewArticleRouter(ac *Controller) *ArticleRouter {
	return &ArticleRouter{
		ArticleController: ac,
	}
}

func (ar *ArticleRouter) Register(front, admin *gin.RouterGroup) {
	articleAdminGroup := admin.Group("/articles")
	{
		articleAdminGroup.POST("/", ar.ArticleController.HandlerCreateArticle)
		articleAdminGroup.DELETE("/:id", ar.ArticleController.HandlerDeleteArticle)
		articleAdminGroup.POST("/images", ar.ArticleController.HandlerImage)
		articleAdminGroup.GET("/", ar.ArticleController.HandlerGetArticleLimit)
		articleAdminGroup.GET("/:id")
	}

	articleFrontGroup := front.Group("/articles")
	{
		articleFrontGroup.GET("", ar.ArticleController.HandlerGetPublishedArticles)
		articleFrontGroup.GET("/:id", ar.ArticleController.HandlerGetArticleDetail)
		articleFrontGroup.POST("/:id/like", ar.ArticleController.HandlerLikeArticle)
	}

}
