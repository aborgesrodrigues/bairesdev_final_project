package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../domain"
	"../service"
	"github.com/gorilla/mux"
)

//CreateQuestion func
func createQuestion(w http.ResponseWriter, r *http.Request) {
	var question domain.Question
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Kindly enter data with the Product title and description only in order to update")
		return
	}
	err = json.Unmarshal(reqBody, &question)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	questionService := service.NewQuestionService()
	createdQuestion, error := questionService.Create(question)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdQuestion)
}

//CreateQuestion func
func updateQuestion(w http.ResponseWriter, r *http.Request) {
	var question domain.Question

	// get the parameter
	ID := mux.Vars(r)["id"]
	intID, intError := strconv.Atoi(ID)

	if intError != nil {
		fmt.Fprintf(w, intError.Error())
		return
	}

	// read the body of the request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Kindly enter data with the Product title and description only in order to update")
		return
	}

	// convert json value to question struct
	err = json.Unmarshal(reqBody, &question)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	if int(question.ID) != intID {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID of the question is different of the parameter")
		return
	}

	// call the service to update the question
	questionService := service.NewQuestionService()
	updatedQuestion, error := questionService.Update(question)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	json.NewEncoder(w).Encode(updatedQuestion)
}

//CreateQuestion func
func deleteQuestion(w http.ResponseWriter, r *http.Request) {
	// get the parameter
	ID := mux.Vars(r)["id"]
	intID, intError := strconv.Atoi(ID)

	if intError != nil {
		fmt.Fprintf(w, intError.Error())
		return
	}

	// call the service to update the question
	questionService := service.NewQuestionService()
	error := questionService.Delete(intID)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	json.NewEncoder(w).Encode("File deleted")
}

// GetOneQuestion func
func getAllQuestions(w http.ResponseWriter, r *http.Request) {
	questionService := service.NewQuestionService()
	questions, error := questionService.GetAll()

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	json.NewEncoder(w).Encode(questions)
}

// GetOneQuestion func
func getQuestionsByUser(w http.ResponseWriter, r *http.Request) {
	// get the user_id var
	userID := mux.Vars(r)["user_id"]
	intUserID, intError := strconv.Atoi(userID)

	if intError != nil {
		fmt.Fprintf(w, intError.Error())
		return
	}

	// call the service to find question by user
	questionService := service.NewQuestionService()
	questions, error := questionService.FindByUser(intUserID)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	json.NewEncoder(w).Encode(questions)

}

// CreateQuestionRouters func
func CreateQuestionRouters(router *mux.Router) {
	router.HandleFunc("/question", createQuestion).Methods("POST")
	router.HandleFunc("/question", getAllQuestions).Methods("GET")
	router.HandleFunc("/question/user/{user_id}", getQuestionsByUser).Methods("GET")
	router.HandleFunc("/question/{id}", updateQuestion).Methods("PUT")
	router.HandleFunc("/question/{id}", deleteQuestion).Methods("DELETE")
}
