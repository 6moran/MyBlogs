package article

import (
	"MyBlogs/internal/model/dto/request"
	"MyBlogs/internal/service"
	bizerrors "MyBlogs/pkg/errors"
	"MyBlogs/pkg/logger"
	"MyBlogs/pkg/minio"
	"MyBlogs/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	MaxImageSize = 5 * 1024 * 1024
)

type Controller struct {
	AS service.ArticleService
}

func NewArticleController(s service.ArticleService) *Controller {
	return &Controller{AS: s}
}

// HandlerCreateArticle 创建文章
func (artC *Controller) HandlerCreateArticle(c *gin.Context) {
	var car request.CreateArticleRequest
	if err := c.ShouldBindJSON(&car); err != nil {
		logger.Warn("JSON解析失败", zap.Error(err))
		response.BadRequest(c, "参数错误")
		return
	}
	if err := artC.AS.CreateNewArticle(car); err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("创建文章失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("创建文章失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}
	response.Success(c, nil)
}

// HandlerGetArticleLimit 获取文章分页列表
func (artC *Controller) HandlerGetArticleLimit(c *gin.Context) {
	var qar request.QueryArticleRequest
	if err := c.ShouldBindQuery(&qar); err != nil {
		logger.Warn("参数解析失败", zap.Error(err))
		response.BadRequest(c, "参数错误")
		return
	}
	res, total, err := artC.AS.GetArticleLimit(qar)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("查询文章失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("查询文章失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}
	response.Success(c, gin.H{
		"data":  res,
		"total": total,
	})
}

// HandlerDeleteArticle 删除文章
func (artC *Controller) HandlerDeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Warn("参数错误", zap.Error(err))
		response.BadRequest(c, "参数错误")
		return
	}
	if err = artC.AS.DeleteArticle(id); err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("删除文章失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("删除文章失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}
	response.Success(c, nil)
}

// HandlerGetPublishedArticles 获取已发布的文章列表
func (artC *Controller) HandlerGetPublishedArticles(c *gin.Context) {
	var req request.PublishedArticleRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Warn("参数解析失败", zap.Error(err))
		response.BadRequest(c, "参数错误")
		return
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 || req.Size > 50 {
		req.Size = 10
	}

	resp, err := artC.AS.GetPublishedArticles(req.Page, req.Size, req.TagID, req.Keyword)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("查询文章失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("查询文章失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}
	response.Success(c, resp)
}

// HandlerGetArticleDetail 获取文章详情
func (artC *Controller) HandlerGetArticleDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	resp, err := artC.AS.GetArticleDetail(id)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("查询文章详情失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("查询文章详情失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}
	response.Success(c, resp)
}

// HandlerLikeArticle 点赞文章
func (artC *Controller) HandlerLikeArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "参数错误")
		return
	}

	likeCount, err := artC.AS.LikeArticle(id)
	if err != nil {
		logger.Error("点赞失败", zap.Error(err))
		response.InternalServerError(c)
		return
	}

	response.Success(c, gin.H{"like_count": likeCount})
}

// HandlerImage 上传图片
func (artC *Controller) HandlerImage(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.PostForm("article_id"))

	fileHeader, err := c.FormFile("image")
	if err != nil {
		logger.Warn("上传图片失败", zap.Error(err))
		response.BadRequest(c, "参数错误")
		return
	}

	if !minio.IsImageFile(fileHeader.Filename) {
		logger.Warn("图片格式有误")
		response.BadRequest(c, "只支持图片格式")
		return
	}
	if fileHeader.Size > MaxImageSize {
		logger.Warn("图片大小超出")
		response.BadRequest(c, "图片不能超过5MB")
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		response.InternalServerError(c)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Error("文件关闭失败", zap.Error(err))
		}
	}()

	filePath, err := artC.AS.SaveImage(file, fileHeader, articleID)
	if err != nil {
		if bizerrors.IsBizError(err) {
			logger.Warn("上传图片失败", zap.Error(err))
			response.BizError(c, err)
		} else {
			logger.Error("上传图片失败", zap.Error(err))
			response.InternalServerError(c)
		}
		return
	}
	response.Success(c, filePath)
}
