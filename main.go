package main

import (
	"project/config"
	// _authHandler "book/delivery/handler/auth"
	_userHandler "project/delivery/handler/user"
	// _authRepository "book/repository/auth"
	_userRepository "project/repository/user"
	// _authUseCase "book/usecase/auth"
	_userUseCase "project/usecase/user"

	// _bookHandler "book/delivery/handler/book"
	// _middlewares "book/delivery/middlewares"
	// _bookRepository "book/repository/book"
	// _bookUseCase "book/usecase/book"

	"fmt"
	"log"

	_routes "project/delivery/routes"
	_utils "project/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	// bookRepo := _bookRepository.NewBookRepository(db)
	// bookUsecase := _bookUseCase.NewBookUseCase(bookRepo)
	// bookHandler := _bookHandler.NewBookHandler(bookUsecase)

	// authRepo := _authRepository.NewAuthRepository(db)
	// authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	// authHandler := _authHandler.NewAuthHandler(authUseCase)
	e := echo.New()
	// e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(_middlewares.CustomLogger())

	_routes.RegisterPath(e, userHandler)
	// _routes.RegisterAuthPath(e, authHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
