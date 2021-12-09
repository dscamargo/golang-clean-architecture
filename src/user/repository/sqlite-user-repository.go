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
	if err := r.db.Preload("Addresses").First(&user, "users.id=?", id); err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return user, false, domain.ErrNotFound
		}
		return user, false, err.Error
	}
	return user, true, nil
}

func (r *SqliteUserRepository) FindByEmail(email string) (domain.User, bool, error) {
	user := domain.User{}
	if err := r.db.First(&user, "email = ?", email); err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return user, false, domain.ErrNotFound
		}
		return user, false, err.Error
	}
	return user, true, nil
}

func (r *SqliteUserRepository) Create(name, email, password string) (string, error) {
	user := domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	if err := r.db.Create(&user).Association("Addresses").Append([]domain.Address{
		{
			Name:   "House one",
			Street: "Street one",
			Number: 1,
		},
		{
			Name:   "House two",
			Street: "Street two",
			Number: 2,
		}}); err != nil {
		return "", err
	}
	return strconv.FormatUint(uint64(user.ID), 10), nil
}
