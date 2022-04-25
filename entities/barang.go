package entities

import "gorm.io/gorm"

type Barang struct {
	gorm.Model
	Nama  string `json:"nama" form:"nama"`
	Jenis string `json:"jenis" form:"jenis"`
	Stok  int    `json:"stok" form:"stok"`
}
