package comment

import (
	"MyBlogs/internal/middleware"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct {
	CommentController *Controller
}

func NewCommentRouter(cc *Controller) *CommentRouter {
	return &CommentRouter{
		CommentController: cc,
	}
}

func (cr *CommentRouter) Register(front, admin *gin.RouterGroup) {
	// 前台评论路由
	articleComments := front.Group("/articles/:id/comments")
	{
		// 获取评论（公开）
		articleComments.GET("", cr.CommentController.HandlerGetComments)

		// 发表评论（需要登录）
		articleComments.POST("", middleware.AuthMiddleware(), cr.CommentController.HandlerCreateComment)
	}

	// 后台评论管理
	adminComments := admin.Group("/comments")
	{
		adminComments.DELETE("/:id", cr.CommentController.HandlerDeleteComment)
	}
}
