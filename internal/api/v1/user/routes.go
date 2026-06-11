package user

import (
	"MyBlogs/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	UserController *Controller
}

func NewUserRouter(uc *Controller) *UserRouter {
	return &UserRouter{
		UserController: uc,
	}
}

func (ur *UserRouter) Register(front, admin *gin.RouterGroup) {
	// 前台用户路由
	userGroup := front.Group("/user")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/me", ur.UserController.HandlerGetUserInfo)
	}

	// 后台管理员路由
	userAdminGroup := admin.Group("/users")
	userAdminGroup.Use(middleware.AdminMiddleware())
	{
		userAdminGroup.GET("/", ur.UserController.HandlerGetUserList)
		userAdminGroup.GET("/:id", ur.UserController.HandlerGetUserByID)
		userAdminGroup.POST("/", ur.UserController.HandlerCreateUser)
		userAdminGroup.PUT("/:id", ur.UserController.HandlerUpdateUser)
		userAdminGroup.DELETE("/:id", ur.UserController.HandlerDeleteUser)
	}
}
