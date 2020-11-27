package main

import (
	server "bairesdev_final_project/server/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	server := server.NewServer()
	server.CreateRouters(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
