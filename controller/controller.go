package controller

import (
	"encoding/json"
	"net/http"

	"main.go/models"
	"main.go/storage"
)

type Controller struct {
	Storage storage.IStorage
}

func New(storage storage.IStorage) Controller {
	return Controller{
		Storage: storage,
	}
}
func handleResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	resp := models.Response{}

	switch code := statusCode; {
	case code < 400:
		resp.Description = "success"
	case code < 500:
		resp.Description = "bad request"
	default:
		resp.Description = "internal server error"
	}

	resp.StatusCode = statusCode
	resp.Data = data

	js, _ := json.Marshal(resp)

	w.WriteHeader(statusCode)
	w.Write(js)
}
