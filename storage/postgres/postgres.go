package postgres

import (
	"database/sql"
	"fmt"

	"main.go/config"
	"main.go/storage"
)

type Storage struct {
	DB *sql.DB
}

func New(cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s sslmode=disable`, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Storage{}, err
	}

	return Storage{
		DB: db,
	}, nil

}
func (s Storage) Close() {
	s.DB.Close()
}
func (s Storage) User() storage.IUserStorage {
	newuser := NewUserRepo(s.DB)
	return newuser

}
func (s Storage) Category() storage.ICategory {
	newcategory := NewCategory(s.DB)
	return newcategory

}
func (s Storage) Products() storage.IProductsStorage {
	newproducts := NewProducts(s.DB)
	return newproducts

}
func (s Storage) Baskets() storage.IBasketsStorage {
	newbasket := NewBasket(s.DB)
	return newbasket

}
func (s Storage) BasketProducts() storage.IBasketProductStorage {
	newbasketproduct := NewBasketProduct(s.DB)
	return newbasketproduct
}
