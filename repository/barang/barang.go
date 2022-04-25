package barang

import (
	"fmt"
	_entities "project/entities"

	"gorm.io/gorm"
)

type BarangRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *BarangRepository {
	return &BarangRepository{
		database: db,
	}
}

func (br *BarangRepository) GetAll() ([]_entities.Barang, error) {
	var barangs []_entities.Barang
	tx := br.database.Find(&barangs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return barangs, nil
}

func (br *BarangRepository) GetById(id int) (_entities.Barang, int, error) {
	var barangs _entities.Barang
	tx := br.database.Find(&barangs, id)
	if tx.Error != nil {
		return barangs, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return barangs, 0, nil
	}
	return barangs, int(tx.RowsAffected), nil
}

func (br *BarangRepository) CreateBarang(barang _entities.Barang) error {

	tx := br.database.Create(&barang)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (br *BarangRepository) DeleteBarang(id int) error {
	var barangs _entities.Barang
	tx := br.database.Where("id = ?", id).Delete(&barangs)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (br *BarangRepository) UpdateBarang(id int, barang _entities.Barang) error {

	tx := br.database.Where("id = ?", id).Updates(&barang)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {

		return fmt.Errorf("%s", "error")
	}
	return nil

}
