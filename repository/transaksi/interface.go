package transaksi

import (
	_entities "project/entities"
)

type TransaksiRepositoryInterface interface {
	GetAll() ([]_entities.Transaksi, error)
	GetById(id int) (_entities.Transaksi, int, error)
	CreateTransaksi(user _entities.Transaksi) error
}
