package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xianshujun/CodeChangeTracker/internal/services"
)

// AnalysisHandler 分析处理器
type AnalysisHandler struct {
	service *services.AnalysisService
}

// NewAnalysisHandler 创建新的分析处理器
func NewAnalysisHandler(service *services.AnalysisService) *AnalysisHandler {
	return &AnalysisHandler{service: service}
}

// GenerateReport 生成分析报告
// @Summary 生成分析报告
// @Description 为项目生成详细的分析报告
// @Tags analysis
// @Produce json
// @Param project_id path int true "项目ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /analysis/reports/{project_id} [get]
func (h *AnalysisHandler) GenerateReport(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "待实现"})
}

// GetMetrics 获取项目度量指标
// @Summary 获取项目度量指标
// @Description 获取项目的各种度量指标
// @Tags analysis
// @Produce json
// @Param project_id path int true "项目ID"
// @Success 200 {object} models.AnalysisMetrics
// @Failure 404 {object} map[string]string
// @Router /analysis/metrics/{project_id} [get]
func (h *AnalysisHandler) GetMetrics(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "待实现"})
}