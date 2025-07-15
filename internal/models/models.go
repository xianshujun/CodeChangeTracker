package models

import (
	"time"

	"gorm.io/gorm"
)

// Project 项目模型
type Project struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;uniqueIndex"`
	Description string         `json:"description"`
	RepoURL     string         `json:"repo_url" gorm:"not null"`
	Language    string         `json:"language"`
	Branch      string         `json:"branch" gorm:"default:main"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关联关系
	Changes []CodeChange `json:"changes,omitempty" gorm:"foreignKey:ProjectID"`
}

// CodeChange 代码变更模型
type CodeChange struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	ProjectID    uint           `json:"project_id" gorm:"not null;index"`
	CommitHash   string         `json:"commit_hash" gorm:"not null;uniqueIndex"`
	Author       string         `json:"author" gorm:"not null"`
	AuthorEmail  string         `json:"author_email"`
	Message      string         `json:"message"`
	FilesChanged StringArray    `json:"files_changed" gorm:"type:text"`
	LinesAdded   int            `json:"lines_added"`
	LinesDeleted int            `json:"lines_deleted"`
	RiskScore    float64        `json:"risk_score" gorm:"index"`
	Status       string         `json:"status" gorm:"default:pending"` // pending, analyzed, failed
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关联关系
	Project     Project          `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
	Assessments []RiskAssessment `json:"assessments,omitempty" gorm:"foreignKey:ChangeID"`
}

// RiskAssessment 风险评估模型
type RiskAssessment struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ChangeID    uint           `json:"change_id" gorm:"not null;index"`
	RiskLevel   string         `json:"risk_level" gorm:"not null"` // LOW, MEDIUM, HIGH, CRITICAL
	Score       float64        `json:"score" gorm:"not null"`
	Factors     StringArray    `json:"factors" gorm:"type:text"`
	Suggestions StringArray    `json:"suggestions" gorm:"type:text"`
	Confidence  float64        `json:"confidence"`
	ModelVersion string        `json:"model_version"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关联关系
	Change CodeChange `json:"change,omitempty" gorm:"foreignKey:ChangeID"`
}

// AnalysisMetrics 分析度量模型
type AnalysisMetrics struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	ProjectID         uint      `json:"project_id" gorm:"not null;index"`
	Date              time.Time `json:"date" gorm:"not null;index"`
	TotalChanges      int       `json:"total_changes"`
	HighRiskChanges   int       `json:"high_risk_changes"`
	MediumRiskChanges int       `json:"medium_risk_changes"`
	LowRiskChanges    int       `json:"low_risk_changes"`
	AverageRiskScore  float64   `json:"average_risk_score"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// User 用户模型
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"not null;uniqueIndex"`
	Email     string         `json:"email" gorm:"not null;uniqueIndex"`
	Role      string         `json:"role" gorm:"default:user"` // admin, user
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}