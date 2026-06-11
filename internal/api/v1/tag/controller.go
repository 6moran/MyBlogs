package tag

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
	TagService service.TagService
}

func NewTagController(s service.TagService) *Controller {
	return &Controller{TagService: s}
}

// HandlerGetTags 获取所有标签
func (tc *Controller) HandlerGetTags(c *gin.Context) {
	tags, err := tc.TagService.GetAllTags()
	if err != nil {
		logger.Error("获取标签失败", zap.Error(err))
		response.InternalServerError(c)
		return
	}
	response.Success(c, tags)
}

// HandlerCreateTag 创建标签
func (tc *Controller) HandlerCreateTag(c *gin.Context) {
	var req request.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("参数解析失败", zap.Error(err))
		response.BadRequest(c, "参数错误")
		return
	}

	if err := tc.TagService.CreateTag(req); err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("创建标签失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("创建标签失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, nil)
}

// HandlerUpdateTag 更新标签
func (tc *Controller) HandlerUpdateTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	var req request.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("参数解析失败", zap.Error(err))
		response.BadRequest(c, "参数错误")
		return
	}

	if err := tc.TagService.UpdateTag(id, req); err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("更新标签失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("更新标签失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, nil)
}

// HandlerDeleteTag 删除标签
func (tc *Controller) HandlerDeleteTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	if err := tc.TagService.DeleteTag(id); err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("删除标签失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("删除标签失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, nil)
}

// HandlerGetTagByID 根据ID获取标签
func (tc *Controller) HandlerGetTagByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	resp, err := tc.TagService.GetTagByID(id)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("获取标签失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("获取标签失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, resp)
}

// HandlerGetTagList 获取标签列表（分页）
func (tc *Controller) HandlerGetTagList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	if page < 1 {
		page = 1
	}
	if size < 1 || size > 50 {
		size = 10
	}

	resp, total, err := tc.TagService.GetTagList(page, size)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("获取标签列表失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("获取标签列表失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}

	response.Success(c, gin.H{
		"data":  resp,
		"total": total,
	})
}
