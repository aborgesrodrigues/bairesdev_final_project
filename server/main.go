package main

import (
	"log"
	"net/http"
	"os"

	"final_project/internal/handler"

	"github.com/gorilla/mux"
)

//"../internal/handler"
//"../cmd"

func main() {
	//cmd.CallDao()
	//cmd.CallService()
	//cmd.CreateDatabase()
	//cmd.PopulateUser()

	os.Setenv("SQLITEPATH", "/home/alessandro/desenvolvimento/workspaces/golang/final_project/server/sqlite-database.db")

	router := mux.NewRouter().StrictSlash(true)
	handler.CreateUserRouters(router)
	handler.CreateQuestionRouters(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
