package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) Products(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateProduct(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIdProducts(w, r)
		} else {
			c.GetListProducts(w, r)
		}
	case http.MethodPut:
		c.UpdateProducts(w, r)
	case http.MethodDelete:
		c.DeleteProducts(w, r)
	}

}

func (c Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	createproduct := models.CreateProduct{}
	if err := json.NewDecoder(r.Body).Decode(&createproduct); err != nil {
		fmt.Println("error while reading data to json from client", err.Error())
		handleResponse(w, 500, err)
		return
	}
	products, err := c.Storage.Products().Create(createproduct)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 201, products)

}
func (c Controller) GetByIdProducts(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values.Get("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	products, err := c.Storage.Products().GetBYID(models.PKProducts{ID: id.String()})
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	js, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, js)
}
func (c Controller) GetListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.Storage.Products().GetListRequest(models.GetAllrequestProducts{})
	if err != nil {
		fmt.Println("error while getting list of products", err.Error())
		handleResponse(w, 500, err)
		return
	}
	js, err := json.Marshal(products)
	if err != nil {
		fmt.Println("error while marsheling", err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, js)

}
func (c Controller) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	products := models.UpdateProducts{}
	if err := json.NewDecoder(r.Body).Decode(&products); err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	updatedproduct, err := c.Storage.Products().Update(products)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, updatedproduct)
}
func (c Controller) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values.Get("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	err = c.Storage.Products().Delete(models.PKProducts{ID: id.String()})
	if err != nil {
		fmt.Println("error while deleting products", err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, "deleted")

}
