package services

import (
	"gorm.io/gorm"
)

// ProjectService 项目服务
type ProjectService struct {
	db *gorm.DB
}

// NewProjectService 创建新的项目服务
func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{db: db}
}

// ChangeService 变更服务
type ChangeService struct {
	db *gorm.DB
}

// NewChangeService 创建新的变更服务
func NewChangeService(db *gorm.DB) *ChangeService {
	return &ChangeService{db: db}
}

// AnalysisService 分析服务
type AnalysisService struct {
	db *gorm.DB
}

// NewAnalysisService 创建新的分析服务
func NewAnalysisService(db *gorm.DB) *AnalysisService {
	return &AnalysisService{db: db}
}