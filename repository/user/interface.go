package user

import (
	_entities "project/entities"
)

type UserRepositoryInterface interface {
	GetAll() ([]_entities.User, error)
	GetById(id int) (_entities.User, int, error)
	CreateUser(user _entities.User) error
	DeleteUser(id int) error
	UpdateUser(id int, user _entities.User) error
}
