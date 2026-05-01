package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyRouter struct {
	articleRouter *ArticleRouter
}

func NewMyRouter(AR *ArticleRouter) *MyRouter {
	return &MyRouter{
		articleRouter: AR,
	}
}

func (r MyRouter) SetUp() *gin.Engine {
	engine := gin.Default()

	engine.StaticFS("/uploads/images", http.Dir("./uploads/images"))

	////配置CORS跨域
	//engine.Use()

	apiGroup := engine.Group("/api")

	r.articleRouter.Register(apiGroup)

	return engine
}
