package comment

import (
	"MyBlogs/internal/service"
	"MyBlogs/pkg/logger"
	"MyBlogs/pkg/response"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Controller struct {
	CommentService service.CommentService
}

func NewCommentController(s service.CommentService) *Controller {
	return &Controller{CommentService: s}
}

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	Content  string `json:"content" binding:"required"`
	ParentID int    `json:"parent_id"`
}

// HandlerGetComments 获取文章评论
func (cc *Controller) HandlerGetComments(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 50 {
		size = 10
	}

	resp, err := cc.CommentService.GetCommentsByArticleID(articleID, page, size)
	if err != nil {
		logger.Error("获取评论失败", zap.Error(fmt.Errorf("获取文章评论失败: %v", err)))
		response.InternalServerError(c)
		return
	}
	response.Success(c, resp)
}

// HandlerCreateComment 发表评论
func (cc *Controller) HandlerCreateComment(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	userID, _ := c.Get("user_id")

	if err := cc.CommentService.CreateComment(articleID, userID.(int), req.Content, req.ParentID); err != nil {
		logger.Error("创建评论失败", zap.Error(fmt.Errorf("创建评论失败: %v", err)))
		response.InternalServerError(c)
		return
	}
	response.Success(c, nil)
}

// HandlerDeleteComment 删除评论（管理员）
func (cc *Controller) HandlerDeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	if err := cc.CommentService.DeleteComment(id); err != nil {
		logger.Error("删除评论失败", zap.Error(fmt.Errorf("删除评论失败: %v", err)))
		response.InternalServerError(c)
		return
	}
	response.Success(c, nil)
}
