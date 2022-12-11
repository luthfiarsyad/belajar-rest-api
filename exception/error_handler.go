package exception

import (
	"belajar-rest-api/model/api"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	internalServerError(err, c)
}

func internalServerError(err error, c echo.Context) {
	c.JSON(http.StatusInternalServerError, api.ApiResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err.Error(),
	})
}

func PanicMiddlewares(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func(c echo.Context) error {
			err := recover()

			validationErrors, ok := err.(validator.ValidationErrors)

			if ok {
				return c.JSON(http.StatusBadRequest, api.ApiResponse{
					Code:   http.StatusBadRequest,
					Status: "BAD REQUEST",
					Data:   validationErrors.Error(),
				})
			}

			notFound, ok := err.(NotFoundError)

			if ok {
				return c.JSON(http.StatusNotFound, api.ApiResponse{
					Code:   http.StatusNotFound,
					Status: "NOT FOUND",
					Data:   notFound.Error.Error(),
				})
			}

			internalServerError, ok := err.(InternalServerError)

			if ok {
				return c.JSON(http.StatusNotFound, api.ApiResponse{
					Code:   http.StatusInternalServerError,
					Status: "INTERNAL SERVER ERROR",
					Data:   internalServerError.Error.Error(),
				})
			}
			return c.NoContent(http.StatusInternalServerError)
		}(c)
		return next(c)
	}
}
