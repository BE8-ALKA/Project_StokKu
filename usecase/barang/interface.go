package barang

import (
	_entities "project/entities"
)

type BarangUseCaseInterface interface {
	GetAll() ([]_entities.Barang, error)
	GetById(id int) (_entities.Barang, int, error)
	CreateBarang(barang _entities.Barang) error
	DeleteBarang(id int) error
	UpdateBarang(id int, barang _entities.Barang) error
}
