package transaksi

import (
	_entities "project/entities"

	"gorm.io/gorm"
)

type TransaksiRepository struct {
	database *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) *TransaksiRepository {
	return &TransaksiRepository{
		database: db,
	}
}

func (tr *TransaksiRepository) GetAll() ([]_entities.Transaksi, error) {
	var transaksi []_entities.Transaksi
	tx := tr.database.Find(&transaksi)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transaksi, nil
}

func (tr *TransaksiRepository) GetById(id int) (_entities.Transaksi, int, error) {
	var transaksi _entities.Transaksi
	tx := tr.database.Find(&transaksi, id)
	if tx.Error != nil {
		return transaksi, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return transaksi, 0, nil
	}
	return transaksi, int(tx.RowsAffected), nil
}

func (tr *TransaksiRepository) CreateTransaksi(transaksi _entities.Transaksi) error {

	tx := tr.database.Create(&transaksi)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}
