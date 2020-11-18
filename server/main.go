package main

import (
	"log"
	"net/http"
	"os"

	"bairesdev_final_project/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	//cmd.CallDao()
	//cmd.CallService()
	//cmd.CreateDatabase()
	//cmd.PopulateUser()

	os.Setenv("SQLITEPATH", "/home/alessandro/go/src/bairesdev_final_project/server/sqlite-database.db")

	router := mux.NewRouter().StrictSlash(true)
	handler.CreateUserRouters(router)
	handler.CreateQuestionRouters(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
