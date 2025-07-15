package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/xianshujun/CodeChangeTracker/internal/models"
)

// Init 初始化数据库连接
func Init(databaseURL string) (*gorm.DB, error) {
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(databaseURL), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 自动迁移数据库表结构
	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database connected and migrated successfully")
	return db, nil
}

// migrate 执行数据库迁移
func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.CodeChange{},
		&models.RiskAssessment{},
		&models.AnalysisMetrics{},
	)
}

// GetDB 获取数据库实例（用于其他包）
func GetDB(db *gorm.DB) *gorm.DB {
	return db
}