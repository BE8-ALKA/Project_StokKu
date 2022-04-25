package transaksi

import (
	_entities "project/entities"
)

type TransaksiUseCaseInterface interface {
	GetAll() ([]_entities.Transaksi, error)
	GetById(id int) (_entities.Transaksi, int, error)
	CreateTransaksi(transaksi _entities.Transaksi) error
}
