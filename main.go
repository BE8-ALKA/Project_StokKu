package main

import (
	"project/config"
	_authHandler "project/delivery/handler/auth"
	_userHandler "project/delivery/handler/user"
	_authRepository "project/repository/auth"
	_userRepository "project/repository/user"
	_authUseCase "project/usecase/auth"
	_userUseCase "project/usecase/user"

	_barangHandler "project/delivery/handler/barang"
	_middlewares "project/delivery/middlewares"
	_barangRepository "project/repository/barang"
	_barangUseCase "project/usecase/barang"

	_transaksiHandler "project/delivery/handler/transaksi"

	_transaksiRepository "project/repository/transaksi"
	_transaksiUseCase "project/usecase/transaksi"

	"fmt"
	"log"

	_routes "project/delivery/routes"
	_utils "project/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	barangRepo := _barangRepository.NewBarangRepository(db)
	barangUseCase := _barangUseCase.NewBarangUseCase(barangRepo)
	barangHandler := _barangHandler.NewBarangHandler(barangUseCase)

	transaksiRepo := _transaksiRepository.NewTransaksiRepository(db)
	transaksiUseCase := _transaksiUseCase.NewTransaksiUseCase(transaksiRepo)
	transaksiHandler := _transaksiHandler.NewTransaksiHandler(transaksiUseCase)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())

	_routes.RegisterPath(e, userHandler, barangHandler, transaksiHandler)
	_routes.RegisterAuthPath(e, authHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
