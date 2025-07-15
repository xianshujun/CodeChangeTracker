package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xianshujun/CodeChangeTracker/internal/services"
)

// ChangeHandler 变更处理器
type ChangeHandler struct {
	service *services.ChangeService
}

// NewChangeHandler 创建新的变更处理器
func NewChangeHandler(service *services.ChangeService) *ChangeHandler {
	return &ChangeHandler{service: service}
}

// AnalyzeChange 分析变更
// @Summary 分析代码变更
// @Description 分析单个代码变更的风险
// @Tags changes
// @Accept json
// @Produce json
// @Param change body map[string]interface{} true "变更信息"
// @Success 201 {object} models.RiskAssessment
// @Failure 400 {object} map[string]string
// @Router /changes/analyze [post]
func (h *ChangeHandler) AnalyzeChange(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "待实现"})
}

// GetChange 获取变更详情
// @Summary 获取变更详情
// @Description 根据ID获取代码变更详情
// @Tags changes
// @Produce json
// @Param id path int true "变更ID"
// @Success 200 {object} models.CodeChange
// @Failure 404 {object} map[string]string
// @Router /changes/{id} [get]
func (h *ChangeHandler) GetChange(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "待实现"})
}

// GetAssessment 获取风险评估结果
// @Summary 获取风险评估结果
// @Description 获取变更的风险评估结果
// @Tags changes
// @Produce json
// @Param id path int true "变更ID"
// @Success 200 {object} models.RiskAssessment
// @Failure 404 {object} map[string]string
// @Router /changes/{id}/assessment [get]
func (h *ChangeHandler) GetAssessment(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "待实现"})
}