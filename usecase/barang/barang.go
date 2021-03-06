package barang

import (
	_entities "project/entities"
	_barangRepository "project/repository/barang"
)

type BarangUseCase struct {
	barangRepository _barangRepository.BarangRepositoryInterface
}

func NewBarangUseCase(barangRepo _barangRepository.BarangRepositoryInterface) BarangUseCaseInterface {
	return &BarangUseCase{
		barangRepository: barangRepo,
	}
}

func (buc *BarangUseCase) GetAll() ([]_entities.Barang, error) {
	barangs, err := buc.barangRepository.GetAll()
	return barangs, err
}

func (buc *BarangUseCase) GetById(id int) (_entities.Barang, error) {
	barang, err := buc.barangRepository.GetById(id)
	return barang, err
}

func (buc *BarangUseCase) CreateBarang(barang _entities.Barang) error {
	err := buc.barangRepository.CreateBarang(barang)
	return err
}

func (buc *BarangUseCase) DeleteBarang(id, userID int) error {
	err := buc.barangRepository.DeleteBarang(id, userID)
	return err
}

func (buc *BarangUseCase) UpdateBarang(id int, barang _entities.Barang) error {
	err := buc.barangRepository.UpdateBarang(id, barang)
	return err
}
