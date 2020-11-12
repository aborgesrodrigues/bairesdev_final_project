package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../domain"
	"../service"
)

//CreateUser func
func createUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Kindly enter data with the Product title and description only in order to update")
		return
	}
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userService := &service.UserService{}
	createdUser, error := userService.Create(user)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// GetOneUser func
func getOneUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	intUserID, intError := strconv.Atoi(userID)

	if intError != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid number")
		return
	}

	userService := &service.UserService{}
	user, error := userService.FindByID(intUserID)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	json.NewEncoder(w).Encode(user)

}

// CreateUserRouters func
func CreateUserRouters() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user", createUser).Methods("POST")
	//router.HandleFunc("/inventory", getAllProducts).Methods("GET")
	router.HandleFunc("/user/{id}", getOneUser).Methods("GET")
	//router.HandleFunc("/inventory/{id}", updatePatchProduct).Methods("PATCH")
	//router.HandleFunc("/inventory/{id}", updateProduct).Methods("PUT")
	//router.HandleFunc("/inventory/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
