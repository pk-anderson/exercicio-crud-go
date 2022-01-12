package routes

import (
	"github.com/gorilla/mux"
	"github.com/pk-anderson/exercicio-crud-go/controllers"
)

func StartRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.MainRoute)
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users", controllers.InsertUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUserById).Methods("DELETE")
}
