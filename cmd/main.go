package main

import (
	v1 "MyBlogs/api/v1"
	"MyBlogs/internal/config"
	"MyBlogs/internal/databases"
	"MyBlogs/internal/repositories"
	"MyBlogs/internal/routers"
	"MyBlogs/internal/services"
	"MyBlogs/internal/utils"
	"fmt"
	"go.uber.org/zap"
	"log"
)

func main() {
	//初始化配置文件
	conf, err := config.NewConfig()
	if err != nil {
		log.Printf("NewConfig() failed,err:%v\n", err)
		return
	}
	fmt.Println("配置文件初始化成功...")

	//初始化日志
	utils.InitLogger(conf)

	//初始化数据库
	db, err := databases.InitDB(conf)
	if err != nil {
		utils.Error("数据库初始化失败", zap.Error(fmt.Errorf("InitDB() failed,err:%v", err)))
		return
	}
	fmt.Println("数据库初始化成功...")

	//依赖注入
	articleRepo := repositories.NewArticleRepo(db)
	articleService := services.NewArticleService(articleRepo)
	articleController := v1.NewArticleController(articleService)

	articleRouter := routers.NewArticleRouter(articleController)
	router := routers.NewMyRouter(articleRouter)
	r := router.SetUp()
	addr := "0.0.0.0:8080"
	err = r.Run(addr)
	if err != nil {
		utils.Error("服务器运行失败", zap.Error(fmt.Errorf("Run() failed,err:%v", err)))
		return
	}

}
