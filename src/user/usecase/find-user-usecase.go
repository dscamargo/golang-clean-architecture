package usecase

import (
	"github.com/dscamargo/crud-clean-architecture/src/domain"
)

type GetUserUsecase struct {
	userRepository domain.UserRepository
}

func NewGetUserUsecase(useRepo domain.UserRepository) GetUserUsecase {
	return GetUserUsecase{
		userRepository: useRepo,
	}
}

func (u *GetUserUsecase) GetUserById(id string) (domain.User, error) {
	user, _, err := u.userRepository.FindById(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
