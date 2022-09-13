package usecase

import (
	"encoding/base64"

	"github.com/t0239184/GoChat/app/crypto"
	"github.com/t0239184/GoChat/app/domain"
)

type UserUsecase struct {
	UserRepository domain.IUserRepository
}

func NewUserUsecase(userRepository domain.IUserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

func (u *UserUsecase) CreateUser(user *domain.User) (id *int64, error error) {
	saltValue := crypto.GenerateSalt(64)
	iteration := 1000
	salt := &domain.Salt{
		Salt: base64.StdEncoding.EncodeToString(saltValue),
		Iteration: int16(iteration),
	}
	user.Password = crypto.HashWithSaltAndIteration(user.Password, saltValue, iteration)
	id, err := u.UserRepository.CreateUser(user, salt)
	if err != nil {
		return nil, err
	}
	return id, nil
}