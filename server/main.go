package main

import (
	"log"
	"net/http"

	"bairesdev_final_project/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	handler.CreateUserRouters(router)
	handler.CreateQuestionRouters(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
