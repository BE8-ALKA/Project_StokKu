package user

import (
	_entities "project/entities"
	_userRepository "project/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) GetAll() ([]_entities.User, error) {
	users, err := uuc.userRepository.GetAll()
	return users, err
}

func (uuc *UserUseCase) GetById(id int) (_entities.User, int, error) {
	user, rows, err := uuc.userRepository.GetById(id)
	return user, rows, err
}

func (uuc *UserUseCase) CreateUser(user _entities.User) error {
	err := uuc.userRepository.CreateUser(user)
	return err
}

func (uuc *UserUseCase) DeleteUser(id int) error {
	err := uuc.userRepository.DeleteUser(id)
	return err
}

func (uuc *UserUseCase) UpdateUser(id int, user _entities.User) error {
	err := uuc.userRepository.UpdateUser(id, user)
	return err
}
