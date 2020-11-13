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

// QuestionHandler struct
type QuestionHandler struct {
	service service.QuestionServiceInterface
}

// NewQuestionHandler constructor to QuestionHandler struct
func NewQuestionHandler() *QuestionHandler {
	return &QuestionHandler{service: service.NewQuestionService()}
}

//CreateQuestion creates a new question
func (qh *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
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

	createdQuestion, error := qh.service.Create(question)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdQuestion)
}

//UpdateQuestion updates existing question
func (qh *QuestionHandler) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
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
	updatedQuestion, error := qh.service.Update(question)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	json.NewEncoder(w).Encode(updatedQuestion)
}

//DeleteQuestion deletes a question
func (qh *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	// get the parameter
	ID := mux.Vars(r)["id"]
	intID, intError := strconv.Atoi(ID)

	if intError != nil {
		fmt.Fprintf(w, intError.Error())
		return
	}

	// call the service to update the question
	error := qh.service.Delete(intID)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	json.NewEncoder(w).Encode("File deleted")
}

// GetAllQuestions return all existing questions
func (qh *QuestionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, error := qh.service.GetAll()

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	json.NewEncoder(w).Encode(questions)
}

// GetQuestionsByUser returns all the questions of an user
func (qh *QuestionHandler) GetQuestionsByUser(w http.ResponseWriter, r *http.Request) {
	// get the user_id var
	userID := mux.Vars(r)["user_id"]
	intUserID, intError := strconv.Atoi(userID)

	if intError != nil {
		fmt.Fprintf(w, intError.Error())
		return
	}

	// call the service to find question by user
	questions, error := qh.service.FindByUser(intUserID)

	if error != nil {
		fmt.Fprintf(w, error.Error())
		return
	}

	json.NewEncoder(w).Encode(questions)
}

// CreateQuestionRouters func
func CreateQuestionRouters(router *mux.Router) {
	qh := NewQuestionHandler()
	router.HandleFunc("/question", qh.CreateQuestion).Methods("POST")
	router.HandleFunc("/question", qh.GetAllQuestions).Methods("GET")
	router.HandleFunc("/question/user/{user_id}", qh.GetQuestionsByUser).Methods("GET")
	router.HandleFunc("/question/{id}", qh.UpdateQuestion).Methods("PUT")
	router.HandleFunc("/question/{id}", qh.DeleteQuestion).Methods("DELETE")
}
