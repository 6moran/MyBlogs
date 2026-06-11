package tag

import (
	"MyBlogs/internal/middleware"

	"github.com/gin-gonic/gin"
)

type TagRouter struct {
	TagController *Controller
}

func NewTagRouter(tc *Controller) *TagRouter {
	return &TagRouter{TagController: tc}
}

func (tr *TagRouter) Register(front, admin *gin.RouterGroup) {
	// 前台标签路由（公开）
	frontGroup := front.Group("/tags")
	{
		frontGroup.GET("", tr.TagController.HandlerGetTags)
	}

	// 后台标签路由（需要管理员权限）
	adminGroup := admin.Group("/tags")
	adminGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminGroup.POST("/", tr.TagController.HandlerCreateTag)
		adminGroup.PUT("/:id", tr.TagController.HandlerUpdateTag)
		adminGroup.DELETE("/:id", tr.TagController.HandlerDeleteTag)
		adminGroup.GET("/", tr.TagController.HandlerGetTagList)
		adminGroup.GET("/:id", tr.TagController.HandlerGetTagByID)
	}
}
