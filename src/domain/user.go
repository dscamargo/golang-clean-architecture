package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `gorm:"size:255;index:unique" json:"email"`
	Password  string         `json:"-"`
	Addresses []Address      `json:"addresses"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserRepository interface {
	FindById(id string) (User, bool, error)
	FindByEmail(email string) (User, bool, error)
	Create(name, email, password string) (string, error)
}
