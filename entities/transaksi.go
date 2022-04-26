package entities

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	Id       int    `json:"id"`
	UserID   uint   `json:"userID" form:"userID"`
	BarangID uint   `json:"barangID" form:"barangID"`
	Nama     string `json:"nama" form:"nama"`
	Barang   string `json:"barang" form:"barang"`
	Stok     int    `json:"stok" form:"stok"`
	Harga    int    `json:"harga" form:"harga"`
}
