package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"MyBlogs/internal/api"
	"MyBlogs/internal/repository"
	"MyBlogs/internal/service"
	"MyBlogs/pkg/config"
	"MyBlogs/pkg/databases"
	"MyBlogs/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// App 应用容器，集中管理基础设施、依赖装配、路由和 HTTP 服务生命周期
type App struct {
	mysqlDB *gorm.DB
	router  *gin.Engine
	cfg     *config.Config
	server  *http.Server
}

// New 创建应用容器
func New() *App {
	return &App{}
}

// Initialize 初始化整个应用
func (a *App) Initialize() error {
	if err := a.initConfig(); err != nil {
		return err
	}
	a.initLogger()
	a.initGin()
	if err := a.initMySQL(); err != nil {
		return err
	}
	a.initDependencies()
	a.initServe()
	return nil
}

// initConfig 加载配置文件
func (a *App) initConfig() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("加载配置失败: %w", err)
	}
	a.cfg = cfg
	return nil
}

// initLogger 初始化日志
func (a *App) initLogger() {
	logger.InitLogger(&a.cfg.Log)
	logger.Info("=========================================")
	logger.Info("配置加载成功")
	logger.Info("=========================================")
}

// initGin 设置 gin 模式
func (a *App) initGin() {
	gin.SetMode(gin.DebugMode)
}

// initMySQL 初始化 MySQL
func (a *App) initMySQL() error {
	db, err := databases.InitDB(&a.cfg.Database.MySQL)
	if err != nil {
		return fmt.Errorf("初始化 MySQL 失败: %w", err)
	}
	a.mysqlDB = db
	logger.Info("数据库初始化完成")
	return nil
}

// initDependencies 集中装配 Repository、Service 和 Controller 依赖
func (a *App) initDependencies() {
	// Repository 层
	articleRepo := repository.NewArticleRepo(a.mysqlDB)
	userRepo := repository.NewUserRepo(a.mysqlDB)
	tagRepo := repository.NewTagRepo(a.mysqlDB)

	// Service 层
	articleService := service.NewArticleService(articleRepo)
	userService := service.NewUserService(userRepo)
	tagService := service.NewTagService(tagRepo)

	// 路由注册
	a.router = api.NewRouter(api.Dependencies{
		ArticleService: articleService,
		UserService:    userService,
		TagService:     tagService,
	})
}

// initServe 初始化 HTTP 服务器
func (a *App) initServe() {
	a.server = &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      a.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

// Run 启动 HTTP 服务并阻塞当前进程
func (a *App) Run() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("服务器启动失败")
			return
		}
	}()
	a.gracefulShutdown()
}

// gracefulShutdown 优雅退出
func (a *App) gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("正在关闭服务器...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		logger.Error("关闭服务器失败")
	}
	logger.Info("服务器已关闭")
}
