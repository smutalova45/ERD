package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"main.go/models"
	"main.go/pkg/check"
)

func (c Controller) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateUser(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIDUser(w, r)
		} else {
			c.GetListUsers(w, r)
		}
	case http.MethodPut:
		c.UpdateUser(w, r)
	case http.MethodDelete:
		c.DeleteUser(w, r)
	}

}

func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := models.CreateUser{}
	if err := json.NewDecoder(r.Body).Decode(&CreateUser); err != nil {
		fmt.Println("error while reading data from client", err.Error())
		handleResponse(w, http.StatusBadRequest, err)
		return
	}
	if !check.Password(CreateUser.Password) {
		fmt.Println("the password format is not correct!")
		handleResponse(w, http.StatusBadRequest, errors.New("password should involves lower and upper case latters"))
		return
	}

	user, err := c.Storage.User().Create(CreateUser)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, 201, user)

}
func (c Controller) GetByIDUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values.Get("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	user, err := c.Storage.User().GetByID(models.PrimaryKey{ID: id.String()})
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
	}
	js, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error while marsheling user", err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, js)
}

func (c Controller) GetListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Storage.User().GetListRequest(models.GetAllRequest{})
	if err != nil {
		fmt.Println("error while getting list of users", err.Error())
		handleResponse(w, 500, err)
		return
	}
	userJS, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error while marsheking json", err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, userJS)

}
func (c Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := models.UpdateUser{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	updateduser, err := c.Storage.User().Update(user)
	if err != nil {
		fmt.Println(err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, updateduser.ID)
}
func (c Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()
	idstr := values.Get("id")
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println("error while parsing to uuid", err.Error())
		handleResponse(w, 500, err)
		return
	}
	err = c.Storage.User().Delete(models.PrimaryKey{ID: id.String()})
	if err != nil {
		fmt.Println("error while deleting user", err.Error())
		handleResponse(w, 500, err)
		return
	}
	handleResponse(w, 200, "deleted")
}
