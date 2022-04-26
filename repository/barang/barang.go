package barang

import (
	"fmt"
	_entities "project/entities"

	"gorm.io/gorm"
)

type BarangRepository struct {
	database *gorm.DB
}

func NewBarangRepository(db *gorm.DB) *BarangRepository {
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

func (br *BarangRepository) GetById(id int) (_entities.Barang, error) {
	var barangs _entities.Barang
	tx := br.database.Where("id = ?", id).First(&barangs)
	if tx.Error != nil {
		return barangs, tx.Error
	}
	return barangs, nil
}

func (br *BarangRepository) CreateBarang(barang _entities.Barang) error {

	tx := br.database.Save(&barang)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (br *BarangRepository) DeleteBarang(id, userID int) error {
	var barangs _entities.Barang
	err := br.database.Where("id =? and user_id = ?", id, userID).First(&barangs).Error
	if err != nil {
		return err
	}
	br.database.Delete(&barangs)
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
