package usecase

import (
	"github.com/dscamargo/crud-clean-architecture/src/domain"
	"github.com/dscamargo/crud-clean-architecture/src/user/repository"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func makeGetUserSut() (GetUserUsecase, domain.UserRepository) {
	repo := repository.NewInMemRepo()
	sut := NewGetUserUsecase(repo)
	return sut, repo
}

func TestUserUsecase_GetUserById_NotFound(t *testing.T) {
	sut, _ := makeGetUserSut()
	_, err := sut.GetUserById("not-found-id")
	assert.Equal(t, err, domain.ErrNotFound)
}

func TestUserUsecase_GetUserById_Success(t *testing.T) {
	sut, repo := makeGetUserSut()
	userId, err := repo.Create("any-name", "any-email", "any-password")
	assert.Nil(t, err)
	user, err := sut.GetUserById(userId)
	assert.Nil(t, err)
	assert.Equal(t, strconv.FormatUint(uint64(user.ID), 10), userId)
	assert.Equal(t, user.Name, "any-name")
	assert.Equal(t, user.Email, "any-email")
	assert.Equal(t, user.Password, "any-password")
}
