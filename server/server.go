package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pk-anderson/exercicio-crud-go/server/routes"
)

func StartMuxServer() {
	router := mux.NewRouter().StrictSlash(true)
	port := ":8080"

	routes.StartRoutes(router)
	fmt.Println("Servidor rodando na porta", port)
	log.Fatal(http.ListenAndServe(port, router))
}
