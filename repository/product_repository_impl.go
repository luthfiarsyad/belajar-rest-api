package repository

import (
	"belajar-rest-api/exception"
	"belajar-rest-api/helper"
	"belajar-rest-api/model/domain"
	"context"
	"database/sql"
)

type productRepository struct {
}

func NewProductRepository() *productRepository {
	return &productRepository{}
}

func (repo *productRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	query := "SELECT id, category_id, name FROM product"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		product := domain.Product{}
		rows.Scan(product.GetId(), product.GetCategoryId(), product.GetName())
		products = append(products, product)
	}
	return products
}
func (repo *productRepository) Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	query := "INSERT INTO product (id, category_id, name) VALUES (?,?,?)"

	res, err := tx.ExecContext(ctx, query, product.GetId(), product.GetCategoryId(), product.GetName())
	helper.PanicIfError(err)
	lastId, err := res.LastInsertId()
	helper.PanicIfError(err)
	last := int(lastId)
	product.SetId(&last)
	return product
}
func (repo *productRepository) Update(ctx context.Context, tx *sql.Tx, product domain.Product) {
	query := "UPDATE product SET category_id=?, name=? WHERE id=?"

	_, err := tx.ExecContext(ctx, query, product.GetCategoryId(), product.GetName(), product.GetId())
	helper.PanicIfError(err)
}

func (repo *productRepository) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	query := "DELETE FROM product WHERE id=?"

	_, err := tx.ExecContext(ctx, query, product.GetId())
	helper.PanicIfError(err)
}

func (repo *productRepository) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Product {
	product := domain.Product{}
	query := "SELECT id, category_id,name FROM product WHERE id=?"
	row := tx.QueryRowContext(ctx, query, id)

	err := row.Scan(product.GetId(), product.GetCategoryId(), product.GetName())
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}
	return product
}

func (repo *productRepository) FindByCategoryId(ctx context.Context, tx *sql.Tx, id int) []domain.Product {
	query := "SELECT id, name FROM product WHERE category_id =?"

	rows, err := tx.QueryContext(ctx, query, id)
	helper.PanicIfError(err)
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		product := domain.Product{}
		rows.Scan(product.GetId(), product.GetName())
		product.SetCategoryId(&id)
		products = append(products, product)
	}

	return products
}
