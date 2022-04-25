package transaksi

import (
	"net/http"
	"project/delivery/helper"
	_entities "project/entities"
	_transaksiUseCase "project/usecase/transaksi"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransaksiHandler struct {
	transaksiUseCase _transaksiUseCase.TransaksiUseCaseInterface
}

func NewTransaksiHandler(transaksiUseCase _transaksiUseCase.TransaksiUseCaseInterface) *TransaksiHandler {
	return &TransaksiHandler{
		transaksiUseCase: transaksiUseCase,
	}
}

func (th *TransaksiHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		transaksi, err := th.transaksiUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all transaksi", transaksi))
	}
}

func (th *TransaksiHandler) GetByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		transaksi, rows, err := th.transaksiUseCase.GetById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all transaksi", transaksi))
	}
}

func (th *TransaksiHandler) CreateTransaksi() echo.HandlerFunc {
	return func(c echo.Context) error {
		var transaksi _entities.Transaksi
		c.Bind(&transaksi)
		err := th.transaksiUseCase.CreateTransaksi(transaksi)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to create transaksi"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success Create transaksi"))
	}
}
