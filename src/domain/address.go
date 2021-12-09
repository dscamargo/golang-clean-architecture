package domain

import (
	"gorm.io/gorm"
	"time"
)

type Address struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name" `
	Street    string         `gorm:"not null" json:"street"`
	Number    int            `gorm:"not null" json:"number"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
