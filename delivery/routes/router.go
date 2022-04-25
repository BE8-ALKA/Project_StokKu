package routes

import (
	_authHandler "project/delivery/handler/auth"
	_barangHandler "project/delivery/handler/barang"
	_transaksiHandler "project/delivery/handler/transaksi"
	_userHandler "project/delivery/handler/user"

	_middlewares "project/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler, bh *_barangHandler.BarangHandler, th *_transaksiHandler.TransaksiHandler) {
	e.GET("/users", uh.GetAllHandler())
	e.GET("/users/:id", uh.GetByIdHandler(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUser())
	e.PUT("/users/:id", uh.UpdateUser(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUser(), _middlewares.JWTMiddleware())

	e.GET("/barangs", bh.GetAllHandler())
	e.GET("/barangs/:id", bh.GetByIdHandler(), _middlewares.JWTMiddleware())
	e.POST("/barangs", bh.CreateBarang(), _middlewares.JWTMiddleware())
	e.PUT("/barangs/:id", bh.UpdateBarang(), _middlewares.JWTMiddleware())
	e.DELETE("/barangs/:id", bh.DeleteBarang(), _middlewares.JWTMiddleware())

	e.GET("/transaksis", th.GetAllHandler())
	e.GET("/transaksis/:id", th.GetByIdHandler(), _middlewares.JWTMiddleware())
	e.POST("/transaksis", th.CreateTransaksi(), _middlewares.JWTMiddleware())
}
