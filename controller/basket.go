package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) Baskets(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		c.CreateBasket(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIDBasket(w, r)
		} else {
			c.GetALLRequestBasket(w, r)
		}
	case http.MethodPut:
		c.UpdateBasket(w, r)
	case http.MethodDelete:
		c.DeleteBasket(w, r)
	}
}
func (c Controller) CreateBasket(w http.ResponseWriter, r *http.Request) {
	createbasket := models.CreateBaskets{}
	if err := json.NewDecoder(r.Body).Decode(&createbasket); err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	basket, err := c.Storage.Baskets().Create(createbasket)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, basket)
}
func (c Controller) GetByIDBasket(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values.Get("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	basket, err := c.Storage.Baskets().GetById(models.PKBaskets{ID: id.String()})
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	js, err := json.Marshal(basket)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, js)

}
func (c Controller) GetALLRequestBasket(w http.ResponseWriter, r *http.Request) {
	basket, err := c.Storage.Baskets().GetListRequest(models.GetALLRequestBasket{})
	if err != nil {
		fmt.Println("error while getting list of baskets", err.Error())
		handleResponse(w, 500, err)
		return
	}
	json, err := json.Marshal(basket)
	if err != nil {
		fmt.Println("error while marsheling to json basket", err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, json)

}
func (c Controller) UpdateBasket(w http.ResponseWriter, r *http.Request) {
	baskets := models.UpdateBaskets{}
	if err := json.NewDecoder(r.Body).Decode(&baskets); err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	updatedbasket, err := c.Storage.Baskets().Update(baskets)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 201, updatedbasket.ID)

}
func (c Controller) DeleteBasket(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values.Get("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	err = c.Storage.Baskets().Delete(models.PKBaskets{ID: id.String()})
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, "deleted")

}
