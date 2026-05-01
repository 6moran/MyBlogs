package v1

import (
	"MyBlogs/internal/models/dto/request"
	"MyBlogs/internal/services"
	"MyBlogs/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

const (
	MaxImageSize = 5 * 1024 * 1024
)

type ArticleController struct {
	AS services.ArticleService
}

func NewArticleController(service services.ArticleService) *ArticleController {
	return &ArticleController{
		AS: service,
	}
}

// 创建文章
func (artC *ArticleController) HandlerCreateArticle(c *gin.Context) {
	var car request.CreateArticleRequest
	err := c.ShouldBindJSON(&car)
	if err != nil {
		utils.Warn("JSON解析失败", zap.Error(fmt.Errorf("ShouldBindJSON() failed,err:%v", err)))
		c.JSON(400, gin.H{
			"msg": "参数错误",
		})
		return
	}
	err = artC.AS.CreateNewArticle(car)
	if err != nil {
		utils.Error("创建文章失败", zap.Error(fmt.Errorf("CreateNewArticle() failed,err:%v", err)))
		c.JSON(500, gin.H{
			"msg": "服务器出错",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "创建成功",
	})
}

func (artC *ArticleController) HandlerGetArticleLimit(c *gin.Context) {

}

// 删除文章
func (artC *ArticleController) HandlerDeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Warn("Atoi转换失败", zap.Error(fmt.Errorf("Atoi() failed,err:%v", err)))
		c.JSON(400, gin.H{
			"msg": "参数错误",
		})
		return
	}
	err = artC.AS.DeleteArticle(id)
	if err != nil {
		utils.Error("删除文章失败", zap.Error(fmt.Errorf("DeleteArticle() failed,err:%v", err)))
		c.JSON(500, gin.H{
			"msg": "服务器出错",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "删除成功",
	})

}

// 上传图片
func (artC *ArticleController) HandlerImage(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		utils.Error("上传图片失败", zap.Error(fmt.Errorf("FormFile() failed,err:%v", err)))
		c.JSON(400, gin.H{
			"msg": "参数错误",
		})
		return
	}

	//校验图片类型
	if !utils.IsImageFile(fileHeader.Filename) {
		utils.Warn("图片格式有误")
		c.JSON(400, gin.H{
			"msg": "只支持图片格式",
		})
		return
	}
	//校验图片大小
	if fileHeader.Size > MaxImageSize {
		utils.Warn("图片大小超出")
		c.JSON(400, gin.H{
			"msg": "图片不能超过5MB",
		})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			utils.Error("文件关闭失败", zap.Error(fmt.Errorf("file.Close() failed,err:%v", err)))
			return
		}
	}()

	//业务逻辑
	filePath, err := artC.AS.SaveImage(file, fileHeader)
	if err != nil {
		utils.Error("上传图片失败", zap.Error(fmt.Errorf("SaveImage() failed,err:%v", err)))
		c.JSON(500, gin.H{
			"msg": "系统错误",
		})
		return
	}

	//成功返回内部路径
	c.JSON(200, gin.H{
		"data": filePath,
	})
}
