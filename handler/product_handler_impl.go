package handler

import (
	"belajar-rest-api/helper"
	"belajar-rest-api/model/api"
	"belajar-rest-api/model/request"
	"belajar-rest-api/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{productService}
}

func (handler *productHandler) FindAll(c echo.Context) error {
	responseProduct := handler.productService.FindAll(c.Request().Context())

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseProduct,
	})
}

func (handler *productHandler) Create(c echo.Context) error {
	requestProduct := new(request.RequestCreateProduct)
	err := c.Bind(requestProduct)
	helper.PanicIfError(err)

	responseProduct := handler.productService.Create(c.Request().Context(), *requestProduct)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseProduct,
	})
}

func (handler *productHandler) Update(c echo.Context) error {
	requestProduct := new(request.RequestUpdateProduct)
	err := c.Bind(requestProduct)
	helper.PanicIfError(err)

	handler.productService.Update(c.Request().Context(), *requestProduct)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
	})
}
func (handler *productHandler) Delete(c echo.Context) error {
	requestProduct := new(request.RequestDeleteProduct)
	err := c.Bind(requestProduct)
	helper.PanicIfError(err)

	handler.productService.Delete(c.Request().Context(), *requestProduct)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
	})
}

func (handler *productHandler) FindById(c echo.Context) error {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	helper.PanicIfError(err)

	responseProduct := handler.productService.FindById(c.Request().Context(), id)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseProduct,
	})
}

func (handler *productHandler) FindByCategoryId(c echo.Context) error {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	helper.PanicIfError(err)

	responseProduct := handler.productService.FindByCategoryId(c.Request().Context(), id)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseProduct,
	})
}
