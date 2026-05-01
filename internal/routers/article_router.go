package routers

import (
	v1 "MyBlogs/api/v1"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
	ArticleController *v1.ArticleController
}

func NewArticleRouter(ac *v1.ArticleController) *ArticleRouter {
	return &ArticleRouter{
		ArticleController: ac,
	}
}

func (ar *ArticleRouter) Register(r *gin.RouterGroup) {
	articleGroup := r.Group("/articles")
	{
		articleGroup.POST("", ar.ArticleController.HandlerCreateArticle)
		articleGroup.DELETE(":id", ar.ArticleController.HandlerDeleteArticle)
		articleGroup.POST("images", ar.ArticleController.HandlerImage)
	}

}
