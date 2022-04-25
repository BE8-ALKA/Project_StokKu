package entities

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	Nama   string `json:"nama" form:"nama"`
	Barang string `json:"barang" form:"barang"`
	Stok   int    `json:"stok" form:"stok"`
	Harga  int    `json:"harga" form:"harga"`
}
