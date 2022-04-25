package routes

import (
	_authHandler "project/delivery/handler/auth"
	// _bookHandler "book/delivery/handler/book"
	_userHandler "project/delivery/handler/user"

	_middlewares "project/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.GET("/users", uh.GetAllHandler())
	e.GET("/users/:id", uh.GetByIdHandler(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUser())
	e.PUT("/users/:id", uh.UpdateUser(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUser(), _middlewares.JWTMiddleware())

	// e.GET("/books", bh.GetAllHandler())
	// e.GET("/books/:id", bh.GetByIdHandler())
	// e.POST("/books", bh.CreateBook(), _middlewares.JWTMiddleware())
	// e.PUT("/books/:id", bh.UpdateBook(), _middlewares.JWTMiddleware())
	// e.DELETE("/books/:id", bh.DeleteBook(), _middlewares.JWTMiddleware())
}
