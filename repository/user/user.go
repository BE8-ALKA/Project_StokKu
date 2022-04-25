package user

import (
	"fmt"
	_entities "project/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) GetAll() ([]_entities.User, error) {
	var users []_entities.User
	tx := ur.database.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (ur *UserRepository) GetById(id int) (_entities.User, int, error) {
	var user _entities.User
	tx := ur.database.Find(&user, id)
	if tx.Error != nil {
		return user, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, 0, nil
	}
	return user, int(tx.RowsAffected), nil
}

func (ur *UserRepository) CreateUser(user _entities.User) error {

	tx := ur.database.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (ur *UserRepository) DeleteUser(id int) error {
	var users _entities.User
	tx := ur.database.Where("id = ?", id).Delete(&users)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (ur *UserRepository) UpdateUser(id int, user _entities.User) error {

	tx := ur.database.Where("id = ?", id).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {

		return fmt.Errorf("%s", "error")
	}
	return nil

}
