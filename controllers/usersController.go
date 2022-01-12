package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pk-anderson/exercicio-crud-go/models"
)

func MainRoute(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Rota principal")
}

func GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	userList := models.SelectAllUsers()
	json.NewEncoder(rw).Encode(userList)
}

func GetUserById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := models.SelectUserById(id)
	json.NewEncoder(rw).Encode(user)
}

func InsertUser(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var newUser models.User
	jsonErr := json.Unmarshal(body, &newUser)
	if jsonErr != nil {
		panic(jsonErr.Error())
	}
	newId := models.CreateNewUser(newUser.Nome, newUser.Email, newUser.Email)

	newUser.Id = newId
	json.NewEncoder(rw).Encode(newUser)
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var updatedUser models.User
	jsonErr := json.Unmarshal(body, &updatedUser)
	if jsonErr != nil {
		panic(jsonErr.Error())
	}
	updatedUser.Id = id

	models.UpdateUser(id, updatedUser.Nome, updatedUser.Email, updatedUser.Fone)
	json.NewEncoder(rw).Encode(updatedUser)
}

func DeleteUserById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	models.DeleteUser(id)
	rw.WriteHeader(http.StatusNoContent)
}
