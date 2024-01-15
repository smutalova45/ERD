package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) Categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateCategory(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIDCategory(w, r)
		} else {
			c.GetListCategory(w, r)
		}
	case http.MethodPut:
		c.UpdateCategory(w, r)
	case http.MethodDelete:
		c.DeleteCategory(w, r)
	}

}
func (c Controller) CreateCategory(w http.ResponseWriter, r *http.Request) {
	createcategory := models.CreateCategory{}
	if err := json.NewDecoder(r.Body).Decode(&createcategory); err != nil {
		fmt.Println("error while reading data from client", err.Error())
		handleResponse(w, http.StatusBadRequest, err)
		return
	}
	user, err := c.Storage.Category().Create(createcategory)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, 200, user)

}
func (c Controller) GetByIDCategory(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()
	idstr := values.Get("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	category, err := c.Storage.Category().GetByID(models.PKCategory{ID: id.String()})
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	js, err := json.Marshal(category)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, js)
}
func (c Controller) GetListCategory(w http.ResponseWriter, r *http.Request) {
	category, err := c.Storage.Category().GetListRequest(models.GetAllRequestCategory{})
	if err != nil {
		fmt.Println("error while getting list of categories", err.Error())
		handleResponse(w, 500, err)
		return
	}
	js, err := json.Marshal(category)
	if err != nil {
		fmt.Println("error while marsheling category", err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, js)

}
func (c Controller) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	updatedcategory, err := c.Storage.Category().Update(category)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return

	}
	handleResponse(w, 200, updatedcategory.ID)

}
func (c Controller) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values.Get("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	err = c.Storage.Category().Delete(models.PKCategory{ID: id.String()})
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, "deleted")

}
