package service

import (
	"belajar-rest-api/helper"
	"belajar-rest-api/model/domain"
	"belajar-rest-api/model/request"
	"belajar-rest-api/model/response"
	"belajar-rest-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type productService struct {
	db                *sql.DB
	productRepository repository.ProductRepository
	validate          *validator.Validate
}

func NewProductService(db *sql.DB, productRepository repository.ProductRepository, validate *validator.Validate) *productService {
	return &productService{db, productRepository, validate}
}

func (service *productService) FindAll(ctx context.Context) []response.ResponseProduct {
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.productRepository.FindAll(ctx, tx)

	responseProducts := []response.ResponseProduct{}
	for _, v := range products {
		responseProducts = append(responseProducts, v.ToResponseProduct())
	}
	return responseProducts
}

func (service *productService) Create(ctx context.Context, request request.RequestCreateProduct) response.ResponseProduct {
	service.validate.Struct(request)

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{}
	product.SetCategoryId(&request.CategoryId)
	product.SetName(&request.Name)
	product = service.productRepository.Create(ctx, tx, product)

	return product.ToResponseProduct()
}
func (service *productService) Update(ctx context.Context, request request.RequestUpdateProduct) {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := service.productRepository.FindById(ctx, tx, request.Id)
	product.SetCategoryId(&request.CategoryId)
	product.SetName(&request.Name)
	service.productRepository.Update(ctx, tx, product)
}

func (service *productService) Delete(ctx context.Context, request request.RequestDeleteProduct) {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := service.productRepository.FindById(ctx, tx, request.Id)
	service.productRepository.Delete(ctx, tx, product)
}

func (service *productService) FindById(ctx context.Context, id int) response.ResponseProduct {
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product := service.productRepository.FindById(ctx, tx, id)
	res := product.ToResponseProduct()
	return res
}

func (service *productService) FindByCategoryId(ctx context.Context, category_id int) []response.ResponseProduct {
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	responseProducts := []response.ResponseProduct{}
	products := service.productRepository.FindByCategoryId(ctx, tx, category_id)
	for _, v := range products {
		responseProducts = append(responseProducts, v.ToResponseProduct())
	}
	return responseProducts
}
