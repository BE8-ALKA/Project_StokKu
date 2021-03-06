package barang

import (
	"net/http"
	"project/delivery/helper"
	"project/delivery/middlewares"
	_entities "project/entities"
	_barangUseCase "project/usecase/barang"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BarangHandler struct {
	barangUseCase _barangUseCase.BarangUseCaseInterface
}

func NewBarangHandler(barangUseCase _barangUseCase.BarangUseCaseInterface) *BarangHandler {
	return &BarangHandler{
		barangUseCase: barangUseCase,
	}
}

func (bh *BarangHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		barangs, err := bh.barangUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all barang", barangs))
	}
}

func (bh *BarangHandler) GetByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		barangs, err := bh.barangUseCase.GetById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get barang by id", barangs))
	}
}

func (bh *BarangHandler) CreateBarang() echo.HandlerFunc {
	return func(c echo.Context) error {
		var barang _entities.Barang
		c.Bind(&barang)
		id := middlewares.ExtractToken(c)
		barang.UserID = uint(id)
		err := bh.barangUseCase.CreateBarang(barang)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to create barang"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success Create barang"))
	}
}

func (bh *BarangHandler) DeleteBarang() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var userID = middlewares.ExtractToken(c)

		err := bh.barangUseCase.DeleteBarang(id, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete barang by id"))
	}
}

func (bh *BarangHandler) UpdateBarang() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var barangs _entities.Barang
		c.Bind(&barangs)

		err := bh.barangUseCase.UpdateBarang(id, barangs)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update barang by id", barangs))
	}
}
