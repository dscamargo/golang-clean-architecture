package repository

import (
	"github.com/dscamargo/crud-clean-architecture/src/domain"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type inmem struct {
	users []domain.User
}

func NewInMemRepo() *inmem {
	return &inmem{
		users: []domain.User{},
	}
}

func (m *inmem) FindById(id string) (domain.User, bool, error) {
	user := domain.User{}
	found := false
	for _, item := range m.users {
		if strconv.FormatUint(uint64(item.ID), 10) == id {
			user = item
			found = true
		}
	}

	if !found {
		return user, false, domain.ErrNotFound
	}

	return user, found, nil
}

func (m *inmem) FindByEmail(email string) (domain.User, bool, error) {
	user := domain.User{}
	found := false
	for _, item := range m.users {
		if item.Email == email {
			user = item
			found = true
		}
	}
	return user, found, nil
}

func (m *inmem) Create(name, email, password string) (string, error) {
	user := domain.User{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
		Name:      name,
		Email:     email,
		Password:  password,
	}
	m.users = append(m.users, user)
	return strconv.FormatUint(uint64(user.ID), 10), nil
}
