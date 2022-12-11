package handler

import "github.com/labstack/echo/v4"

type CategoryHandlerInterface interface {
	FindAll(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	
	FindById(c echo.Context) error
}