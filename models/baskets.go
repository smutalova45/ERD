package models

type Baskets struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerid"`
	TotalSum   int    `json:"totalsum"`
}
type CreateBaskets struct {
	CustomerID string `json:"customerid"`
	TotalSum   int    `json:"totalsum"`
}
type UpdateBaskets struct {
	ID       string `json:"id"`
	TotalSum int    `json:"totalsum"`
}

type PKBaskets struct {
	ID string `json:"id"`
}
type BasketResponse struct {
	Baskets []Baskets `json:"baskets"`
	Count   int       `json:"count"`
}
type GetALLRequestBasket struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
