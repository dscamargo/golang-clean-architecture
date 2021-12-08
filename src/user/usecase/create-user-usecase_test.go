package usecase

import (
	"github.com/dscamargo/crud-clean-architecture/src/adapters"
	"github.com/dscamargo/crud-clean-architecture/src/domain"
	"github.com/dscamargo/crud-clean-architecture/src/user/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func makeSut() UserUsecase {
	repo := repository.NewInMemRepo()
	hasher := adapters.NewHasher()
	return NewUserUsecase(repo, hasher)
}

func makeFakeUser() map[string]string {
	return map[string]string{
		"name":                 "any-name",
		"email":                "any-email",
		"password":             "any-password",
		"passwordConfirmation": "any-password",
	}
}

func TestUserUsecase_CreateUser_RequiredFields(t *testing.T) {
	sut := makeSut()
	inputs := []map[string]string{
		{
			"email":                "any-email",
			"password":             "any-password",
			"passwordConfirmation": "any-password",
			"emptyField":           "name",
		},
		{
			"name":                 "any-name",
			"password":             "any-password",
			"passwordConfirmation": "any-password",
			"emptyField":           "email",
		},
		{
			"name":                 "any-name",
			"email":                "any-email",
			"passwordConfirmation": "any-password",
			"emptyField":           "password",
		},
		{
			"name":       "any-name",
			"email":      "any-email",
			"password":   "any-password",
			"emptyField": "passwordConfirmation",
		},
	}
	for _, input := range inputs {
		_, err := sut.CreateUser(input["name"], input["email"], input["password"], input["passwordConfirmation"])
		assert.NotNil(t, err)
		assert.Equal(t, err, domain.ErrMissingParam(input["emptyField"]))
	}
}

func TestUserUsecase_CreateUser_CheckPasswordConfirmation(t *testing.T) {
	input := map[string]string{
		"name":                 "any-name",
		"email":                "any-email",
		"password":             "any-password",
		"passwordConfirmation": "other-password",
	}
	sut := makeSut()
	_, err := sut.CreateUser(input["name"], input["email"], input["password"], input["passwordConfirmation"])
	assert.NotNil(t, err)
	assert.Equal(t, err, domain.ErrInvalidParam("passwordConfirmation"))

}

func TestUserUsecase_CreateUser_Success(t *testing.T) {
	input := makeFakeUser()
	sut := makeSut()
	userId, err := sut.CreateUser(input["name"], input["email"], input["password"], input["passwordConfirmation"])
	assert.Nil(t, err)
	assert.NotEmpty(t, userId)
}
