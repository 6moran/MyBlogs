package response

import (
	bizerrors "MyBlogs/pkg/errors"
	"MyBlogs/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Body API 统一响应结构
type Body struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Success 写入成功响应
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Body{
		Code:    bizerrors.CodeSuccess,
		Message: "成功",
		Data:    data,
	})
}

// BadRequest 写入参数错误响应
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Body{
		Code:    bizerrors.CodeBadRequest,
		Message: message,
	})
}

// Unauthorized 写入未认证响应
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Body{
		Code:    bizerrors.CodeUnauthorized,
		Message: message,
	})
}

// Forbidden 写入无权限响应
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, Body{
		Code:    bizerrors.CodeForbidden,
		Message: message,
	})
}

// InternalServerError 写入服务器内部错误响应
func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Body{
		Code:    bizerrors.CodeServerError,
		Message: "服务器内部错误，请稍后再试",
	})
}

// BizError 写入业务错误响应
func BizError(c *gin.Context, err error) {
	e, ok := err.(*bizerrors.Error)
	if !ok {
		logger.Error("Error 断言失败，非业务错误", zap.Error(err))
		InternalServerError(c)
		return
	}
	c.JSON(http.StatusOK, Body{
		Code:    e.Code,
		Message: e.Message,
	})
}
