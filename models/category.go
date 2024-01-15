package models

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type CreateCategory struct {
	Name string `json:"name"`
}
type PKCategory struct {
	ID string `json:"id"`
}
type CategoryResponse struct {
	Categories []Category `json:"category"`
	Count      int        `json:"count"`
}
type GetAllRequestCategory struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
