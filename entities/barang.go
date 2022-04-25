package entities

import (
	"time"

	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	Nama         string    `json:"nama" form:"nama"`
	Jenis        string    `json:"jenis" form:"jenis"`
	Stok         string    `json:"stok" form:"stok"`
	TanggalInput time.Time `json:"tanggalinput" form:"tanggalinput"`
}
