package repository

import (
	"github.com/dscamargo/crud-clean-architecture/src/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		if item.ID == id {
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
		ID:       primitive.NewObjectID().String(),
		Name:     name,
		Email:    email,
		Password: password,
	}
	m.users = append(m.users, user)
	return user.ID, nil
}
