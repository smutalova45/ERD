package storage

import "main.go/models"

type IStorage interface { //hamma crudlani ozini ichida saqlidi
	Close()
	User() IUserStorage
	Category() ICategory
	Products() IProductsStorage
	Baskets() IBasketsStorage
	BasketProducts() IBasketProductStorage
}
type IUserStorage interface {
	Create(models.CreateUser) (models.User, error)
	GetByID(models.PrimaryKey) (models.User, error)
	GetListRequest(models.GetAllRequest) (models.UsersResponse, error)
	Update(models.UpdateUser) (models.User, error)
	Delete(models.PrimaryKey) error
}
type ICategory interface {
	Create(models.CreateCategory) (models.Category, error)
	GetByID(models.PKCategory) (models.Category, error)
	GetListRequest(models.GetAllRequestCategory) (models.CategoryResponse, error)
	Update(models.Category) (models.Category, error)
	Delete(models.PKCategory) error
}
type IProductsStorage interface {
	Create(models.CreateProduct) (models.Products, error)
	GetBYID(models.PKProducts) (models.Products, error)
	GetListRequest(models.GetAllrequestProducts) (models.ProductsResponse, error)
	Update(models.UpdateProducts) (models.Products, error)
	Delete(models.PKProducts) error
}
type IBasketsStorage interface {
	Create(models.CreateBaskets) (models.Baskets, error)
	GetById(models.PKBaskets) (models.Baskets, error)
	GetListRequest(models.GetALLRequestBasket) (models.BasketResponse, error)
	Update(models.UpdateBaskets) (models.Baskets, error)
	Delete(models.PKBaskets) error
}
type IBasketProductStorage interface {
	Create(models.CreateBP) (models.BasketProducts, error)
	GetById(models.PKBasketProducts) (models.BasketProducts, error)
	GetListRequest(models.GetAllRequestBP) (models.BasketProductsResponse, error)
	Update(models.UpdateBP) (models.BasketProducts, error)
	Delete(models.PKBasketProducts) error
}
