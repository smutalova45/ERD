package models

type BasketProducts struct {
	ID        string `json:"id"`
	BasketID  string `json:"basketid"`
	ProductID string `json:"productid"`
	Quantity  int    `json:"qunatity"`
}
type CreateBP struct {
	BasketID  string `json:"basketid"`
	ProductID string `json:"productid"`
	Quantity  int    `json:"qunatity"`
}
type UpdateBP struct {
	ID       string `json:"id"`
	Quantity int    `json:"qunatity"`
}
type PKBasketProducts struct {
	ID string `json:"id"`
}
type BasketProductsResponse struct {
	BasketProducts []BasketProducts `json:"basketproducts"`
	Count          int              `json:"count"`
}
type GetAllRequestBP struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
