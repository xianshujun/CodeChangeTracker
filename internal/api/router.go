package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"github.com/xianshujun/CodeChangeTracker/internal/api/handlers"
	"github.com/xianshujun/CodeChangeTracker/internal/api/middleware"
	"github.com/xianshujun/CodeChangeTracker/internal/config"
	"github.com/xianshujun/CodeChangeTracker/internal/services"
)

// SetupRouter 设置路由
func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.New()

	// 中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// 初始化服务
	projectService := services.NewProjectService(db)
	changeService := services.NewChangeService(db)
	analysisService := services.NewAnalysisService(db)

	// 初始化处理器
	projectHandler := handlers.NewProjectHandler(projectService)
	changeHandler := handlers.NewChangeHandler(changeService)
	analysisHandler := handlers.NewAnalysisHandler(analysisService)

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API路由组
	v1 := router.Group("/api/v1")
	{
		// 项目相关路由
		projects := v1.Group("/projects")
		{
			projects.POST("", projectHandler.CreateProject)
			projects.GET("", projectHandler.GetProjects)
			projects.GET("/:id", projectHandler.GetProject)
			projects.PUT("/:id", projectHandler.UpdateProject)
			projects.DELETE("/:id", projectHandler.DeleteProject)
			projects.POST("/:id/analyze", projectHandler.AnalyzeProject)
			projects.GET("/:id/changes", projectHandler.GetProjectChanges)
			projects.GET("/:id/risk-trends", projectHandler.GetRiskTrends)
		}

		// 变更相关路由
		changes := v1.Group("/changes")
		{
			changes.POST("/analyze", changeHandler.AnalyzeChange)
			changes.GET("/:id", changeHandler.GetChange)
			changes.GET("/:id/assessment", changeHandler.GetAssessment)
		}

		// 分析相关路由
		analysis := v1.Group("/analysis")
		{
			analysis.GET("/reports/:project_id", analysisHandler.GenerateReport)
			analysis.GET("/metrics/:project_id", analysisHandler.GetMetrics)
		}
	}

	// Swagger文档
	if cfg.Environment != "production" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}