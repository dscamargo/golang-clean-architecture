package repository

import (
	"github.com/dscamargo/crud-clean-architecture/src/domain"
	"gorm.io/gorm"
	"strconv"
)

type SqliteUserRepository struct {
	db *gorm.DB
}

func NewSQLiteUserRepository(database *gorm.DB) *SqliteUserRepository {
	return &SqliteUserRepository{database}
}

func (r *SqliteUserRepository) FindById(id string) (domain.User, bool, error) {
	user := domain.User{}
	result := r.db.First(&user, "id=?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, false, domain.ErrNotFound
		}
		return user, false, result.Error
	}
	return user, true, nil
}

func (r *SqliteUserRepository) FindByEmail(email string) (domain.User, bool, error) {
	user := domain.User{}
	result := r.db.First(&user, "email=?", email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, false, domain.ErrNotFound
		}
		return user, false, result.Error
	}
	return user, true, nil
}

func (r *SqliteUserRepository) Create(name, email, password string) (string, error) {
	user := domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	result := r.db.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return strconv.FormatUint(uint64(user.ID), 10), nil
}
