package models

import "gorm.io/gorm"

type Project struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	DBName    string `gorm:"size:255"`
	CreatedAt string
}

func MigrateProjectSchema(db *gorm.DB) {
	db.AutoMigrate(&Project{})
}
