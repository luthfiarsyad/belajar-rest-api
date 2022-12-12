package main

import (
	"belajar-rest-api/database"
	"belajar-rest-api/exception"
	"belajar-rest-api/handler"
	"belajar-rest-api/middleware"
	"belajar-rest-api/repository"
	"belajar-rest-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	db := database.GetConnection()
	validate := validator.New()

	// Category Handler Initialization
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(db, categoryRepository, validate)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Product Handler Initialization
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(db, productRepository, validate)
	productHandler := handler.NewProductHandler(productService)

	echo := echo.New()
	echo.Use(exception.PanicMiddlewares, middleware.AuthMiddleware)

	// API Grouping Echo
	api := echo.Group("/api")

	// Categories
	api.GET("/categories", categoryHandler.FindAll)
	api.POST("/categories", categoryHandler.Create)
	api.PUT("/categories", categoryHandler.Update)
	api.DELETE("/categories", categoryHandler.Delete)
	api.GET("/categories/:id", categoryHandler.FindById)

	// Products
	api.GET("/products", productHandler.FindAll)
	api.POST("/products", productHandler.Create)
	api.PUT("/products", productHandler.Update)
	api.DELETE("/products", productHandler.Delete)
	api.GET("/products/:id", productHandler.FindById)
	api.GET("/category/:category_id/products", productHandler.FindByCategoryId)

	// Echo run server on port :3000
	echo.Start(":3000")
}
