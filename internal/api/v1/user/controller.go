package user

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/service"
	bizerrors "MyBlogs/pkg/errors"
	"MyBlogs/pkg/logger"
	"MyBlogs/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Controller struct {
	UserService service.UserService
}

func NewUserController(s service.UserService) *Controller {
	return &Controller{UserService: s}
}

// HandlerGetUserInfo 获取当前用户信息
func (uc *Controller) HandlerGetUserInfo(c *gin.Context) {
	userID, _ := c.Get("user_id")

	resp, err := uc.UserService.GetUserInfo(userID.(int))
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("获取用户信息失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("获取用户信息失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, resp)
}

// HandlerGetUserByID 根据ID获取用户信息
func (uc *Controller) HandlerGetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	resp, err := uc.UserService.GetUserByID(id)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("获取用户信息失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("获取用户信息失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, resp)
}

// HandlerGetUserList 获取用户列表
func (uc *Controller) HandlerGetUserList(c *gin.Context) {
	var req request.QueryUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Warn("参数解析失败", zap.Error(err))
		response.BadRequest(c, "参数错误")
		return
	}

	resp, total, err := uc.UserService.GetUserList(req)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("获取用户列表失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("获取用户列表失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, gin.H{
		"data":  resp,
		"total": total,
	})
}

// HandlerCreateUser 创建用户
func (uc *Controller) HandlerCreateUser(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("参数解析失败", zap.Error(err))
		response.BadRequest(c, "参数错误")
		return
	}

	if err := uc.UserService.CreateUser(req); err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("创建用户失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("创建用户失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, nil)
}

// HandlerUpdateUser 更新用户
func (uc *Controller) HandlerUpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("参数解析失败", zap.Error(err))
		response.BadRequest(c, "参数错误")
		return
	}

	if err := uc.UserService.UpdateUser(id, req); err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("更新用户失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("更新用户失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, nil)
}

// HandlerDeleteUser 删除用户
func (uc *Controller) HandlerDeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	if err := uc.UserService.DeleteUser(id); err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("删除用户失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("删除用户失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, nil)
}
