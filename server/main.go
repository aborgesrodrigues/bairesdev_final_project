package main

import (
	"log"
	"net/http"

	"../internal/handler"
	"github.com/gorilla/mux"
)

//"../internal/handler"
//"../cmd"

func main() {
	//cmd.CallDao()
	//cmd.CallService()
	//cmd.CreateDatabase()
	//cmd.PopulateUser()

	router := mux.NewRouter().StrictSlash(true)
	handler.CreateUserRouters(router)
	handler.CreateQuestionRouters(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
