package models

type User struct {
	ID       string `json:"id"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Cash     uint   `json:"cash"`
	UserType string `json:"usertype"`
}

type CreateUser struct {
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Cash     uint   `json:"cash"`
	UserType string `json:"usertype"`
}

type UpdateUser struct { // but id will not be updated
	ID       string `json:"id"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Cash     uint   `json:"cash"`
}

type PrimaryKey struct {
	ID string `json:"id"`
}

type UsersResponse struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}

type GetAllRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
