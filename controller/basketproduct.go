package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) BasketProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateBasketProducts(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIdBasketProducts(w, r)
		} else {
			c.GetListBasketProducts(w, r)
		}
	case http.MethodPut:
		c.UpdateBasketProducts(w, r)
	case http.MethodDelete:
		c.DeleteBasketProducts(w, r)
	}
}
func (c Controller) CreateBasketProducts(w http.ResponseWriter, r *http.Request) {
	createbasketproducts := models.CreateBP{}
	if err := json.NewDecoder(r.Body).Decode(&createbasketproducts); err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	basketproduct, err := c.Storage.BasketProducts().Create(createbasketproducts)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return

	}
	handleResponse(w, 200, basketproduct)
}
func (c Controller) GetByIdBasketProducts(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values.Get("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}

	basketproduct, err := c.Storage.BasketProducts().GetById(models.PKBasketProducts{ID: id.String()})
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	js, err := json.Marshal(basketproduct)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, js)
}
func (c Controller) GetListBasketProducts(w http.ResponseWriter, r *http.Request) {
	basketproduct, err := c.Storage.BasketProducts().GetListRequest(models.GetAllRequestBP{})
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	js, err := json.Marshal(basketproduct)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, js)

}
func (c Controller) UpdateBasketProducts(w http.ResponseWriter, r *http.Request) {
	basketproduct := models.UpdateBP{}
	if err := json.NewDecoder(r.Body).Decode(&basketproduct); err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	updatedbasketproduct, err := c.Storage.BasketProducts().Update(basketproduct)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 201, updatedbasketproduct)

}
func (c Controller) DeleteBasketProducts(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values.Get("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println("error while parsing to uuid", err.Error())
		handleResponse(w, 500, err)
		return
	}
	err = c.Storage.BasketProducts().Delete(models.PKBasketProducts{ID: id.String()})
	if err != nil {
		fmt.Println("error while deleting basket products", err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, "deleted")
}
