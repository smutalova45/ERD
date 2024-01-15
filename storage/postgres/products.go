package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"main.go/models"
	"main.go/storage"
)

type productsRepo struct {
	db *sql.DB
}

func NewProducts(db *sql.DB) storage.IProductsStorage {
	return productsRepo{
		db: db,
	}

}
func (p productsRepo) Create(prpduct models.CreateProduct) (models.Products, error) {
	id := uuid.New()
	if _, err := p.db.Exec(`insert into products values($1,$2,$3,$4,$5)`, id, prpduct.Name, prpduct.Price, prpduct.OriginalPrice, prpduct.Quantity, prpduct.CategoryID); err != nil {
		return models.Products{}, err
	}
	return models.Products{}, nil

}
func (p productsRepo) GetBYID(id models.PKProducts) (models.Products, error) {
	products := models.Products{}
	if err := p.db.QueryRow(`select id, name_,price, originalprice, quantity, categoryid where id=$1`, id.ID).Scan(
		&products.ID,
		&products.Name,
		&products.Price,
		&products.OriginalPrice,
		&products.Quantity,
	); err != nil {
		return models.Products{}, err
	}
	return models.Products{}, nil
}
func (p productsRepo) GetListRequest(models.GetAllrequestProducts) (models.ProductsResponse, error) {
	rows, err := p.db.Query(`select * from products`)
	if err != nil {
		return models.ProductsResponse{}, err
	}
	products := []models.Products{}
	for rows.Next() {
		product := models.Products{}
		if err != nil {
			return models.ProductsResponse{}, err
		}
		products = append(products, product)
	}
	response := models.ProductsResponse{
		Products: products,
		Count:    len(products),
	}
	return response, nil
}
func (p productsRepo) Update(products models.UpdateProducts) (models.Products, error) {

	_, err := p.db.Exec(`update products set name_=$1, price=$2,originalprice=$3, quantity=$4 where id=$5`, products.Name, products.Price, products.OriginalPrice, products.Quantity, products.ID)
	if err != nil {
		return models.Products{}, err
	}
	updatedproduct, err := p.GetBYID(models.PKProducts{ID: products.ID})
	if err != nil {
		return models.Products{}, err
	}
	return updatedproduct, nil

}
func (p productsRepo) Delete(products models.PKProducts) error {
	if _, err := p.db.Exec(`delete from basketproducts where id=$1`, products.ID); err != nil {
		return err
	}
	if _, err := p.db.Exec(`delete from products where id=$1`, products.ID); err != nil {
		return err
	}
	return nil
}
