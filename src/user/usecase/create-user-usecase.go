package usecase

import (
	"github.com/dscamargo/crud-clean-architecture/src/adapters"
	"github.com/dscamargo/crud-clean-architecture/src/domain"
)

type CreateUserUsecase struct {
	userRepository domain.UserRepository
	hasher         adapters.Hasher
}

func NewCreateUserUsecase(useRepo domain.UserRepository, hasher adapters.Hasher) CreateUserUsecase {
	return CreateUserUsecase{
		userRepository: useRepo,
		hasher:         hasher,
	}
}

func (u *CreateUserUsecase) CreateUser(name, email, password, passwordConfirmation string) (string, error) {
	input := map[string]string{
		"name":                 name,
		"email":                email,
		"password":             password,
		"passwordConfirmation": passwordConfirmation,
	}

	requiredFields := []string{"name", "email", "password", "passwordConfirmation"}
	for _, requiredField := range requiredFields {
		if input[requiredField] == "" {
			return "", domain.ErrMissingParam(requiredField)
		}
	}

	_, found, _ := u.userRepository.FindByEmail(email)
	if found {
		return "", domain.ErrConflict
	}
	if password != passwordConfirmation {
		return "", domain.ErrInvalidParam("passwordConfirmation")
	}
	hashedPassword := u.hasher.Hash(password)
	newUserId, err := u.userRepository.Create(name, email, hashedPassword)
	if err != nil {
		return "", err
	}
	return newUserId, nil
}
