package barang

import (
	_entities "project/entities"
)

type BarangRepositoryInterface interface {
	GetAll() ([]_entities.Barang, error)
	GetById(id int) (_entities.Barang, error)
	CreateBarang(barang _entities.Barang) error
	DeleteBarang(id, userID int) error
	UpdateBarang(id int, barang _entities.Barang) error
}
