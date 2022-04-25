package transaksi

import (
	_entities "project/entities"
	_transaksiRepository "project/repository/transaksi"
)

type TransaksiUseCase struct {
	transaksiRepository _transaksiRepository.TransaksiRepositoryInterface
}

func NewTransaksiUseCase(transaksiRepo _transaksiRepository.TransaksiRepositoryInterface) TransaksiUseCaseInterface {
	return &TransaksiUseCase{
		transaksiRepository: transaksiRepo,
	}
}

func (tuc *TransaksiUseCase) GetAll() ([]_entities.Transaksi, error) {
	transaksi, err := tuc.transaksiRepository.GetAll()
	return transaksi, err
}

func (tuc *TransaksiUseCase) GetById(id int) (_entities.Transaksi, int, error) {
	transaksi, rows, err := tuc.transaksiRepository.GetById(id)
	return transaksi, rows, err
}

func (tuc *TransaksiUseCase) CreateTransaksi(transaksi _entities.Transaksi) error {
	err := tuc.transaksiRepository.CreateTransaksi(transaksi)
	return err
}
