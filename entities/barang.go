package entities

import "gorm.io/gorm"

type Barang struct {
	gorm.Model
	UserID     string `json:"UserID" form:"UserID"`
	Nama       string `json:"nama" form:"nama"`
	Jenis      string `json:"jenis" form:"jenis"`
	Stok       int    `json:"stok" form:"stok"`
	Transaksis []Transaksi
}
