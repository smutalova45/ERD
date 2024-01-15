package models

type Products struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"orginalprice"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"categoryid"`
}
type CreateProduct struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"orginalprice"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"categoryid"`
}
type UpdateProducts struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"orginalprice"`
	Quantity      int    `json:"quantity"`
}
type PKProducts struct {
	ID string `json:"id"`
}
type ProductsResponse struct {
	Products []Products `json:"products"`
	Count    int        `json:"count"`
}
type GetAllrequestProducts struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
